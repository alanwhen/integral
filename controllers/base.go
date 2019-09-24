package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
	folderName     string
	controllerName string
	actionName     string
}

func initSite() {

}

func (this *BaseController) Prepare() {
	this.folderName = "admin"
	this.controllerName, this.actionName = this.GetControllerAndAction()
}
