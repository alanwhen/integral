package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MemberRoleQueryParam struct {
	BaseQueryParam
	NameLike string
}

type MemberRole struct {
	Id                    int    `orm:"column(role_id)" form:"Id"`
	RoleName              string `orm:"column(role_name)" from:"Role"`
	Description           string `orm:"column(description)" from:"Description"`
	ListOrder             int
	MemberRoleResourceRel []*MemberRoleResourceRel `orm:"reverse(many)" json:"-"`
	MemberRoleRel         []*MemberRoleRel         `orm:"reverse(many)" json:"-"`
}

func init() {
	orm.RegisterModel(new(MemberRole))
}

func MemberRoleTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "sys_member_role"
}

func (this *MemberRole) TableName() string {
	return MemberRoleTBName()
}

func MemberRolePageList(params *MemberRoleQueryParam) ([]*MemberRole, int64) {
	query := orm.NewOrm().QueryTable(MemberRoleTBName())
	data := make([]*MemberRole, 0)

	sortOrder := "Id"
	switch params.Sort {
	case "Id":
		sortOrder = "Id"
	case "ListOrder":
		sortOrder = "ListOrder"
	}

	if params.Order == "desc" {
		sortOrder = "-" + sortOrder
	}
	query = query.Filter("role_name__istartswith", params.NameLike)
	total, _ := query.Count()
	query.OrderBy(sortOrder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

func MemberRoleDataList(params *MemberRoleQueryParam) []*MemberRole {
	params.Limit = -1
	params.Sort = "ListOrder"
	params.Order = "asc"
	data, _ := MemberRolePageList(params)
	return data
}

func MemberRoleBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(MemberRoleTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func MemberRoleOne(id int) (*MemberRole, error) {
	o := orm.NewOrm()
	m := MemberRole{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
