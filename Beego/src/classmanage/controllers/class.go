package controllers

import (
	"classmanage/models"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type ClassController struct {
	//继承beego.Controller的全部属性
	beego.Controller
}

func (this *ClassController) ClassList() {

	var classes []*models.Class

	//初始化一个orm
	o := orm.NewOrm()

	//查询全部字段，将其映射到class结构体中

	//传递到模板中展示

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

	cnt, err := o.QueryTable("class").Count()

	var size int = 5

	offset := (page - 1) * size

	pagecount := math.Ceil(float64(cnt) / float64(size))

	qs := o.QueryTable("class").Limit(size, offset)

	qs.All(&classes)

	this.Data["count"] = cnt
	this.Data["page"] = int(page)
	this.Data["prev"] = int(page) - 1
	this.Data["next"] = int(page) + 1
	this.Data["pagecount"] = int(pagecount)
	this.Data["over"] = int(pagecount) + 1

	//确定布局模板是哪个

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username
	this.Data["classes"] = classes

	this.Layout = "layout/layout.html"
	this.TplName = "class/classlist.html"

}

func (this *ClassController) ClassDelete() {

	o := orm.NewOrm()

	//获得url中的传递的id值，如果发生错误就终止程序运行
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {

		this.StopRun()

	}

	//删除值为id的数据
	class := models.Class{
		Id: id,
	}

	o.Delete(&class)

	//制作URL给模板使用
	beego.URLFor("ClassController.ClassDelete")

	//删除完成后跳转回list界面，然后return
	this.Redirect(beego.URLFor("ClassController.ClassList"), http.StatusFound)

	this.StopRun()
}

func (this *ClassController) ClassAddList() {

	flash := beego.ReadFromRequest(&this.Controller)

	a, exists := flash.Data["ClassName"]

	if exists {
		this.Data["classname"] = a
	}
	b, exists := flash.Data["ClassDesc"]

	if exists {
		this.Data["classdesc"] = b
	}

	this.Data["old"] = map[string]string{
		"className":        flash.Data["className"],
		"classDescription": flash.Data["classDescription"],
	}

	//当请求为get时，仅仅需要展示页面

	beego.URLFor("ClassController.ClassAddList")

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	this.Layout = "layout/layout.html"
	this.TplName = "class/classedit.html"
}

func (this *ClassController) ClassAddAction() {

	//当请求为post时，需要处理数据，先初始化orm，再获得class对象，再讲浏览器请求的数据获得添加到class属性中，插入数据库中，再跳转回list页面
	o := orm.NewOrm()

	class := models.Class{
		ClassName: this.Input().Get("className"),
		ClassDesc: this.Input().Get("classDescription"),
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}

	flash := beego.NewFlash()

	valid := validation.Validation{}

	errResult, errRule := valid.Valid(&class)

	if nil != errRule {
		log.Println(errRule)
	}

	if !errResult {

		for _, errResult := range valid.Errors {

			flash.Set(errResult.Field, errResult.Message)

		}

		for k, v := range this.Input() {
			flash.Set(k, v[0])

		}

		flash.Store(&this.Controller)

		this.Redirect(beego.URLFor("ClassController.ClassAddList"), http.StatusFound)

		this.StopRun()
	}

	o.Insert(&class)

	this.Redirect(beego.URLFor("ClassController.ClassList"), http.StatusFound)

	this.StopRun()

}

func (this *ClassController) ClassUpdateList() {
	//当请求为get时，仅仅需要展示页面，但是还需要展示数据便于修改。
	//先初始化orm，再通过请求的id获得数据库中的数据，将其展示到模板。
	flash := beego.ReadFromRequest(&this.Controller)

	a, exists := flash.Data["ClassName"]

	if exists {
		this.Data["classname"] = a
	}
	b, exists := flash.Data["ClassDesc"]

	if exists {
		this.Data["classdesc"] = b
	}

	this.Data["old"] = map[string]string{
		"className":        flash.Data["className"],
		"classDescription": flash.Data["classDescription"],
	}

	o := orm.NewOrm()

	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {
		this.StopRun()
	}

	class := models.Class{
		Id: id,
	}

	o.Read(&class)

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	this.Data["classinfo"] = class
	beego.URLFor("ClassController.ClassUpdateList")

	this.Layout = "layout/layout.html"
	this.TplName = "class/classedit.html"
}
func (this *ClassController) ClassUpdateAction() {

	//当请求为post时，先初始化orm，再获得请求数据中的id但是不需要像展示页面一样获得数据，获得class结构体对象以后，直接获得修改的数据
	//将其重写class的属性，然后通过指定修数据库中的某些数据完成更新，更新完成以后需要跳转回list页面

	o := orm.NewOrm()

	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {
		this.StopRun()
	}
	class := models.Class{
		Id: id,
	}

	o.Read(&class)

	class.ClassName = this.Input().Get("className")

	flash := beego.NewFlash()

	valid := validation.Validation{}

	errResult, errRule := valid.Valid(&class)

	if nil != errRule {
		log.Println(errRule)
	}

	if !errResult {

		for _, errResult := range valid.Errors {

			flash.Set(errResult.Field, errResult.Message)

		}

		for k, v := range this.Input() {
			flash.Set(k, v[0])

		}

		flash.Store(&this.Controller)

		this.Redirect(this.Ctx.Request.Referer(), http.StatusFound)

		this.StopRun()
	}

	o.Update(&class)

	this.Redirect(beego.URLFor("ClassController.ClassList"), http.StatusFound)

	this.StopRun()
}
