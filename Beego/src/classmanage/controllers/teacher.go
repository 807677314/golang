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

type TeacherController struct {
	beego.Controller
}

func (this *TeacherController) TeacherList() {

	o := orm.NewOrm()

	teacher := []*models.Teacher{}

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

	cnt, err := o.QueryTable("teacher").Count()

	var size int = 5

	offset := (page - 1) * size

	pagecount := math.Ceil(float64(cnt) / float64(size))

	qs := o.QueryTable("teacher").Limit(size, offset)

	qs.All(&teacher)

	this.Data["count"] = cnt
	this.Data["page"] = int(page)
	this.Data["prev"] = int(page) - 1
	this.Data["next"] = int(page) + 1
	this.Data["pagecount"] = int(pagecount)
	this.Data["over"] = int(pagecount) + 1

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	this.Data["teacher"] = teacher

	this.Layout = "layout/layout.html"
	this.TplName = "teacher/teacherlist.html"
}

func (this *TeacherController) TeacherDelete() {

	o := orm.NewOrm()

	//获得url中的传递的id值，如果发生错误就终止程序运行
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {

		this.StopRun()

	}

	//删除值为id的数据
	teacher := models.Teacher{
		Id: id,
	}

	o.Delete(&teacher)

	//制作URL给模板使用
	beego.URLFor("TeacherController.TeacherDelete")

	//删除完成后跳转回list界面，然后return
	this.Redirect(beego.URLFor("TeacherController.TeacherList"), http.StatusFound)

	this.StopRun()
}

func (this *TeacherController) TeacherAddList() {

	//当请求为get时，仅仅需要展示页面

	flash := beego.ReadFromRequest(&this.Controller)

	a, exists := flash.Data["TeacherName"]

	if exists {
		this.Data["teachername"] = a
	}

	this.Data["old"] = map[string]string{
		"teacherName": flash.Data["teacherName"],
		"gender":      flash.Data["gender"],
	}

	beego.URLFor("TeacherController.TeacherAddList")

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	this.Layout = "layout/layout.html"
	this.TplName = "teacher/teacheredit.html"
}

func (this *TeacherController) TeacherAddAction() {

	//当请求为post时，需要处理数据，先初始化orm，再获得class对象，再讲浏览器请求的数据获得添加到class属性中，插入数据库中，再跳转回list页面
	o := orm.NewOrm()

	teacher := models.Teacher{
		TeacherName: this.Input().Get("teacherName"),
		Gender:      this.Input().Get("gender"),
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}

	flash := beego.NewFlash()

	valid := validation.Validation{}

	errResult, errRule := valid.Valid(&teacher)

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

		this.Redirect(beego.URLFor("TeacherController.TeacherAddList"), http.StatusFound)

		this.StopRun()
	}

	row, err := o.Insert(&teacher)
	log.Println(row, err)

	this.Redirect(beego.URLFor("TeacherController.TeacherList"), http.StatusFound)

	this.StopRun()

}

func (this *TeacherController) TeacherUpdateList() {
	//当请求为get时，仅仅需要展示页面，但是还需要展示数据便于修改。
	//先初始化orm，再通过请求的id获得数据库中的数据，将其展示到模板。

	flash := beego.ReadFromRequest(&this.Controller)

	a, exists := flash.Data["TeacherName"]

	if exists {
		this.Data["teachername"] = a
	}

	this.Data["old"] = map[string]string{
		"teacherName": flash.Data["teacherName"],
		"gender":      flash.Data["gender"],
	}

	o := orm.NewOrm()

	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {
		this.StopRun()
	}

	teacher := models.Teacher{
		Id: id,
	}

	o.Read(&teacher)

	this.Data["teacherinfo"] = teacher

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	beego.URLFor("TeacherController.TeacherUpdateList")

	this.Layout = "layout/layout.html"
	this.TplName = "teacher/teacheredit.html"
}
func (this *TeacherController) TeacherUpdateAction() {

	//当请求为post时，先初始化orm，再获得请求数据中的id但是不需要像展示页面一样获得数据，获得class结构体对象以后，直接获得修改的数据
	//将其重写class的属性，然后通过指定修数据库中的某些数据完成更新，更新完成以后需要跳转回list页面

	o := orm.NewOrm()

	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {
		this.StopRun()
	}

	teacher := models.Teacher{
		Id: id,
	}

	o.Read(&teacher)

	teacher.TeacherName = this.Input().Get("teacherName")
	teacher.Gender = this.Input().Get("gender")

	flash := beego.NewFlash()

	valid := validation.Validation{}

	errResult, errRule := valid.Valid(&teacher)

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

	o.Update(&teacher)

	this.Redirect(beego.URLFor("TeacherController.TeacherList"), http.StatusFound)

	this.StopRun()
}
