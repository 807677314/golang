package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"session"
	"strings"
	"time"
)

func LoginAction(w http.ResponseWriter, r *http.Request) {

	//判断数据上传的方式是否为post，如果不是则仅仅展示页面，如果是则执行校验工作
	if r.Method == "POST" {

		sessionData := session.Read(w, r)

		data, exists := sessionData["code"]

		if !exists || r.FormValue("captcha") == "" || strings.ToLower(data.(string)) != strings.ToLower(r.FormValue("captcha")) {
			http.Redirect(w, r, "/login/", 302)
			return
		}

		//解析数据获得提交的user和password
		r.ParseForm()
		user := r.PostForm.Get("user")
		password := r.PostForm.Get("password")
		remember := r.PostForm.Get("rememberme")

		//再数据库中查询相匹配的user和password,如果存在则说明用户存在，不存在则跳转会登录页面
		row, err := Dao.Table("admin").Where("user = ? and password = MD5(?)", user, password).FetchRow()

		if nil != err {
			log.Println(err)
		}

		if len(row) > 0 {

			if "1" == remember {
				salt := "salt"
				http.SetCookie(w, &http.Cookie{
					Name:    "rememberID",
					Value:   fmt.Sprintf("%x", md5.Sum([]byte(row["id"]+salt))),
					Expires: time.Now().Add(30 * 24 * 3600 * time.Second),
					Path:    "/",
				})
				http.SetCookie(w, &http.Cookie{
					Name:    "rememberPwd",
					Value:   fmt.Sprintf("%x", md5.Sum([]byte(row["password"]+salt))),
					Expires: time.Now().Add(30 * 24 * 3600 * time.Second),
					Path:    "/",
				})

			}

			sessionData := session.Read(w, r)

			sessionData["loginName"] = row["name"]

			loginRedirect := "/Index/"

			redirect, exists := sessionData["Redirect"]

			if exists {

				loginRedirect = redirect.(string)
				delete(sessionData, "Redirect")

			}

			session.Write(w, r, sessionData)
			http.Redirect(w, r, loginRedirect, 302)

			return
		} else {

			http.Redirect(w, r, "/login/", 302)
			return
		}
	} else {

		//解析模板
		tpl, err := template.ParseFiles("./template/index/login.html")

		if nil != err {
			log.Println(err)
		}
		data := map[string]interface{}{}
		tpl.Execute(w, data)
	}

}

func SignupAction(w http.ResponseWriter, r *http.Request) {

	//解析模板
	tpl, err := template.ParseFiles("./template/index/signup.html")

	if nil != err {
		log.Println(err)
	}

	//如果用户未填写提交数据时，即第一次点击进入页面时，仅仅展示页面
	if r.Method == "POST" {

		r.ParseForm()
		user := r.PostForm.Get("user")
		password := r.PostForm.Get("password")
		password_repeat := r.PostForm.Get("password_repeat")
		name := r.PostForm.Get("name")

		////判断用户名和昵称不能为空，如为空，则跳转回注册页面，不需要验证密码不为空，因为在html里设置了最小密码长度为6位
		if user == "" || name == "" {
			http.Redirect(w, r, "/signup/", 302)
			return
		}

		//判断2次输入的密码是否相等，如不相等，则跳转回注册页面
		if password != password_repeat {
			http.Redirect(w, r, "/signup/", 302)
			return
		}

		data := map[string]interface{}{}
		data["user"] = user

		//为password进行md5加密，加密以后返回值为16个值的字节切片，需要将其以16进制格式化输入
		data["password"] = fmt.Sprintf("%x", (md5.Sum([]byte(password))))
		data["create_at"] = time.Now().Format("2006-01-02 15:03:05")
		data["name"] = name

		//插入数据库中
		row, err := Dao.Table("admin").Insert(data)
		if nil != err {
			log.Println(err)
		}

		//当row=0的时候，说明插入未成功，跳转会注册页面
		if row != 0 {

			http.Redirect(w, r, "/login/", 302)
		}

	}

	tpl.Execute(w, nil)
}

func LogoutAction(w http.ResponseWriter, r *http.Request) {

	sessionData := session.Read(w, r)

	delete(sessionData, "loginName")

	session.Write(w, r, sessionData)

	http.SetCookie(w, &http.Cookie{
		Name:    "rememberID",
		Value:   "1123",
		Expires: time.Now(),
		Path:    "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "rememberPwd",
		Value:   "1123",
		Expires: time.Now(),
		Path:    "/",
	})

	http.Redirect(w, r, "/login/", 302)

}
