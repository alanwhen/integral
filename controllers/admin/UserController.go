package admin

import (
	"encoding/json"
	"fmt"
	"github.com/alanwhen/education-mini/enums"
	"github.com/alanwhen/education-mini/models"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

type UserController struct {
	BaseController
}

func (this *UserController) Prepare() {
	this.BaseController.Prepare()

	this.checkLogin()
}

func (this *UserController) Index() {
	this.Data["pageTitle"] = "用户列表"
	this.Data["showMoreQuery"] = true
	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)

	this.setTpl("admin/user/index.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/user/index_header.html"
	this.LayoutSections["footer"] = "admin/user/index_footer.html"

	this.Data["canEdit"] = this.checkActionAuthor("UserController", "Edit")
	this.Data["canDelete"] = this.checkActionAuthor("UserController", "Delete")
}

func (this *UserController) DataGrid() {
	var params models.UserQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)

	data, total := models.UserPageList(&params)
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *UserController) Edit() {
	if this.Ctx.Request.Method == "POST" {
		this.Save()
	}

	Id, _ := this.GetInt64(":id", 0)
	m := &models.User{}
	var err error
	if Id > 0 {
		m, err = models.UserOne(Id)
		if err != nil {
			this.pageError("数据无效，请刷新后再试")
		}
	} else {

	}
	this.Data["m"] = m
	this.setTpl("admin/user/edit.html", "shared/layout_pull_box.html")

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["footer"] = "admin/user/edit_footer.html"
}

func (this *UserController) Save() {
	m := models.User{}
	o := orm.NewOrm()
	var err error
	if err = this.ParseForm(&m); err != nil {
		this.jsonResult(enums.JRCodeFailed, "数据获取失败", m.Id)
	}

	if m.Id == 0 {
		password := []byte(m.Password)
		hashPwd, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err == nil {
			this.jsonResult(enums.JRCodeFailed, "密码生成错误", m.Id)
		}
		m.Password = string(hashPwd)
		if _, err := o.Insert(&m); err != nil {
			this.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		if oM, err := models.UserOne(m.Id); err != nil {
			this.jsonResult(enums.JRCodeFailed, "数据无效，请刷新后再试", m.Id)
		} else {
			m.Password = strings.TrimSpace(m.Password)
			if len(m.Password) == 0 {
				m.Password = oM.Password
			} else {
				hashPwd, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
				if err == nil {
					this.jsonResult(enums.JRCodeFailed, "密码生成错误", m.Id)
				}
				m.Password = string(hashPwd)
			}
			m.Avatar = oM.Avatar
		}
		if _, err := o.Update(&m); err != nil {
			this.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
	this.jsonResult(enums.JRCodeSuccess, "保存成功", m.Id)
}

func (this *UserController) Delete() {
	str := this.GetString("ids")
	ids := make([]int, 0, len(str))
	for _, it := range strings.Split(str, ",") {
		if id, err := strconv.Atoi(it); err == nil {
			ids = append(ids, id)
		}
	}

	query := orm.NewOrm().QueryTable(models.UserTBName())
	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		this.jsonResult(enums.JRCodeSuccess, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		this.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
