package controllers

import (
	"github.com/alanwhen/education-mini/enums"
	"github.com/astaxie/beego"
	"github.com/yunnet/gardens/models"
)

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

func (this *BaseController) showJsonResult(status enums.JsonResultStatus, tips string, data interface{}) {
	res := &models.JsonResult{Status: status, Tips: tips, Data: data}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) pageLogin() {
	url := this.URLFor("ManageController.Login")
	this.Redirect(url, 302)
	this.StopRun()
}
