package sysinit

import "github.com/astaxie/beego"

func init() {
	//启动Session
	beego.BConfig.WebConfig.Session.SessionOn = true

	helpers.InitLogs()

	helpers.InitCache()

	InitDatabase()
}
