package data

import (
	"errors"
	"html/template"
	"io/ioutil"
	"regexp"

	blackfriday "gopkg.in/russross/blackfriday.v2"
	yaml "gopkg.in/yaml.v2"
)

type Data struct {
	Path     string
	Title    string
	Content  template.HTML
	Category string
	Tags     []string
	Author   string
	Date     string
	Img      string
}

//构造获得markdown文件内容的方法
func Newdata(path string) (*Data, error) {
	//快速读取全部文件内容
	content, err := ioutil.ReadFile(path)

	if err != nil {

		return nil, errors.New("文件读取错误")
	}
	//正则表达式将-------之间的内容分割出来
	//(.*?)这么处理是因为将其更改为非贪婪模式，保证需要提取内容不会匹配到上下用于分割作用的---
	pattern := "(?s)^-{6,}(.*?)-{6,}\\s*(.*)$"

	reg, err1 := regexp.Compile(pattern)

	if err1 != nil {

		return nil, errors.New("文件解析失败")
	}
	//在content文件中把需要的内容挑选出来，得到的是[][]string类型的数据，filereg[0][1]是需要的元数据，filereg[0][2]是网页需要解析的文档内容
	result := reg.FindAllStringSubmatch(string(content), -1)

	//声明一map类型的变量，key和value都是接口类型
	mateinfo := make(map[interface{}]interface{})

	//解析获得元文件内容
	err2 := yaml.Unmarshal([]byte(result[0][1]), &mateinfo)
	if err2 != nil {
		return nil, errors.New("文件解析错误")
	}

	//通过blackfriday将markdown文件转换为http语言文件
	filecontent := blackfriday.Run([]byte(result[0][2]))

	//通过类型断言将interface{}类型转换为我们需要的类型
	t, _ := mateinfo["title"].(string)
	a, _ := mateinfo["author"].(string)
	d, _ := mateinfo["date"].(string)
	i, _ := mateinfo["img"].(string)
	cat, _ := mateinfo["category"].(string)

	//声明一个与Tags一样的变量类型，为[]string类型
	var ta []string

	//遍历meteinfo["tags"],将得到的[]interface{}类型的v通过断言转换为string类型
	//mateinfo["tags"]本来的类型是[]interface{},不能直接转换为[]string类型，但也不能遍历，所以需要强制再转换为[]interface{}类型
	for _, v := range mateinfo["tags"].([]interface{}) {
		ta = append(ta, v.(string))
	}

	//检查结构体是否存在
	//当path不存在时，Path所得到的值为空，没有意义，只有当path存在时才能最终为p赋值，使判定存在意义。
	P := &Data{
		Path:     path,
		Content:  template.HTML(string(filecontent)),
		Title:    t,
		Author:   a,
		Date:     d,
		Img:      i,
		Category: cat,
		Tags:     ta,
	}

	// fmt.Println(template.HTML(string(filecontent)))

	return P, nil

}
