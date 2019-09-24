package sysinit

import (
	"github.com/alanwhen/integral/helpers"
	"github.com/astaxie/beego"
)

func init() {
	//启动Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//初始化日志
	helpers.InitLogs()
	//初始化缓存
	helpers.InitCache()
	//初始化数据库
	InitDatabase()
}
