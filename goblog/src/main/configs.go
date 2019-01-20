package main

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//声明一个map类型的变量，Key和value都是接口interface{}
var Configs map[interface{}]interface{}

//解析配置文件
func configparse() error {

	//读取
	info, err := ioutil.ReadFile("./config/site.yaml")

	if err != nil {

		return errors.New("文件读取失败")

	}

	//通过yaml.v2整理的解析出来的内容，传入Configs中
	err = yaml.Unmarshal(info, &Configs)

	if err != nil {
		return errors.New("文件解析失败")
	}

	return nil
}
