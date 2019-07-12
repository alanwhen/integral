package models

import (
	"github.com/astaxie/beego"
)

func tableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}
