package admin

import (
	"github.com/alanwhen/education-mini/enums"
	"github.com/astaxie/beego/orm"
)

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
	var params models.StudentQueryParam{}
	data := models.StudentDataList(&params)
	this.jsonResult(enums.JRCodeSuccess, "", data)
}

func (this *StudentController) Edit() {
	if this.Ctx.Request.Method == "POST" {
		this.Save()
	}

	Id, _ := this.GetInt(":id", 0)
	m := models.Student{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			this.pageError("数据无效，请刷新后重试")
		}
	} else {
		//设置m的默认值
	}

	this.Data["m"] = m
	this.setTpl("admin/student/edit.html", "shared/layout_pull_box.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["footer"] = "admin/student/edit_footer.html"
}

func (this *StudentController) Save() {

}
