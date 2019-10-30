package admin

import (
	"fmt"
	"github.com/alanwhen/education-mini/enums"
	"github.com/alanwhen/education-mini/helpers"
	"github.com/alanwhen/education-mini/models"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

type ManageController struct {
	BaseController
}

func (this *ManageController) Login() {
	if this.Ctx.Request.Method == "POST" {
		remoteAddr := this.Ctx.Request.RemoteAddr
		addrs := strings.Split(remoteAddr, "::1")
		if len(addrs) > 1 {
			remoteAddr = "localhost"
		}
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		if err := models.LoginTraceAdd(username, remoteAddr, time.Now()); err != nil {
			beego.Error("LoginTraceAdd error.")
		}
		beego.Info(fmt.Sprintf("login: %s IP: %s", username, remoteAddr))

		if len(username) == 0 || len(password) == 0 {
			this.jsonResult(enums.JRCodeFailed, "用户名或密码不正确", "")
		}
		user, err := models.SysMemberOneByUserName(username)
		if user != nil && err == nil {
			if user.IfLock == 1 {
				this.jsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
			}
			password = helpers.String2md5(helpers.String2md5(password) + user.Encrypt)
			if user.Password != password {
				this.jsonResult(enums.JRCodeFailed, "密码不正确", "")
			}

			this.setMemberInfo2Session(user.Id)

			this.jsonResult(enums.JRCodeSuccess, "登录成功", "")
		} else {
			this.jsonResult(enums.JRCodeFailed, "用户不存在", "")
		}

	}
	this.Data["pageTitle"] = beego.AppConfig.String("site.app") + beego.AppConfig.String("site.name") + " - 登录"
	this.Data["siteVersion"] = beego.AppConfig.String("site.version")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/manage/login_header.html"
	this.LayoutSections["footer"] = "admin/manage/login_footer.html"
	this.setTpl("admin/manage/login.html", "shared/layout_base.html")
}

func (this *ManageController) Logout() {
	user := models.SysMember{}
	this.SetSession("sys_member", user)
	this.pageLogin()
}

func (this *ManageController) Index() {
	this.Data["pageTitle"] = "首页"

	this.checkLogin()
	this.setTpl("admin/manage/index.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/manage/index_header.html"
	this.LayoutSections["footer"] = "admin/manage/index_footer.html"
}

func (this *ManageController) ReloadCache() {

}

func (this *ManageController) appConf() {

}

func (this *ManageController) addConf() {

}

func (this *ManageController) Error() {
	this.Data["error"] = this.GetString(":error")
	this.setTpl("manage/error.html", "shared/layout_pull_box.html")
}
