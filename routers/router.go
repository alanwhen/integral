package routers

import "github.com/astaxie/beego"

func init() {
	beego.Router("/admin/manage/index", &admin.ManageController{}, "*:Index")
}
