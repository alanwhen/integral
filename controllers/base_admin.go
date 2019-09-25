package controllers

import "strings"

type BaseAdminController struct {
	BaseController
}

func (this *BaseController) setTpl(template ...string) {
	layout := "shared/layout_page.html"

	var tplName string
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		ctrlName := strings.ToLower(this.controllerName[0 : len(this.controllerName)-10])
		actionName := strings.ToLower(this.actionName)
		tplName = "admin" + "/" + ctrlName + "/" + actionName + ".html"
	}
	this.Layout = layout
	this.TplName = tplName
}
