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

func ClassroomListAction(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{}

	

	tpl, err := template.ParseFiles("./template/classroom/classroomlist.html", "./template/partial/top.html", "./template/partial/menu.html", "./template/partial/header.html")
	if nil != err {
		log.Println(err)
		return
	}

	p := r.URL.Query().Get("page")
	if "" == p {
		p = "1"
	}
	page, err := strconv.ParseFloat(p, 64)
	if nil != err {
		log.Println(err)
		return
	}

	countstr, err := Dao.Table("classroom").Field("count(*)").FetchValue()
	if nil != err {
		log.Println(err)
		return
	}
	count, err := strconv.Atoi(countstr)
	if nil != err {
		log.Println(err)
		return
	}
	size := 5.0
	offset := (page - 1.0) * size
	pagecount := math.Ceil(float64(count) / size)

	ob := r.URL.Query().Get("ob")

	if "" == ob {
		ob = "DESC"
	}
	of := r.URL.Query().Get("of")

	if "" != of {
		log.Println(of)

		Dao.OrderBy(fmt.Sprintf("%s %s", of, ob))
	}

	keywords := r.URL.Query().Get("keywords")

	if "" != keywords {
		Dao.Where("classroomName like ? or classroomAdress like ?", keywords+"%", keywords+"%")
	}

	res, err := Dao.Table("classroom").Limit(strconv.Itoa(int(size))).Offset(strconv.Itoa(int(offset))).FetchRows()

	if nil != err {
		log.Println(err)
		return
	}

	data["keywords"] = keywords
	data["count"] = count
	data["page"] = int(page)
	data["prev"] = int(page) - 1
	data["next"] = int(page) + 1
	data["pagecount"] = int(pagecount)
	data["end"] = int(pagecount) - 1
	data["classroom"] = res
	data["ob"] = ob
	data["of"] = of
	tpl.Execute(w, data)

}

func ClassroomDelAction(w http.ResponseWriter, r *http.Request) {


	id := r.URL.Query().Get("id")

	_, err := Dao.Table("classroom").Where("id=?", id).Delete()

	if nil != err {
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/classroom/list/", 302)
}

func ClassroomUpdateAction(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{}


	

	tpl, err := template.ParseFiles("./template/classroom/classroomupdate.html", "./template/partial/top.html", "./template/partial/menu.html", "./template/partial/header.html")
	if nil != err {
		log.Println(err)
		return
	}

	id := r.URL.Query().Get("id")

	res, err := Dao.Table("classroom").Where("id=?", id).FetchRow()

	if nil != err {
		log.Println(err)
		return
	}

	data["classroom"] = res

	if r.Method == "POST" {

		r.ParseForm()
		c := r.PostForm
		content := map[string]interface{}{}
		for key, value := range c {
			content[key] = value[0]
		}
		if "" == id {
			content["update_at"] = time.Now().Format("2006-01-02 15:04:03")

			_, err = Dao.Table("classroom").Insert(content)

			if nil != err {
				log.Println(err)
				return
			}
		} else {
			content["create_at"] = time.Now().Format("2006-01-02 15:04:03")

			_, err = Dao.Table("classroom").Where("id=?", id).Update(content)
			if nil != err {
				log.Println(err)
				return
			}
			http.Redirect(w, r, "/classroom/list/", 302)
		}
	}

	tpl.Execute(w, data)

}

func ClassroomBatchDelAction(w http.ResponseWriter, r *http.Request) {



	r.ParseForm()

	ids := r.PostForm["ids"]
	for _, value := range ids {
		_, err := Dao.Table("classroom").Where("id=?", value).Delete()
		if nil != err {
			log.Println(err)
			return
		}
	}
	http.Redirect(w, r, "/classroom/list/", 302)
}
