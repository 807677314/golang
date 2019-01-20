package main

import (
	_ "classmanage/routers"
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterDataBase("default", "mysql", "root:ld123456@tcp(127.0.0.1:3306)/classmanage?charset=utf8&loc=Asia%2FShanghai", 30)

	//orm.RunCommand()
}

func main() {

	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}
