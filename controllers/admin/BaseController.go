package admin

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}
