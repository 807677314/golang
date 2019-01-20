package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {




	c.Redirect(beego.URLFor("IndexController.Get"), http.StatusFound)

}
