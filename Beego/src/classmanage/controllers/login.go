package controllers

import (
	"classmanage/models"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"net/http"

	"github.com/astaxie/beego"
)

type LoginController struct {
	//继承beego.Controller的全部属性
	beego.Controller
}

func (this *LoginController) LoginList() {

	flash := beego.ReadFromRequest(&this.Controller)



	a, exists := flash.Data["loginError"]

	if exists {
		this.Data["loginError"] =a

		log.Println(this.Data)
	}

	this.TplName = "login/login.html"

}

func (this *LoginController) LoginAction() {

	user := this.Input().Get("user")
	password := fmt.Sprintf("%x", md5.Sum([]byte(this.Input().Get("password"))))

	o := orm.NewOrm()

	userinfo := models.Userinfo{
		User: user,
	}
	errread := o.Read(&userinfo, "User")

	if nil != errread {
		log.Println(errread)
	}

	if user == userinfo.User && password == userinfo.Password {

		this.SetSession("username", userinfo.UserName)

		this.Redirect(beego.URLFor("MainController.Get"), http.StatusFound)

		this.StopRun()

	} else {

		flash := beego.NewFlash()

		flash.Set("loginError", "用户名或密码错误")

		flash.Store(&this.Controller)

		this.Redirect(beego.URLFor("LoginController.LoginList"), http.StatusFound)

		this.StopRun()
	}

}
