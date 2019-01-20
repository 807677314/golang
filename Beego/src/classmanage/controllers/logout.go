package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type LogoutController struct {
	//继承beego.Controller的全部属性
	beego.Controller
}

func (this *LogoutController) LogoutAction() {

	this.DelSession("username")

	this.DestroySession()

	this.SessionRegenerateID()

	this.Redirect(beego.URLFor("LoginController.LoginList"), http.StatusFound)

	this.StopRun()

}
