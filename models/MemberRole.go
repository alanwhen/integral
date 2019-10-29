package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MemberRoleQueryParam struct {
	BaseQueryParam
}

type MemberRole struct {
	Id                    int                      `orm:"column(role_id)" form:"Id"`
	RoleName              string                   `orm:"column(role_name)" from:"role_name"`
	ListOrder             int                      `orm:"column(list_order)" form:"ListOrder"`
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
