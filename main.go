package main

import (
	_ "github.com/alanwhen/education-mini/routers"
	_ "github.com/alanwhen/education-mini/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/statics", "statics")
	beego.Run()
}
