package main

import (
	"net/http"
)

func main() {

	//集中处理DAO的调用问题，方便以后可以直接调用
	InitDAO()

	//首页处理
	http.HandleFunc("/Index/", IndexAction)

	//登录版面的处理
	http.HandleFunc("/login/", LoginAction)
	http.HandleFunc("/signup/", SignupAction)
	http.HandleFunc("/logout/", LogoutAction)
	http.HandleFunc("/login/captcha/", Captcha)

	//overall 版面的处理
	http.HandleFunc("/overall/", OverallAction)
	http.HandleFunc("/overall/add/", OverallAddAction)
	http.HandleFunc("/overall/delete/", OverallDelAction)
	http.HandleFunc("/overall/update/", OverallUpdateAction)
	http.HandleFunc("/overall/batchDel/", OverallBatchDelAction)

	//class版面的处理
	http.HandleFunc("/class/list/", ClassListAction)
	http.HandleFunc("/class/list/add/", ClassUpdateAction)
	http.HandleFunc("/class/list/delete/", ClassDelAction)
	http.HandleFunc("/class/list/update/", ClassUpdateAction)
	http.HandleFunc("/class/batchDel/", ClassBatchDelAction)

	//claaroom版面的处理
	http.HandleFunc("/classroom/list/", ClassroomListAction)
	http.HandleFunc("/classroom/list/add/", ClassroomUpdateAction)
	http.HandleFunc("/classroom/list/delete/", ClassroomDelAction)
	http.HandleFunc("/classroom/list/update/", ClassroomUpdateAction)
	http.HandleFunc("/classroom/batchDel/", ClassroomBatchDelAction)

	//teacher版面的处理
	http.HandleFunc("/teacher/list/", TeacherListAction)
	http.HandleFunc("/teacher/list/add/", TeacherUpdateAction)
	http.HandleFunc("/teacher/list/delete/", TeacherDelAction)
	http.HandleFunc("/teacher/list/update/", TeacherUpdateAction)
	http.HandleFunc("/teacher/batchDel/", TeacherBatchDelAction)

	//静态资源的处理
	http.HandleFunc("/js/", AssetAction)
	http.HandleFunc("/css/", AssetAction)
	http.HandleFunc("/fonts/", AssetAction)
	http.HandleFunc("/images/", AssetAction)

	//监听
	http.ListenAndServe(":8083", nil)
}
