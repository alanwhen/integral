package admin

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
	folderName     string
	controllerName string
	actionName     string
}

func (this *BaseController) Prepare() {
	this.folderName = "admin"
	this.controllerName, this.actionName = this.GetControllerAndAction()

}
