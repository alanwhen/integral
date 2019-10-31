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

type RoleController struct {
	BaseController
}

func (this *RoleController) Prepare() {
	this.BaseController.Prepare()

	this.checkAuthor("DataGrid", "DataList", "UpdateSeq")
}

func (this *RoleController) Index() {
	this.Data["pageTitle"] = "角色管理"

	this.Data["showMoreQuery"] = false

	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)

	this.setTpl("admin/role/index.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/role/index_header.html"
	this.LayoutSections["footer"] = "admin/role/index_footer.html"

	this.Data["canEdit"] = this.checkActionAuthor("RoleController", "Edit")
	this.Data["canDelete"] = this.checkActionAuthor("RoleController", "Delete")
	this.Data["canAllocate"] = this.checkActionAuthor("RoleController", "Allocate")
}

func (this *RoleController) DataGrid() {
	var params models.MemberRoleQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)

	data, total := models.MemberRolePageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	this.Data["json"] = result
	this.ServeJSON()
}

//DataList 角色列表
func (this *RoleController) DataList() {
	var params = models.MemberRoleQueryParam{}

	//获取数据列表和总数
	data := models.MemberRoleDataList(&params)

	//定义返回的数据结构
	this.jsonResult(enums.JRCodeSuccess, "", data)
}

//Edit 添加、编辑角色界面
func (this *RoleController) Edit() {
	if this.Ctx.Request.Method == "POST" {
		this.Save()
	}

	Id, _ := this.GetInt(":id", 0)
	m := models.MemberRole{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			this.pageError("数据无效，请刷新后重试")
		}
	}
	this.Data["m"] = m
	this.setTpl("admin/role/edit.html", "shared/layout_pull_box.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["footer"] = "admin/role/edit_footer.html"
}

//Save 添加、编辑页面 保存
func (this *RoleController) Save() {
	var err error
	m := models.MemberRole{}

	//获取form里的值
	if err = this.ParseForm(&m); err != nil {
		this.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	o := orm.NewOrm()
	if m.Id == 0 {
		if _, err = o.Insert(&m); err == nil {
			this.jsonResult(enums.JRCodeSuccess, "添加成功", m.Id)
		} else {
			this.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}

	} else {
		if _, err = o.Update(&m); err == nil {
			this.jsonResult(enums.JRCodeSuccess, "编辑成功", m.Id)
		} else {
			this.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
}

//Delete 批量删除
func (this *RoleController) Delete() {
	strs := this.GetString("ids")
	ids := make([]int, 0, len(strs))

	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}

	if num, err := models.MemberRoleBatchDelete(ids); err == nil {
		this.jsonResult(enums.JRCodeSuccess, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		this.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

//Allocate 给角色分配资源界面
func (this *RoleController) Allocate() {
	roleId, _ := this.GetInt("id", 0)
	strs := this.GetString("ids")

	o := orm.NewOrm()
	m := models.MemberRole{Id: roleId}
	if err := o.Read(&m); err != nil {
		this.jsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", "")
	}

	//删除已关联的历史数据
	if _, err := o.QueryTable(models.MemberRoleResourceRelTBName()).Filter("MemberRole__id", m.Id).Delete(); err != nil {
		this.jsonResult(enums.JRCodeFailed, "删除历史关系失败", "")
	}

	var relations []models.MemberRoleResourceRel
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			r := models.MemberResource{Id: id}
			relation := models.MemberRoleResourceRel{MemberRole: &m, MemberResource: &r}
			relations = append(relations, relation)
		}
	}

	if len(relations) > 0 {
		//批量添加
		if _, err := o.InsertMulti(len(relations), relations); err == nil {
			this.jsonResult(enums.JRCodeSuccess, "保存成功", "")
		}
	}

	this.jsonResult(0, "保存失败", "")
}

func (this *RoleController) UpdateSeq() {
	Id, _ := this.GetInt("pk", 0)
	oM, err := models.MemberRoleOne(Id)
	if err != nil || oM == nil {
		this.jsonResult(enums.JRCodeFailed, "选择的数据无效", 0)
	}

	value, _ := this.GetInt("value", 0)
	oM.ListOrder = value

	o := orm.NewOrm()
	if _, err := o.Update(oM); err == nil {
		this.jsonResult(enums.JRCodeSuccess, "修改成功", oM.Id)
	} else {
		this.jsonResult(enums.JRCodeFailed, "修改失败", oM.Id)
	}
}
