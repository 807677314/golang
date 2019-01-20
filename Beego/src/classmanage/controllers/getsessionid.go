package controllers

import (
	"github.com/astaxie/beego/context"
)

func GetSessionName(ctx *context.Context) string {
	sessionId := ctx.Input.Session("username")

	if nil != sessionId {
		return sessionId.(string)
	}
	return ""

}
