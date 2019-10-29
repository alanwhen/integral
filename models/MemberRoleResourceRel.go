package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MemberRoleResourceRel struct {
	MemberRole     *MemberRole     `orm:"rel(fk)"`
	MemberResource *MemberResource `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(MemberRoleResourceRel))
}

func MemberRoleResourceRelTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "sys_member_role_resource_rel"
}

func (this *MemberRoleResourceRel) TableName() string {
	return MemberRoleResourceRelTBName()
}
