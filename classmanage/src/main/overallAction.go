package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func OverallAction(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{}

	

	
	//解析模板
	tpl, err := template.ParseFiles("./template/overall/overalllist.html", "./template/partial/top.html", "./template/partial/menu.html", "./template/partial/header.html")
	if nil != err {
		log.Println(err)
		return
	}

	//查询overall表中的记录数，导入模板中让用户可以知道一共存在多少条数据
	c, err := Dao.Table("overall").Field("count(*)").FetchValue()
	if nil != err {
		log.Println(err)
		return
	}

	//利用strconv包中的ParseFloat 函数将获得记录数转换为float64类型，因为之后math.Ceil()向上取整时需要float64的参数
	count, err := strconv.ParseFloat(c, 64)
	if nil != err {
		log.Println(err)
		return
	}

	//获得URL中page的值，即用户当前浏览的页数，如果用户传入的值为负数或0时，将其重置为1，因为页数不能为负或者0
	p := r.URL.Query().Get("page")

	if "" == p {
		p = "1"
	}
	page, err := strconv.Atoi(p)
	if nil != err {
		log.Println(err)
		return
	}
	if page <= 0 {
		page = 1
	}

	//size代表每页显示几条数字，这里设置为浮点数，同样是因为之后math.Ceil()向上取整时需要float64的参数
	size := 5.0

	//pagecount代表全部数据一共需要显示几页
	pagecount := math.Ceil(count / size)

	//offset代表limit分页显示时，每一页的第一条数据从哪一条数据开始显示
	offset := (page - 1) * int(size)

	//调用Dao设置limit，需要string类型的参数，所以都需要转化为string类型，因为limit会返回一个*DAO,被Dao包里的结构体DAO接收，
	//所以这里不需要接收，下面依旧可以使用这里更新的数据
	Dao.Limit(strconv.Itoa(int(size))).Offset(strconv.Itoa(offset))

	//接收URL中ob的值，即用户需要的排序的是升序还是降序，当用户第一次点击到页面时，接收不到ob的值，即为空。这里默认设置为降序。
	ob := r.URL.Query().Get("ob")

	if "" == ob {
		ob = "DESC"
	}

	//接收URL中of的值，即用户需要以哪一个字段为基准排序，如果接收到的数据不为空，即进行排序，如果为空，则跳过。
	of := r.URL.Query().Get("of")

	if "" != of {

		Dao.OrderBy(fmt.Sprintf("%s %s", of, ob))
	}

	//现将*DAO里储存的数据更新，避免下面操作时Dao的操作过于长
	Dao.Field("o.id as overallid , c.id , r.id , t.id ,c.className , r.classroomName , t.teacherName , o.update_at").
		Table("overall").As("o")

	//获取URL中Keywords的值，如果Keywords不为空，即用户需要筛选数据，则进行筛选，如果为空，则跳过
	keywords := r.URL.Query().Get("keywords")

	if "" != keywords {
		Dao.Where("c.className like ? or r.classroomName like ? or t.teacherName like ?", keywords+"%", keywords+"%", keywords+"%")
	}

	//同理先将*DAO里储存的数据更新，避免下面操作时Dao的操作过于长
	Dao.LeftJoin("class", "c", "o.classid=c.id").
		LeftJoin("classroom", "r", "o.classroomid=r.id").
		LeftJoin("teacher", "t", "o.teacherid=t.id")

	//综合上面所有更新后进行查询
	res, err := Dao.FetchRows()

	if nil != err {
		log.Println(err)
		return
	}

	//将需要传给模板的值在这里都传入data中，用于展示
	data["overall"] = res
	data["keywords"] = keywords
	data["count"] = int(count)
	data["pagecount"] = int(pagecount)
	//当前页
	data["page"] = int(page)
	//上一页
	data["prev"] = int(page) - 1
	//下一页
	data["next"] = int(page) + 1
	//在模板展示时，最后一页我们只需要展示到最后一页为止，在模板中判定当下一页大于等于最大页数+1页时不会再显示下一页
	//当一共只有10页时，最大页数+1页为11页，当浏览到第10页时，下一页为11页，等于最大页数+1页，所以不会显示下一页
	data["end"] = int(pagecount) + 1
	data["ob"] = ob
	data["of"] = of
	tpl.Execute(w, data)

}

