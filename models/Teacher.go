package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Teacher struct {
	Id int
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
