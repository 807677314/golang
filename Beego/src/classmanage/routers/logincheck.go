package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)

func Check(ctx *context.Context) {

	sessionId := ctx.Input.Session("username")

	if nil == sessionId {

		ctx.Redirect(http.StatusFound, beego.URLFor("LoginController.LoginList"))

	}

}

