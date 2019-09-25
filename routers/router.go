package routers

import (
	"github.com/alanwhen/education-mini/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/manage/index", &admin.ManageController{}, "Get:Index")
	beego.Router("/admin/manage/cache", &admin.ManageController{}, "Get:ReloadCache")
	beego.Router("/admin/manage/appConf", &admin.ManageController{}, "Get,Post:appConf")
	beego.Router("/admin/manage/addConf", &admin.ManageController{}, "Get,Post:addConf")
}
