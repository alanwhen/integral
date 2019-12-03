package admin

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Prepare() {
	this.BaseController.Prepare()

	this.checkLogin()
}

func (this *CategoryController) Index() {
	this.Data["pageTitle"] = "栏目管理"
	this.Data["activeSideBarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)
	//查看的权限
	this.checkAuthor()

	this.setTpl("admin/category/index.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/category/index_header.html"
	this.LayoutSections["footer"] = "admin/category/index_footer.html"

	this.Data["canEdit"] = this.checkActionAuthor("CategoryController", "Edit")
	this.Data["canDelete"] = this.checkActionAuthor("CategoryController", "Delete")
}

func (this *CategoryController) TreeGrid() {

}
