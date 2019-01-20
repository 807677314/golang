package routers

import (
	"classmanage/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Index/", &controllers.IndexController{})
	beego.Router("/login/",&controllers.LoginController{} ,"GET:LoginList;POST:LoginAction")
	beego.Router("/logout/",&controllers.LogoutController{} ,"GET:LogoutAction")

	classNS := beego.NewNamespace("class",
		beego.NSBefore(Check),
		beego.NSRouter("/list/", &controllers.ClassController{}, "GET:ClassList"),
		beego.NSRouter("/add/?:id:int", &controllers.ClassController{}, "GET:ClassAddList;POST:ClassAddAction"),
		beego.NSRouter("/delete/?:id:int", &controllers.ClassController{}, "Get:ClassDelete"),
		beego.NSRouter("/update/?:id:int", &controllers.ClassController{}, "Get:ClassUpdateList;POST:ClassUpdateAction"),
	)
	beego.AddNamespace(classNS)

	overallNS := beego.NewNamespace("overall",
		beego.NSBefore(Check),
		beego.NSRouter("/list/", &controllers.OverallController{}, "GET:OverallList"),
		beego.NSRouter("/add/?:id:int", &controllers.OverallController{}, "GET:OverallAddList;POST:OverallAddAction"),
		beego.NSRouter("/delete/?:id:int", &controllers.OverallController{}, "Get:OverallDelete"),
		beego.NSRouter("/update/?:id:int", &controllers.OverallController{}, "Get:OverallUpdateList;POST:OverallUpdateAction"),
	)
	beego.AddNamespace(overallNS)

	classroomNS := beego.NewNamespace("classroom",
		beego.NSBefore(Check),
		beego.NSRouter("/list/", &controllers.ClassroomController{}, "GET:ClassroomList"),
		beego.NSRouter("/add/?:id:int", &controllers.ClassroomController{}, "GET:ClassroomAddList;POST:ClassroomAddAction"),
		beego.NSRouter("/delete/?:id:int", &controllers.ClassroomController{}, "Get:ClassroomDelete"),
		beego.NSRouter("/update/?:id:int", &controllers.ClassroomController{}, "Get:ClassroomUpdateList;POST:ClassroomUpdateAction"),
	)
	beego.AddNamespace(classroomNS)

	teacherNS := beego.NewNamespace("teacher",
		beego.NSBefore(Check),
		beego.NSRouter("/list/", &controllers.TeacherController{}, "GET:TeacherList"),
		beego.NSRouter("/add/?:id:int", &controllers.TeacherController{}, "GET:TeacherAddList;POST:TeacherAddAction"),
		beego.NSRouter("/delete/?:id:int", &controllers.TeacherController{}, "Get:TeacherDelete"),
		beego.NSRouter("/update/?:id:int", &controllers.TeacherController{}, "Get:TeacherUpdateList;POST:TeacherUpdateAction"),
	)
	beego.AddNamespace(teacherNS)
}
