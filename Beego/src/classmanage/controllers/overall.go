package controllers

import (
	"classmanage/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type OverallController struct {
	//继承beego.Controller的全部属性
	beego.Controller
}

func (this *OverallController) OverallList() {

	o := orm.NewOrm()

	overall := []*models.Overall{}

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

	cnt, err := o.QueryTable("overall").Count()

	var size int = 5

	offset := (page - 1) * size

	pagecount := math.Ceil(float64(cnt) / float64(size))

	qs := o.QueryTable("overall").RelatedSel().Limit(size, offset)

	qs.All(&overall)

	this.Data["count"] = cnt
	this.Data["page"] = int(page)
	this.Data["prev"] = int(page) - 1
	this.Data["next"] = int(page) + 1
	this.Data["pagecount"] = int(pagecount)
	this.Data["over"] = int(pagecount) + 1

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	this.Data["overall"] = overall

	this.Layout = "layout/layout.html"
	this.TplName = "overall/overalllist.html"

}

func (this *OverallController) OverallDelete() {

	o := orm.NewOrm()

	//获得url中的传递的id值，如果发生错误就终止程序运行
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {

		this.StopRun()

	}

	//删除值为id的数据
	overall := models.Overall{
		Id: id,
	}

	o.Delete(&overall)

	//制作URL给模板使用
	beego.URLFor("OverallController.OverallDelete")

	//删除完成后跳转回list界面，然后return
	this.Redirect(beego.URLFor("OverallController.OverallList"), http.StatusFound)

	this.StopRun()
}

func (this *OverallController) OverallAddList() {

	//当请求为get时，仅仅需要展示页面

	o := orm.NewOrm()
	class := []*models.Class{}
	classroom := []*models.Classroom{}
	teacher := []*models.Teacher{}

	o.QueryTable("class").All(&class)
	o.QueryTable("classroom").All(&classroom)
	o.QueryTable("teacher").All(&teacher)

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	this.Data["class"] = class
	this.Data["classroom"] = classroom
	this.Data["teacher"] = teacher

	beego.URLFor("OverallController.OverallAddList")

	this.Layout = "layout/layout.html"
	this.TplName = "overall/overalledit.html"
}

func (this *OverallController) OverallAddAction() {

	//当请求为post时，需要处理数据，先初始化orm，再获得class对象，再讲浏览器请求的数据获得添加到class属性中，插入数据库中，再跳转回list页面
	o := orm.NewOrm()

	classid, err := strconv.Atoi(this.Input().Get("classid"))

	if nil != err {

		this.Redirect(beego.URLFor("OverallController.OverallAddList"), http.StatusFound)
		this.StopRun()
	}
	classroomid, err := strconv.Atoi(this.Input().Get("classroomid"))

	if nil != err {

		this.Redirect(beego.URLFor("OverallController.OverallAddList"), http.StatusFound)
		this.StopRun()
	}
	teacherid, err := strconv.Atoi(this.Input().Get("teacherid"))

	if nil != err {

		this.Redirect(beego.URLFor("OverallController.OverallAddList"), http.StatusFound)
		this.StopRun()
	}

	overall := models.Overall{
		Class:     &models.Class{Id: classid},
		Classroom: &models.Classroom{Id: classroomid},
		Teacher:   &models.Teacher{Id: teacherid},
		UpdateAt:  time.Now(),
	}

	o.Insert(&overall)

	this.Redirect(beego.URLFor("OverallController.OverallList"), http.StatusFound)

	this.StopRun()

}

func (this *OverallController) OverallUpdateList() {
	//当请求为get时，仅仅需要展示页面，但是还需要展示数据便于修改。
	//先初始化orm，再通过请求的id获得数据库中的数据，将其展示到模板。

	o := orm.NewOrm()

	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {
		this.StopRun()
	}

	overall := models.Overall{
		Id: id,
	}

	o.Read(&overall)

	class := []*models.Class{}
	classroom := []*models.Classroom{}
	teacher := []*models.Teacher{}

	o.QueryTable("class").All(&class)
	o.QueryTable("classroom").All(&classroom)
	o.QueryTable("teacher").All(&teacher)

	username := GetSessionName(this.Ctx)
	this.Data["username"] = username

	this.Data["class"] = class
	this.Data["classroom"] = classroom
	this.Data["teacher"] = teacher

	this.Data["overallinfo"] = overall
	beego.URLFor("OverallController.OverallUpdateList")

	this.Layout = "layout/layout.html"
	this.TplName = "overall/overalledit.html"
}

func (this *OverallController) OverallUpdateAction() {

	o := orm.NewOrm()

	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if nil != err {
		this.StopRun()
	}

	overall := models.Overall{

		Id: id,
	}

	o.Read(&overall)

	classid, err := strconv.Atoi(this.Input().Get("classid"))

	if nil != err {

		this.Redirect(beego.URLFor("OverallController.OverallAddList"), http.StatusFound)
		this.StopRun()
	}
	classroomid, err := strconv.Atoi(this.Input().Get("classroomid"))

	if nil != err {

		this.Redirect(beego.URLFor("OverallController.OverallAddList"), http.StatusFound)
		this.StopRun()
	}
	teacherid, err := strconv.Atoi(this.Input().Get("teacherid"))

	if nil != err {

		this.Redirect(beego.URLFor("OverallController.OverallAddList"), http.StatusFound)
		this.StopRun()
	}

	overall.Class = &models.Class{Id: classid}
	overall.Classroom = &models.Classroom{Id: classroomid}
	overall.Teacher = &models.Teacher{Id: teacherid}

	o.Update(&overall)

	this.Redirect(beego.URLFor("OverallController.OverallList"), http.StatusFound)

	this.StopRun()

}
