package models

import (
	"fmt"
	"github.com/alanwhen/education-mini/helpers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MemberResource struct {
	Id                    int
	Name                  string          `orm:"size(64)"`
	Parent                *MemberResource `orm:"null;rel(fk)"`
	Rtype                 int
	Seq                   int
	Sons                  []*MemberResource        `orm:"reverse(many)"` //fk 的反向关系
	SonNum                int                      `orm:"-"`
	Icon                  string                   `orm:"size(32)"`
	LinkUrl               string                   `orm:"-"`
	UrlFor                string                   `orm:"size(255)" Json:"-"`
	HtmlDisabled          int                      `orm:"-"`
	Level                 int                      `orm:"-"`
	MemberRoleResourceRel []*MemberRoleResourceRel `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(MemberResource))
}

func (this *MemberResource) TableName() string {
	return MemberResourceTBName()
}

func MemberResourceTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "sys_member_resource"
}

func MemberResourceOne(id int) (*MemberResource, error) {
	o := orm.NewOrm()
	m := MemberResource{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func ResourceTreeGrid() []*MemberResource {
	o := orm.NewOrm()
	query := o.QueryTable(MemberResourceTBName()).OrderBy("seq", "id")
	list := make([]*MemberResource, 0)
	query.All(&list)
	return resourceList2TreeGrid(list)
}

func ResourceTreeGrid4Parent(id int) []*MemberResource {
	tree := ResourceTreeGrid()
	if id == 0 {
		return tree
	}
	var index = -1
	for i, _ := range tree {
		if tree[i].Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		return tree
	} else {
		tree[index].HtmlDisabled = 1
		for _, item := range tree[index+1:] {
			if item.Level > tree[index].Level {
				item.HtmlDisabled = 1
			} else {
				break
			}
		}
	}
	return tree
}

func ResourceTreeGridByMemberId(memberId, maxrtype int) []*MemberResource {
	cacheKey := fmt.Sprintf("eds_ResourceTreeGridByMemberId_%v_%v", memberId, maxrtype)
	var list []*MemberResource
	if err := helpers.GetCache(cacheKey, &list); err == nil {
		return list
	}

	o := orm.NewOrm()
	user, err := SysMemberOne(memberId)
	if err != nil || user == nil {
		return list
	}

	var sql string
	if user.GroupId == 1 {
		sql = fmt.Sprintf(`SELECT id,name,parent_id,rtype,icon,seq,url_for FROM %s WHERE rtype <= ? Order By seq asc,Id ASC`, MemberResourceTBName())
		o.Raw(sql, maxrtype).QueryRows(&list)
	} else {
		sql = fmt.Sprintf(`SELECT DISTINCT T0.resource_id,T2.id,T2.name,T2.parent_id,T2.rtype,T2.icon,T2.seq,T2.url_for
		FROM %s AS T0
		INNER JOIN %s AS T1 ON T0.role_id = T1.role_id
		INNER JOIN %s AS T2 ON T2.id = T0.resource_id
		WHERE T1.backend_user_id = ? and T2.rtype <= ?  Order By T2.seq asc,T2.id asc`, MemberRoleResourceRelTBName(), MemberRoleRelTBName(), MemberResourceTBName())
		o.Raw(sql, memberId, maxrtype).QueryRows(&list)
	}
	result := resourceList2TreeGrid(list)
	helpers.SetCache(cacheKey, result, 30)
	return result
}

func resourceList2TreeGrid(list []*MemberResource) []*MemberResource {
	result := make([]*MemberResource, 0)
	for _, item := range list {
		if item.Parent == nil || item.Parent.Id == 0 {
			item.Level = 0
			result = append(result, item)
			result = resourceAddSons(item, list, result)
		}
	}
	return result
}

func resourceAddSons(cur *MemberResource, list, result []*MemberResource) []*MemberResource {
	for _, item := range list {
		if item.Parent != nil && item.Parent.Id == cur.Id {
			cur.SonNum++
			item.Level = cur.Level + 1
			result = append(result, item)
			result = resourceAddSons(item, list, result)
		}
	}
	return result
}
