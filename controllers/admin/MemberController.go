package admin

import (
	"encoding/json"
	"fmt"
	"github.com/alanwhen/education-mini/enums"
	"github.com/alanwhen/education-mini/helpers"
	"github.com/alanwhen/education-mini/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type MemberController struct {
	BaseController
}

func (this *MemberController) Prepare() {
	this.BaseController.Prepare()

	this.checkLogin()
}

func (this *MemberController) Index() {
	this.Data["pageTitle"] = "管理列表"

	this.Data["showMoreQuery"] = true

	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)

	this.setTpl("admin/member/index.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/member/index_header.html"
	this.LayoutSections["footer"] = "admin/member/index_footer.html"

	this.Data["canEdit"] = this.checkActionAuthor("MemberController", "Edit")
	this.Data["canDelete"] = this.checkActionAuthor("MemberController", "Delete")
}

func (this *MemberController) DataGrid() {
	var params models.SysMemberQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)

	data, total := models.SysMemberPageList(&params)
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MemberController) Edit() {
	if this.Ctx.Request.Method == "POST" {
		this.Save()
	}

	Id, _ := this.GetInt(":id", 0)
	m := &models.SysMember{}
	var err error
	if Id > 0 {
		m, err = models.SysMemberOne(Id)
		if err != nil {
			this.pageError("数据无效，请刷新后再试")
		}
	} else {
		m.IfLock = enums.Disabled
	}
	this.Data["m"] = m

	this.setTpl("admin/member/edit.html", "shared/layout_pull_box.html")

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["footer"] = "admin/member/edit_footer.html"
}

func (this *MemberController) Save() {
	m := models.SysMember{}
	o := orm.NewOrm()
	var err error
	if err = this.ParseForm(&m); err != nil {
		this.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	if m.Id == 0 {
		m.Password = helpers.String2md5(helpers.String2md5(m.Password + helpers.RandomString(5)))
		if _, err := o.Insert(&m); err != nil {
			this.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		if oM, err := models.SysMemberOne(m.Id); err != nil {
			this.jsonResult(enums.JRCodeFailed, "数据无效，请刷新后再试", m.Id)
		} else {
			m.Password = strings.TrimSpace(m.Password)
			if len(m.Password) == 0 {
				m.Password = oM.Password
			} else {
				m.Password = helpers.String2md5(helpers.String2md5(m.Password + oM.Encrypt))
			}
			m.Avatar = oM.Avatar
		}
		if _, err := o.Update(&m); err != nil {
			this.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
	this.jsonResult(enums.JRCodeSuccess, "保存成功", m.Id)
}

func (this *MemberController) Delete() {
	str := this.GetString("ids")
	ids := make([]int, 0, len(str))
	for _, it := range strings.Split(str, ",") {
		if id, err := strconv.Atoi(it); err == nil {
			ids = append(ids, id)
		}
	}

	query := orm.NewOrm().QueryTable(models.SysMemberTBName())
	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		this.jsonResult(enums.JRCodeSuccess, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		this.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (this *MemberController) Profile() {
	Id := this.curUser.Id
	m, err := models.SysMemberOne(Id)
	if m == nil || err != nil {
		this.pageError("数据无效，请刷新后重试")
	}
	this.Data["hasAvatar"] = len(m.Avatar) > 0
	helpers.LogDebug(m.Avatar)

	this.Data["m"] = m
	this.setTpl("admin/member/profile.html")

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/member/profile_header.html"
	this.LayoutSections["footer"] = "admin/member/profile_footer.html"
}

func (this *MemberController) BasicInfoSave() {
	Id := this.curUser.Id
	oM, err := models.SysMemberOne(Id)
	if oM == nil || err != nil {
		this.jsonResult(enums.JRCodeFailed, "数据无效，请刷新后再试", "")
	}
	m := models.SysMember{}
	if err = this.ParseForm(&m); err != nil {
		this.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	oM.FullName = m.FullName
	oM.Mobile = m.Mobile
	oM.Email = m.Email
	oM.Avatar = this.GetString("ImageUrl")
	if len(oM.Avatar) == 0 {
		oM.Avatar = "/statics/upload/tiger.png"
	}

	o := orm.NewOrm()
	if _, err := o.Update(oM); err != nil {
		this.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
	} else {
		this.setMemberInfo2Session(Id)
		this.jsonResult(enums.JRCodeSuccess, "保存成功", m.Id)
	}
}

func (this *MemberController) PasswordSave() {
	Id := this.curUser.Id
	oM, err := models.SysMemberOne(Id)
	if oM == nil || err != nil {
		this.pageError("数据无效，请稍后再试")
	}
	oldPwd := strings.TrimSpace(this.GetString("UserPwd", ""))
	newPwd := strings.TrimSpace(this.GetString("NewUserPwd", ""))
	confirmPwd := strings.TrimSpace(this.GetString("ConfirmPwd", ""))
	md5Str := helpers.String2md5(helpers.String2md5(oldPwd + oM.Encrypt))

	if oM.Password != md5Str {
		this.jsonResult(enums.JRCodeFailed, "原密码错误", "")
	}

	if len(newPwd) == 0 {
		this.jsonResult(enums.JRCodeFailed, "请输入新密码", "")
	}

	if newPwd != confirmPwd {
		this.jsonResult(enums.JRCodeFailed, "两次输入密码不一致", "")
	}

	oM.Password = helpers.String2md5(helpers.String2md5(newPwd + oM.Encrypt))
	o := orm.NewOrm()
	if _, err := o.Update(oM); err != nil {
		this.jsonResult(enums.JRCodeFailed, "保存失败", oM.Id)
	} else {
		this.setMemberInfo2Session(Id)
		this.jsonResult(enums.JRCodeSuccess, "保存成功", oM.Id)
	}
}

func (this *MemberController) UploadImage() {
	//这里type没有用，只是为了演示传值
	stype, _ := this.GetInt32("type", 0)
	if stype > 0 {
		f, h, err := this.GetFile("fileImageUrl")
		if err != nil {
			this.jsonResult(enums.JRCodeFailed, "上传失败", "")
		}
		defer f.Close()

		filePath := "statics/upload/" + h.Filename
		// 保存位置在 static/upload, 没有文件夹要先创建
		this.SaveToFile("fileImageUrl", filePath)
		this.jsonResult(enums.JRCodeSuccess, "上传成功", "/"+filePath)
	} else {
		this.jsonResult(enums.JRCodeFailed, "上传失败", "")
	}
}
