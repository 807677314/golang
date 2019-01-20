package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username
	this.Layout = "layout/layout.html"
	this.TplName = "Index/index.html"
}
