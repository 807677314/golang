package wechat

import (
	"beego_grpc/initial"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/esap/wechat" // 微信SDK包
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type menu struct {
	Button map[string][]sub_button `json:"menu"`
}

type sub_button struct {
	Typ       string            `json:"type,omitempty"`
	Name      string            `json:"name,omitempty"`
	Key       string            `json:"key,omitempty"`
	Url       string            `json:"url,omitempty"`
	SubButton []send_sub_button `json:"sub_button"`
}

type revice_button struct {
	Button []revice_sub_button `json:"button"`
}

type revice_sub_button struct {
	Id        int                 `json:"id,omitempty"`
	Pid       int                 `json:"pid,omitempty"`
	MenuLevel int                 `json:"menu_level,omitempty"`
	Typ       string              `json:"type,omitempty"`
	Name      string              `json:"name,omitempty"`
	Key       string              `json:"key,omitempty"`
	Url       string              `json:"url,omitempty"`
	SubButton []revice_sub_button `json:"sub_button"`
}

type send_button struct {
	Button []send_sub_button `json:"button"`
}

type send_sub_button struct {
	Id        int               `json:"-"`
	Pid       int               `json:"-"`
	MenuLevel int               `json:"-"`
	Typ       string            `json:"type,omitempty"`
	Name      string            `json:"name,omitempty"`
	Key       string            `json:"key,omitempty"`
	Url       string            `json:"url,omitempty"`
	SubButton []send_sub_button `json:"sub_button"`
}

type WechatControllers struct {
	beego.Controller
}

func GetACSkey() string {
	return wechat.GetAccessToken()
}

func GetJson(w http.ResponseWriter, r *http.Request) revice_sub_button {

	r.ParseForm()
	subbutton := revice_sub_button{}

	body, err := ioutil.ReadAll(r.Body)

	if nil != err {
		log.Println(err)
	}

	err = json.Unmarshal(body, &subbutton)

	if nil != err {
		log.Println(err)
	}
	return subbutton

}

func GetMenu() {
	var Id_1 = 100
	var Id_2 = 200

	key := GetACSkey()

	uri := "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=" + key

	resp, err := http.Get(uri)
	if nil != err {
		log.Println(err)
	}

	resBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	menu := menu{}
	json.Unmarshal(resBytes, &menu)
	reviceButton := revice_button{}
	a := menu.Button["button"]
	for k, v := range a {
		if nil == reviceButton.Button {
			reviceButton.Button = make([]revice_sub_button, 3)
		}
		reviceButton.Button[k].Key = v.Key
		reviceButton.Button[k].Name = v.Name
		reviceButton.Button[k].Typ = v.Typ
		reviceButton.Button[k].Url = v.Url
		reviceButton.Button[k].Id = Id_1 + k
		reviceButton.Button[k].MenuLevel = 0
		for key, value := range a[k].SubButton {
			if nil == reviceButton.Button[k].SubButton {
				reviceButton.Button[k].SubButton = make([]revice_sub_button, 5)
			}
			reviceButton.Button[k].SubButton[key].Typ = value.Typ
			reviceButton.Button[k].SubButton[key].Name = value.Name
			reviceButton.Button[k].SubButton[key].Key = value.Key
			reviceButton.Button[k].SubButton[key].Url = value.Url
			reviceButton.Button[k].SubButton[key].Pid = reviceButton.Button[k].Id
			reviceButton.Button[k].SubButton[key].Id = Id_2 + key
			reviceButton.Button[k].SubButton[key].MenuLevel = 1
		}
	}

	revBtn := revice_button{}

	for k, v := range reviceButton.Button {
		keys := []int{}
		if 0 != v.Id {
			revBtn.Button = append(revBtn.Button, reviceButton.Button[k])

		}

		for key, value := range reviceButton.Button[k].SubButton {
			if 0 == value.Id {
				keys = append(keys, key)
			}
		}

		lenth := len(keys)
		if lenth > 2 {
			revBtn.Button[k].SubButton = append(append([]revice_sub_button{}, revBtn.Button[k].SubButton[:keys[0]]...), revBtn.Button[k].SubButton[keys[lenth-1]+1:]...)
		} else if lenth == 1 {
			revBtn.Button[k].SubButton = append(append([]revice_sub_button{}, revBtn.Button[k].SubButton[:0]...), revBtn.Button[k].SubButton[1:]...)
		}
	}

	initial.Bm.Put("revBtn", revBtn, 600*time.Second)

}

func sendMsg(res revice_button, uri string) []byte {
	sendcontent := send_button{}

	bytesMenu, _ := json.Marshal(res)
	json.Unmarshal(bytesMenu, &sendcontent)

	key := GetACSkey()

	uploaduri := uri + key
	bytesData, err := json.Marshal(sendcontent)
	if nil != err {
		log.Fatalf("json编码失败：%s", err)
	}

	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", uploaduri, reader)
	if err != nil {
		log.Fatalf("建立请求失败：%s", err)
		return nil
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalf("请求发送失败：%s", err)
		return nil
	}
	resBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return resBytes
}

var Menu revice_button
var Id1 int = 100
var Id2 int = 200

//创建菜单
//{"name":"淘宝","url":"www.taobao.com","key":"shopping","type":"click","menu_level":0,"pid":101}
//创建二级菜单必须需要pid
func (c *WechatControllers) MenuCreate() {

	data := GetJson(c.Ctx.ResponseWriter, c.Ctx.Request)
	GetMenu()
	result := initial.Bm.Get("revBtn")
	Menu = result.(revice_button)

	if 0 == data.MenuLevel {
		if 0 == len(Menu.Button) {
			btn1 := revice_sub_button{}
			btn1 = data
			btn1.Id = Id1
			Menu.Button = append(Menu.Button, btn1)
		} else if len(Menu.Button) < 3 {
			a := 0
			for _, v := range Menu.Button {
				if data.Id == v.Id {
					a++
				}
			}

			if 0 == a {
				btn2 := revice_sub_button{}
				btn2 = data
				Id1++
				btn2.Id = Id1
				Menu.Button = append(Menu.Button, btn2)
			}
		}
	} else if 1 == data.MenuLevel {
		a := 0
		for k := range Menu.Button {
			if Menu.Button[k].Id == data.Pid {
				a++
			}
		}

		if 0 == a {
			log.Println("先创建上级菜单")
			c.StopRun()
		}
		b := 1
		for _, v := range Menu.Button[data.Pid-100].SubButton {
			if v.Id == data.Id {
				b++
			}
		}
		if 1 == a && 1 == b && len(Menu.Button[data.Pid-100].SubButton) < 5 {
			btn := revice_sub_button{}
			btn = data
			btn.Id = Id2
			Id2++
			Menu.Button[data.Pid-100].SubButton = append(Menu.Button[data.Pid-100].SubButton, btn)
		} else if 1 == a && len(Menu.Button[data.Pid-100].SubButton) < 5 {
			btn := revice_sub_button{}
			btn = data
			btn.Id = Id2
			Menu.Button[data.Pid-100].SubButton[data.Pid-100] = data
		}

	}
	uri := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="

	resBytes := sendMsg(Menu, uri)
	c.Ctx.ResponseWriter.Header().Set("Content-Length", strconv.Itoa(len(resBytes)))
	c.Ctx.ResponseWriter.Write(resBytes)
	c.StopRun()

}

//查询菜单
func (c *WechatControllers) MenuQuery() {

	GetMenu()
	result := initial.Bm.Get("revBtn")
	resBytes, err := json.Marshal(result)
	if nil != err {
		log.Println(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Length", strconv.Itoa(len(resBytes)))
	c.Ctx.ResponseWriter.Write(resBytes)
	c.StopRun()

}

//删除菜单
//{"id":101,"pid":101}
//删除一级菜单可以不需要pid，删除二级菜单必要pid
func (c *WechatControllers) MenuDelete() {

	GetMenu()
	data := GetJson(c.Ctx.ResponseWriter, c.Ctx.Request)
	result := initial.Bm.Get("revBtn")
	res := revice_button{}
	res = result.(revice_button)
	for k, v := range res.Button {
		if v.Id == data.Id {
			res.Button[k].Id = 0
			break
		} else {
			for key, value := range res.Button[k].SubButton {
				if value.Id == data.Id && value.Pid == data.Pid {
					res.Button[k].SubButton[key].Id = 0
					break
				}
			}
		}
	}

	revBtn := revice_button{}
	for k, v := range res.Button {
		keys := []int{}
		if 0 != v.Id {
			revBtn.Button = append(revBtn.Button, res.Button[k])
		} else {
			for key, value := range res.Button[k].SubButton {

				if 0 == value.Id {
					keys = append(keys, key)
				}

				lenth := len(keys)
				if lenth > 2 {
					revBtn.Button[k].SubButton = append(append([]revice_sub_button{}, revBtn.Button[k].SubButton[:keys[0]]...), revBtn.Button[k].SubButton[keys[lenth-1]+1:]...)
				} else if lenth == 1 {
					revBtn.Button[k].SubButton = append(append([]revice_sub_button{}, revBtn.Button[k].SubButton[:0]...), revBtn.Button[k].SubButton[1:]...)
				}
			}

		}

	}
	var uri string
	if nil == revBtn.Button {
		uri = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="
	} else {
		uri = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="
	}

	resBytes := sendMsg(revBtn, uri)
	c.Ctx.ResponseWriter.Header().Set("Content-Length", strconv.Itoa(len(resBytes)))
	c.Ctx.ResponseWriter.Write(resBytes)
	c.StopRun()

}

//修改菜单
//{"id":100,"name":"淘宝","url":"www.taobao.com","key":"shopping","type":"click","menu_level":0,"pid":101}
//创建二级菜单必须需要pid
func (c *WechatControllers) MenuUpdate() {
	GetMenu()
	data := GetJson(c.Ctx.ResponseWriter, c.Ctx.Request)
	result := initial.Bm.Get("revBtn")
	res := revice_button{}

	res = result.(revice_button)

	for k, v := range res.Button {

		if v.Id == data.Id {
			res.Button[k].Key = data.Key
			res.Button[k].Name = data.Name
			res.Button[k].Typ = data.Typ
			res.Button[k].Url = data.Url
			break
		} else {
			for key, value := range res.Button[k].SubButton {
				if value.Id == data.Id && value.Pid == data.Pid {
					res.Button[k].SubButton[key].Key = data.Key
					res.Button[k].SubButton[key].Name = data.Name
					res.Button[k].SubButton[key].Typ = data.Typ
					res.Button[k].SubButton[key].Url = data.Url
					break
				}
			}
		}
	}
	uri := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="
	resBytes := sendMsg(res, uri)
	c.Ctx.ResponseWriter.Header().Set("Content-Length", strconv.Itoa(len(resBytes)))
	c.Ctx.ResponseWriter.Write(resBytes)
	c.StopRun()
}
