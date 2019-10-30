package admin

import (
	"fmt"
	"github.com/alanwhen/education-mini/enums"
	"github.com/alanwhen/education-mini/helpers"
	"github.com/alanwhen/education-mini/models"
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
	folderName     string
	controllerName string
	actionName     string
	curUser        models.SysMember
}

func (this *BaseController) Prepare() {
	this.folderName = "admin"
	this.controllerName, this.actionName = this.GetControllerAndAction()
	this.Data["siteApp"] = beego.AppConfig.String("site.app")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.Data["siteVersion"] = beego.AppConfig.String("site.version")
	//从Session里获取数据，设置用户信息
	this.adapterUserInfo()
}

func (this *BaseController) adapterUserInfo() {
	a := this.GetSession("sys_member")
	if a != nil {
		this.curUser = a.(models.SysMember)
		this.Data["memberInfo"] = a
	}
}

func (this *BaseController) checkLogin() {
	if this.curUser.Id == 0 {
		url := this.URLFor("ManageController.Login") + "?url="

		//登录成功后返回的地址为当前页
		returnUrl := this.Ctx.Request.URL.Path

		if this.Ctx.Input.IsAjax() {
			returnUrl := this.Ctx.Input.Refer()
			this.jsonResult(enums.JRCode302, "请登录", url+returnUrl)
		}
		this.Redirect(url+returnUrl, 302)
		this.StopRun()
	}
}

func (this *BaseController) checkActionAuthor(ctrlName, ActName string) bool {
	if this.curUser.Id == 0 {
		return false
	}
	//从session获取用户信息
	user := this.GetSession("sys_member")
	//类型断言
	v, ok := user.(models.SysMember)
	if ok {
		if v.GroupId == 1 {
			return true
		}
		//遍历用户所负责的资源列表
		for i, _ := range v.ResourceUrlForList {
			urlFor := strings.TrimSpace(v.ResourceUrlForList[i])
			if len(urlFor) == 0 {
				continue
			}
			str := strings.Split(urlFor, ",")
			if len(str) > 0 && str[0] == (ctrlName+"."+ActName) {
				return true
			}
		}
	}
	return false
}

func (this *BaseController) pageLogin() {
	url := this.URLFor("ManageController.Login")
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) setTpl(template ...string) {
	layout := "shared/layout_page.html"

	var tplName string
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		ctrlName := strings.ToLower(this.controllerName[0 : len(this.controllerName)-10])
		actionName := strings.ToLower(this.actionName)
		tplName = "admin" + "/" + ctrlName + "/" + actionName + ".html"
	}

	this.Layout = layout
	this.TplName = tplName
}

func (this *BaseController) jsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	res := &models.JsonResult{Code: code, Msg: msg, Obj: obj}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) checkAuthor(ignores ...string) {
	this.checkLogin()

	for _, ignore := range ignores {
		if ignore == this.actionName {
			return
		}
	}

	hasAuthor := this.checkActionAuthor(this.controllerName, this.actionName)
	if !hasAuthor {
		helpers.LogDebug(fmt.Sprintf("author control: path=%s.%s userid=%v  无权访问", this.controllerName, this.actionName, this.curUser.Id))

		//如果没有权限
		if this.Ctx.Input.IsAjax() {
			this.jsonResult(enums.JRCode401, "无权访问", "")
		} else {
			this.pageError("无权访问")
		}
	}
}

func (this *BaseController) setMemberInfo2Session(memberId int) error {
	m, err := models.SysMemberOne(memberId)
	if err != nil {
		return err
	}
	resourceList := models.ResourceTreeGridByMemberId(memberId, 1000)
	for _, item := range resourceList {
		m.ResourceUrlForList = append(m.ResourceUrlForList, strings.TrimSpace(item.UrlFor))
	}
	this.SetSession("sys_member", *m)
	return nil
}

func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) pageError(msg string) {
	errUrl := this.URLFor("ManageController.Error") + "/" + msg
	this.Redirect(errUrl, 302)
	this.StopRun()
}
