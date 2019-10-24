package admin

import "github.com/astaxie/beego"

type ManageController struct {
	BaseController
}

func (this *ManageController) Login() {
	this.Data["pageTitle"] = beego.AppConfig.String("site.app") + beego.AppConfig.String("site.name") + " - 登录"
	this.Data["siteVersion"] = beego.AppConfig.String("site.version")

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "home/header.html"
	this.LayoutSections["footer"] = "home/footer.html"
	this.setTpl("")
}

func (this *ManageController) Index() {

}

func (this *ManageController) ReloadCache() {

}

func (this *ManageController) appConf() {

}

func (this *ManageController) addConf() {

}
