package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//解析文件Index

	//文件解析完毕，执行，回应HTTP的请求，没有错误的话就是nil

	tep, err := template.ParseFiles("./html/classic/index.html")

	if err != nil {
		log.Print(err)
		return

	}

	//将Markdown文件内容和配置文件内容整合在一起传到index.html中
	Sumdatas := map[interface{}]interface{}{"Datas": Datas, "Configs": Configs, "Categories": Categories}

	tep.Execute(w, Sumdatas)
}
func CategoryActive(w http.ResponseWriter, r *http.Request) {

	pattern := `^/category/(.*?)/?$`

	reg, err := regexp.Compile(pattern)

	if err != nil {
		return
	}

	result := reg.FindAllStringSubmatch(r.URL.Path, -1)

	CategoryName := result[0][1] //问题出现在CategoryName,循环出现，value出现带no value，一次不带no value

	fmt.Println(result[0][0])

	Datas := Categories[CategoryName].Datalist

	tep, err := template.ParseFiles("./html/classic/index.html")

	if err != nil {

		return

	}

	//将Markdown文件内容和配置文件内容整合在一起传到index.html中
	Sumdatas := map[interface{}]interface{}{"Datas": Datas, "Configs": Configs, "Categories": Categories}

	tep.Execute(w, Sumdatas)
}

func Asset(w http.ResponseWriter, r *http.Request) {

	//将剩下的功能性文件，CSS JS导入。
	http.ServeFile(w, r, "/html/classic"+r.URL.Path)
}
