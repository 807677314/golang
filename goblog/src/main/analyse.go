package main

import (
	"data"
	"fmt"
	"io/ioutil"
	"log"
)

type Datacategory struct {
	Name       string
	Datacounts int
	Datalist   []*data.Data
}

func Newdatacategory(n string) *Datacategory {

	return &Datacategory{

		Name:       n,
		Datacounts: 0,
		Datalist:   []*data.Data{},
	}
}

var Categories = make(map[string]*Datacategory)

//声明一个结构体切片类型的变量
var Datas []*data.Data

func Dataanalyse() {

	//	遍历data文件夹
	infos, err := ioutil.ReadDir("./data")

	if err != nil {
		log.Print(err)

	}

	//遍历data内所有的markdown文件
	for _, info := range infos {

		fmt.Println("info:", info.Name())
		//把makedown文件传给Newdata函数处理
		details, err := data.Newdata("./data/" + info.Name())

		if err != nil {

			log.Print(err)

		}
		//把每一个结构体整合
		Datas = append(Datas, details)

		category, exists := Categories[details.Category]

		if exists {

			category.Datacounts++

		} else {

			category = Newdatacategory(details.Category)
			Categories[details.Category] = category
			category.Datacounts++

		}

		category.Datalist = append(category.Datalist, details)
	}

	// fmt.Println(Categories)
}