func OverallAddAction(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{}

	


	//解析模板
	tpl, err := template.ParseFiles("./template/overall/overalladd.html", "./template/partial/top.html", "./template/partial/menu.html", "./template/partial/header.html")
	if nil != err {
		log.Println(err)
		return
	}
	//当页面提交方式为POST时，才会进行以下操作，因为当用户第一次点击进入页面时，没有任何数据的提交，仅仅只需要展示页面
	if r.Method == "POST" {

		//解码数据
		r.ParseForm()
		c := r.PostForm
		content := map[string]interface{}{}

		//解码后的数据为，map[string][]interface{}型数据，需要遍历数据，然后获得value[0]的值，即用户传入的数据，当用户传入的数据不为"请选择"时，
		//才插入数据库中，避免用户没有进行选择时，点击了提交，导致传入的数据为“请选择”，误传入数据库中。
		for key, value := range c {

			if "请选择" != value[0] {

				content[key] = value[0]

			}

		}

		//更新数据的更新时间
		content["update_at"] = time.Now().Format("2006-01-02 15:04:03")

		//将用户传入的数据传入数据库中
		_, err := Dao.Table("overall").Insert(content)
		if nil != err {
			log.Println(err)
			return
		}
	}

	//将其他表的数据展示出来，供用户选择，展示为名字，实则上传为ID，上传的ID记录在overall表中，这样就实现了多对多的选择
	class, err := Dao.Table("class").Field("id,className").FetchRows()
	if nil != err {
		log.Println(err)
		return
	}
	classroom, err := Dao.Table("classroom").Field("id,classroomName").FetchRows()
	if nil != err {
		log.Println(err)
		return
	}
	teacher, err := Dao.Table("teacher").Field("id,teacherName").FetchRows()
	if nil != err {
		log.Println(err)
		return
	}

	data["class"] = class
	data["classroom"] = classroom
	data["teacher"] = teacher
	tpl.Execute(w, data)

}

func OverallDelAction(w http.ResponseWriter, r *http.Request) {

	

	//获得URL中ID的值
	id := r.URL.Query().Get("id")

	//通过ID的值删除相应的数据
	_, err := Dao.Table("overall").Where("id=?", id).Delete()

	if nil != err {
		log.Println(err)
		return
	}
	//删除完毕后跳转回页面
	http.Redirect(w, r, "/overall/list/", 302)
}

func OverallUpdateAction(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{}

	
	

	//解析模板
	tpl, err := template.ParseFiles("./template/overall/overallupdate.html", "./template/partial/top.html", "./template/partial/menu.html", "./template/partial/header.html")
	if nil != err {
		log.Println(err)
		return
	}

	//当页面提交方式为POST时，才会进行以下操作，因为当用户第一次点击进入页面时，没有任何数据的提交，仅仅只需要展示页面
	if r.Method == "POST" {

		//获得URL中ID的值
		id := r.URL.Query().Get("id")

		//解码数据
		r.ParseForm()
		c := r.PostForm

		//解码后的数据为，map[string][]interface{}型数据，需要遍历数据，然后获得value[0]的值，即用户传入的数据
		content := map[string]interface{}{}
		for key, value := range c {

			content[key] = value[0]
		}

		//更新数据的更新时间
		content["update_at"] = time.Now().Format("2006-01-02 15:04:03")

		//根据ID将用户传入的数据在数据库中进行更新
		_, err := Dao.Table("overall").Where("id=?", id).Update(content)
		if nil != err {
			log.Println(err)
			return
		}
		//更新完毕后跳转回页面
		http.Redirect(w, r, "/overall/list/", 302)
	}

	//将其他表的数据展示出来，供用户选择，展示为名字，实则上传为ID，上传的ID记录在overall表中，这样就实现了多对多的选择
	class, err := Dao.Table("class").Field("id,className").FetchRows()
	if nil != err {
		log.Println(err)
		return
	}
	classroom, err := Dao.Table("classroom").Field("id,classroomName").FetchRows()
	if nil != err {
		log.Println(err)
		return
	}
	teacher, err := Dao.Table("teacher").Field("id,teacherName").FetchRows()
	if nil != err {
		log.Println(err)
		return
	}
	id := r.URL.Query().Get("id")
	overall, err := Dao.Table("overall").Where("id=?", id).FetchRow()
	if nil != err {
		log.Println(err)
		return
	}
	data["class"] = class
	data["classroom"] = classroom
	data["teacher"] = teacher
	data["overall"] = overall
	tpl.Execute(w, data)

}

func OverallBatchDelAction(w http.ResponseWriter, r *http.Request) {



	//解码数据，获得的IDS的为用户选择的需要删除的数据ID的集合，遍历ids时，每一次获得id都通过ID将数据删除，但是这个方法有个弊端，就是
	//这样删除需要频繁的调用数据库连接池和频繁的使用sql语句，导致速度下降
	r.ParseForm()
	ids := r.PostForm["ids"]
	for _, v := range ids {
		_, err := Dao.Table("overall").Where("id=?", v).Delete()
		if nil != err {
			log.Println(err)
			return
		}
	}
	//删除完毕后跳转回页面
	http.Redirect(w, r, "/overall/list/", 302)
}
