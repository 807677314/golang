package main

import (
	"html/template"
	"log"
	"net/http"
	"session"
)

//首页的处理器
func IndexAction(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{}

	sessionData := session.Read(w, r)

	if !CheckSession(w, r, sessionData) {
		sessionData["Redirect"] = "/Index/"
		session.Write(w, r, sessionData)
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	data["name"] = sessionData["loginName"]

	//解析模板
	tpl, err := template.ParseFiles("./template/index/index.html", "./template/partial/top.html", "./template/partial/menu.html", "./template/partial/header.html")
	if nil != err {
		log.Println(err)
		return
	}

	tpl.Execute(w, data)

}

//静态资源的处理
func AssetAction(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./template/"+r.URL.Path)

}
