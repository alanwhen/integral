package admin

type ResourceController struct {
	BaseController
}

func (this *ResourceController) Prepare() {
	this.BaseController.Prepare()

	this.checkLogin()
}

func (this *ResourceController) Index() {
	this.Data["pageTitle"] = "资源管理"

	//需要权限控制
	this.checkAuthor()
}
