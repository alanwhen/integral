package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type MemberRoleResourceRel struct {
	Id             int
	MemberRole     *MemberRole     `orm:"rel(fk)"`
	MemberResource *MemberResource `orm:"rel(fk)"`
	Created        time.Time       `orm:"auto_now_add;type(datetime)"`
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
