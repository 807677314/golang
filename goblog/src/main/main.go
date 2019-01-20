package main

import (
	"fmt"
	"net/http"
)

func main() {

	//解析ymal配置文件
	configparse()
	Dataanalyse()
	for name, c := range Categories {
		fmt.Println(name, c.Datacounts, c.Datalist)
	}

	http.HandleFunc("/", Index)
	http.HandleFunc("/category/", CategoryActive)

	http.HandleFunc("/css/", Asset)
	http.HandleFunc("/js/", Asset)
	http.HandleFunc("/fonts/", Asset)
	http.HandleFunc("/img/", Asset)

	//持续监听，如果没有错误，就为空
	http.ListenAndServe(":8088", nil)

}
