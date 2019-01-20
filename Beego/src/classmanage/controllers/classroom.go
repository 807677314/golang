package controllers

import (
	"classmanage/models"
	"github.com/astaxie/beego/validation"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ClassroomController struct {
	beego.Controller
}

func (this *ClassroomController) ClassroomList() {

	o := orm.NewOrm()

	classroom := []*models.Classroom{}

	p := this.Input().Get("page")

	if p == "" {
		p = "1"
	}

	page, err := strconv.Atoi(p)

	if nil != err {
		log.Println(err)
	}

	if page <= 0 {
		page = 1
	}

	cnt, err := o.QueryTable("classroom").Count()

	var size int = 5

	offset := (page - 1) * size

	pagecount := math.Ceil(float64(cnt) / float64(size))

	qs := o.QueryTable("classroom").Limit(size, offset)

	qs.All(&classroom)

	this.Data["count"] = cnt
	this.Data["page"] = int(page)
	this.Data["prev"] = int(page) - 1
	this.Data["next"] = int(page) + 1
	this.Data["pagecount"] = int(pagecount)
	this.Data["over"] = int(pagecount) + 1

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	this.Data["classroom"] = classroom

	this.Layout = "layout/layout.html"
	this.TplName = "classroom/classroomlist.html"

}

func (this *ClassroomController) ClassroomDelete() {

	o := orm.NewOrm()

	//获得url中的传递的id值，如果发生错误就终止程序运行
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {

		this.StopRun()

	}

	//删除值为id的数据
	classroom := models.Classroom{
		Id: id,
	}

	o.Delete(&classroom)

	//制作URL给模板使用
	beego.URLFor("ClassroomController.ClassroomDelete")

	//删除完成后跳转回list界面，然后return
	this.Redirect(beego.URLFor("ClassroomController.ClassroomList"), http.StatusFound)

	this.StopRun()
}

func (this *ClassroomController) ClassroomAddList() {

	//当请求为get时，仅仅需要展示页面

	flash := beego.ReadFromRequest(&this.Controller)

	a, exists := flash.Data["ClassroomName"]

	if exists {
		this.Data["classroomname"] = a
	}
	b, exists := flash.Data["ClassroomAdress"]

	if exists {
		this.Data["classroomadress"] = b
	}

	this.Data["old"] = map[string]string{
		"classroomName":   flash.Data["classroomName"],
		"classroomAdress": flash.Data["classroomAdress"],
	}

	beego.URLFor("ClassroomController.ClassroomAddList")

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	this.Layout = "layout/layout.html"
	this.TplName = "classroom/classroomedit.html"
}

func (this *ClassroomController) ClassroomAddAction() {

	//当请求为post时，需要处理数据，先初始化orm，再获得class对象，再讲浏览器请求的数据获得添加到class属性中，插入数据库中，再跳转回list页面
	o := orm.NewOrm()

	classroom := models.Classroom{
		ClassroomName:   this.Input().Get("classroomName"),
		ClassroomAdress: this.Input().Get("classroomAdress"),
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
	}

	flash := beego.NewFlash()

	valid := validation.Validation{}

	errResult, errRule := valid.Valid(&classroom)

	if nil != errRule {
		log.Println(errRule)
	}

	if !errResult {

		for _, errResult := range valid.Errors {

			flash.Set(errResult.Field, errResult.Message)

		}

		for k, v := range this.Input() {
			flash.Set(k, v[0])
			log.Println(k, v[0])
		}

		flash.Store(&this.Controller)

		this.Redirect(beego.URLFor("ClassroomController.ClassroomAddList"), http.StatusFound)

		this.StopRun()
	}

	o.Insert(&classroom)

	this.Redirect(beego.URLFor("ClassroomController.ClassroomList"), http.StatusFound)

	this.StopRun()

}

func (this *ClassroomController) ClassroomUpdateList() {
	//当请求为get时，仅仅需要展示页面，但是还需要展示数据便于修改。
	//先初始化orm，再通过请求的id获得数据库中的数据，将其展示到模板。

	flash := beego.ReadFromRequest(&this.Controller)

	a, exists := flash.Data["ClassroomName"]

	if exists {
		this.Data["classroomname"] = a
	}
	b, exists := flash.Data["ClassroomAdress"]

	if exists {
		this.Data["classroomadress"] = b
	}

	this.Data["old"] = map[string]string{
		"classroomName":   flash.Data["classroomName"],
		"classroomAdress": flash.Data["classroomAdress"],
	}

	o := orm.NewOrm()

	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {
		this.StopRun()
	}

	classroom := models.Classroom{
		Id: id,
	}

	o.Read(&classroom)

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	this.Data["classroominfo"] = classroom
	beego.URLFor("ClassroomController.ClassroomUpdateList")

	this.Layout = "layout/layout.html"
	this.TplName = "classroom/classroomedit.html"
}
func (this *ClassroomController) ClassroomUpdateAction() {

	//当请求为post时，先初始化orm，再获得请求数据中的id但是不需要像展示页面一样获得数据，获得class结构体对象以后，直接获得修改的数据
	//将其重写class的属性，然后通过指定修数据库中的某些数据完成更新，更新完成以后需要跳转回list页面

	o := orm.NewOrm()

	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {
		this.StopRun()
	}

	classroom := models.Classroom{
		Id: id,
	}

	o.Read(&classroom)
	classroom.ClassroomName = this.Input().Get("classroomName")
	classroom.ClassroomAdress = this.Input().Get("classroomAdress")

	flash := beego.NewFlash()

	valid := validation.Validation{}

	errResult, errRule := valid.Valid(&classroom)

	if nil != errRule {
		log.Println(errRule)
	}

	if !errResult {

		for _, errResult := range valid.Errors {

			flash.Set(errResult.Field, errResult.Message)

		}

		for k, v := range this.Input() {
			flash.Set(k, v[0])
			log.Println(k, v[0])
		}

		flash.Store(&this.Controller)

		this.Redirect(this.Ctx.Request.Referer(), http.StatusFound)

		this.StopRun()
	}

	o.Update(&classroom)

	this.Redirect(beego.URLFor("ClassroomController.ClassroomList"), http.StatusFound)

	this.StopRun()
}
