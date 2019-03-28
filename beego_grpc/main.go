package main

import (
	_ "beego_grpc/initial"
	_ "beego_grpc/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.Run()

}
