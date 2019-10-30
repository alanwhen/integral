package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MemberRoleRel struct {
	Id         int
	MemberRole *MemberRole `orm:"rel(fk)"`
	SysMember  *SysMember  `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(MemberRoleRel))
}

func MemberRoleRelTBName() string {
	return beego.AppConfig.String("db_dt_prefix") + "sys_member_role_rel"
}

func (this *MemberRoleRel) TableName() string {
	return MemberRoleRelTBName()
}
