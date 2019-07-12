package main

import (
	_ "github.com/alanwhen/integral/routers"
	_ "github.com/alanwhen/integral/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
