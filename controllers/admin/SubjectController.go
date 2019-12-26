package admin

import (
	"encoding/json"
	"fmt"
	"github.com/alanwhen/education-mini/enums"
	"github.com/alanwhen/education-mini/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type SubjectController struct {
	BaseController
}

func (this *SubjectController) Prepare() {
	this.BaseController.Prepare()
	this.checkAuthor("DataGrid")
}

func (this *SubjectController) Index() {
	this.Data["pageTitle"] = "学校管理"

	this.Data["showMoreQuery"] = true
	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)

	this.setTpl("admin/subject/index.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/subject/index_header.html"
	this.LayoutSections["footer"] = "admin/subject/index_footer.html"

	this.Data["canEdit"] = this.checkActionAuthor(this.controllerName, "Edit")
	this.Data["canDelete"] = this.checkActionAuthor(this.controllerName, "Delete")
}

func (this *SubjectController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.SubjectQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)

	data, total := models.SubjectPageList(&params)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *SubjectController) Edit() {
	if this.Ctx.Request.Method == "POST" {
		this.Save()
	}

	Id, _ := this.GetInt(":id", 0)
	m := models.Subject{Id: Id}
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
	this.setTpl("admin/school/edit.html", "shared/layout_pull_box.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["footer"] = "admin/school/edit_footer.html"
}

func (this *SubjectController) Save() {
	var err error
	m := models.Subject{}

	if err = this.ParseForm(&m); err != nil {
		this.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	o := orm.NewOrm()
	if m.Id == 0 {
		if err = o.Begin(); err != nil {
			this.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
			return
		}

		if _, err = o.Insert(&m); err == nil {
			if err = o.Commit(); err != nil {
				this.jsonResult(enums.JRCodeFailed, "添加提交失败", m.Id)
				o.Rollback()
			} else {
				this.jsonResult(enums.JRCodeSuccess, "添加成功", m.Id)
			}
		} else {
			if err = o.Rollback(); err != nil {
				this.jsonResult(enums.JRCodeFailed, "添加回滚失败", m.Id)
			} else {
				this.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
			}
		}
	} else {
		if _, err = o.Update(&m); err == nil {
			this.jsonResult(enums.JRCodeSuccess, "编辑成功", m.Id)
		} else {
			this.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
}

func (this *SubjectController) Delete() {
	strS := this.GetString("ids")
	ids := make([]int, 0, len(strS))
	for _, str := range strings.Split(strS, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}

	if num, err := models.SubjectBatchDelete(ids); err == nil {
		this.jsonResult(enums.JRCodeSuccess, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		this.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
