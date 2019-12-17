package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Teacher struct {
	Id           int
	School       *School
	Subject      *Subject
	TeacherName  string `orm:"size()"`
	TeacherNo    string `orm:"size()"`
	Introduction string `orm:"type(text)"`
	Sex          byte
	Avatar       string    `orm:"size(150)"`
	Birthday     time.Time `orm:"type(date)"`
	Mobile       string    `orm:"size(22)"`
	Password     string
	AddTime      int
	Openid       string `orm:"size(50)"`
}

func init() {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("db_tb_prefix"), new(Teacher))
}

func (this *Teacher) TableName() string {
	return TeacherTBName()
}

func TeacherTBName() string {
	return "teacher"
}
