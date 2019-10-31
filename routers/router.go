package routers

import (
	"github.com/alanwhen/education-mini/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/manage/login", &admin.ManageController{}, "*:Login")
	beego.Router("/admin/manage/index", &admin.ManageController{}, "Get:Index")
	beego.Router("/admin/manage/cache", &admin.ManageController{}, "Get:ReloadCache")

	beego.Router("/admin/resource/index", &admin.ResourceController{}, "*:Index")
	beego.Router("/admin/resource/treeGrid", &admin.ResourceController{}, "POST:TreeGrid")
	beego.Router("/admin/resource/edit/?:id", &admin.ResourceController{}, "Get,Post:Edit")
	beego.Router("/admin/resource/parent", &admin.ResourceController{}, "Post:ParentTreeGrid")
	beego.Router("/admin/resource/delete", &admin.ResourceController{}, "Post:Delete")

	beego.Router("admin/resource/select", &admin.ResourceController{}, "Get:Select")
	beego.Router("admin/resource/chooseIcon", &admin.ResourceController{}, "Get:ChooseIcon")

	beego.Router("/admin/resource/usermenutree", &admin.ResourceController{}, "POST:UserMenuTree")
	beego.Router("/admin/resource/checkUrlFor", &admin.ResourceController{}, "POST:CheckUrlFor")

	beego.Router("/", &admin.ManageController{}, "*:Index")
}
