package admin

import (
	"fmt"
	"github.com/alanwhen/education-mini/enums"
	"github.com/alanwhen/education-mini/helpers"
	"github.com/alanwhen/education-mini/models"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"strconv"
	"strings"
)

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

	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)
	this.setTpl("admin/resource/index.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/resource/index_header.html"
	this.LayoutSections["footer"] = "admin/resource/index_footer.html"

	this.Data["canEdit"] = this.checkActionAuthor("ResourceController", "Edit")
	this.Data["canDelete"] = this.checkActionAuthor("ResourceController", "Delete")
}

func (this *ResourceController) TreeGrid() {
	tree := models.ResourceTreeGrid()

	this.UrlFor2Link(tree)
	this.jsonResult(enums.JRCodeSuccess, "", tree)
}

func (this *ResourceController) UserMenuTree() {
	memberId := this.curUser.Id

	tree := models.ResourceTreeGridByMemberId(memberId, 1)

	this.UrlFor2Link(tree)
	this.jsonResult(enums.JRCodeSuccess, "", tree)
}

func (this *ResourceController) ParentTreeGrid() {
	Id, _ := this.GetInt("id", 0)
	tree := models.ResourceTreeGrid4Parent(Id)

	this.UrlFor2Link(tree)
	this.jsonResult(enums.JRCodeSuccess, "", tree)
}

func (this *ResourceController) UrlFor2LinkOne(urlfor string) string {
	if len(urlfor) == 0 {
		return ""
	}

	str := strings.Split(urlfor, ",")
	if len(str) == 1 {
		return this.URLFor(str[0])
	} else if len(str) > 1 {
		var values []interface{}
		for _, val := range str[1:] {
			values = append(values, val)
		}

		return this.URLFor(str[0], values...)
	}
	return ""
}

func (this *ResourceController) UrlFor2Link(src []*models.MemberResource) {
	for _, item := range src {
		item.LinkUrl = this.UrlFor2LinkOne(item.UrlFor)
	}
}

func (this *ResourceController) Edit() {
	//权限
	//this.checkAuthor()

	if this.Ctx.Request.Method == "POST" {
		this.Save()
	}

	Id, _ := this.GetInt(":id", 0)
	m := &models.MemberResource{}
	var err error
	if Id == 0 {
		m.Seq = 100
	} else {
		m, err = models.MemberResourceOne(Id)
		if err != nil {
			this.pageError("数据无效，请刷新后再试")
		}
	}

	if m.Parent != nil {
		this.Data["parent"] = m.Parent.Id
	} else {
		this.Data["parent"] = 0
	}

	this.Data["parents"] = models.ResourceTreeGrid4Parent(Id)

	m.LinkUrl = this.UrlFor2LinkOne(m.UrlFor)

	this.Data["m"] = m

	this.setTpl("admin/resource/edit.html", "shared/layout_pull_box.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["footer"] = "admin/resource/edit_footer.html"
}

func (this *ResourceController) Save() {
	var err error
	o := orm.NewOrm()
	parent := &models.MemberResource{}
	m := models.MemberResource{}
	parentId, _ := this.GetInt("Parent", 0)

	//获取form里的值
	if err = this.ParseForm(&m); err != nil {
		helpers.LogDebug(err)
		this.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	//获取父节点
	if parentId > 0 {
		parent, err = models.MemberResourceOne(parentId)
		if err == nil && parent != nil {
			m.Parent = parent
		} else {
			this.jsonResult(enums.JRCodeFailed, "父节点无效", "")
		}
	}

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

func (this *ResourceController) Delete() {
	this.checkAuthor()
	Id, _ := this.GetInt("Id", 0)
	if Id == 0 {
		this.jsonResult(enums.JRCodeFailed, "选择数据无效", 0)
	}

	query := orm.NewOrm().QueryTable(models.MemberResourceTBName())
	if _, err := query.Filter("id", Id).Delete(); err == nil {
		this.jsonResult(enums.JRCodeSuccess, fmt.Sprintf("删除成功"), 0)
	} else {
		this.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

// Select 通用选择面板
func (this *ResourceController) Select() {
	//获取调用者的类别 1表示 角色
	desttype, _ := this.GetInt("desttype", 0)

	//获取调用者的值
	destval, _ := this.GetInt("destval", 0)

	//返回的资源列表
	var selectedIds []string
	o := orm.NewOrm()
	if desttype > 0 && destval > 0 {
		//如果都大于0,则获取已选择的值，例如：角色，就是获取某个角色已关联的资源列表
		switch desttype {
		case 1:
			{
				role := models.MemberRole{Id: destval}
				o.LoadRelated(&role, "MemberRoleResourceRel")
				for _, item := range role.MemberRoleResourceRel {
					selectedIds = append(selectedIds, strconv.Itoa(item.MemberResource.Id))
				}
			}
		}
	}

	this.Data["selectedIds"] = strings.Join(selectedIds, ",")
	this.setTpl("admin/resource/select.html", "shared/layout_pull_box.html")

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/resource/select_header.html"
	this.LayoutSections["footer"] = "admin/resource/select_footer.html"
}

func (this *ResourceController) ChooseIcon() {
	filename := "statics/plugins/font-awesome/less/variables.less"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}

	var iconList []string
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, "@fa-var-") {
			tempStr := line[8:]
			idx := strings.Index(tempStr, ":")
			icon := tempStr[:idx]
			iconList = append(iconList, icon)
		}
	}
	this.Data["IconList"] = iconList
	this.setTpl("admin/resource/choose_icon.html", "shared/layout_pull_box.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = "admin/resource/choose_icon_header.html"
	this.LayoutSections["footer"] = "admin/resource/choose_icon_footer.html"
}

//CheckUrlFor 填写UrlFor时进行验证
func (this *ResourceController) CheckUrlFor() {
	urlfor := this.GetString("urlfor")
	link := this.UrlFor2LinkOne(urlfor)
	if len(link) > 0 {
		this.jsonResult(enums.JRCodeSuccess, "解析成功", link)
	} else {
		this.jsonResult(enums.JRCodeFailed, "解析失败", link)
	}
}

func (this *ResourceController) UpdateSeq() {
	Id, _ := this.GetInt("pk", 0)

	oM, err := models.MemberResourceOne(Id)
	if err != nil || oM == nil {
		this.jsonResult(enums.JRCodeFailed, "选择的数据无效", 0)
	}

	value, _ := this.GetInt("value", 0)
	oM.Seq = value
	if _, err := orm.NewOrm().Update(oM); err == nil {
		this.jsonResult(enums.JRCodeSuccess, "修改成功", oM.Id)
	} else {
		this.jsonResult(enums.JRCodeFailed, "修改失败", oM.Id)
	}
}
