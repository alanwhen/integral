package admin

type StudentController struct {
	BaseController
}

func (this *StudentController) Prepare() {
	this.BaseController.Prepare()

	this.checkLogin()
}

func (this *StudentController) Index() {
	this.Data["pageTitle"] = "学籍列表"
	this.Data["showMoreQuery"] = true
	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)

	this.setTpl("admin/student/index.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/student/index_header.html"
	this.LayoutSections["footer"] = "admin/student/index_footer.html"

	this.Data["canEdit"] = this.checkActionAuthor(this.controllerName, "Edit")
	this.Data["canDelete"] = this.checkActionAuthor(this.controllerName, "Delete")
}

func (this *StudentController) DataGrid() {

}
