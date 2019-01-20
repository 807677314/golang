package main

import (
	"dao"
	"fmt"
	"log"
	"net/http"
)

//设置全局变量，方便之后需要调用连接池的时候不必重复写，直接调用即可
var Dao *dao.DAO

func InitDAO() {

	//打开连接池
	d, err := dao.New("root:ld123456@tcp(127.0.0.1:3306)/classmanage?charset=utf8mb4")

	if nil != err {
		log.Println(err)
		return
	}
	Dao = d
}

func CheckSession(w http.ResponseWriter, r *http.Request, sessionData map[string]interface{}) bool {

	_, exists := sessionData["loginName"]

	if !exists {

		id, err := r.Cookie("rememberID")
		if nil != err {
			log.Println(err)
			return false
		}
		pwd, err := r.Cookie("rememberPwd")
		if nil != err {
			log.Println(err)
			return false
		}

		if "" != id.Value && "" != pwd.Value {

			salt := "salt"

			row, err := Dao.Table("admin").Where(fmt.Sprintf("md5(concat(id,'%s'))=? and md5(concat(password,'%s'))=?", salt, salt), id.Value, pwd.Value).FetchRow()

			if nil != err {
				log.Println(err)
				return false
			}

			if len(row) > 0 {

				sessionData["loginName"] = row["name"]

				return true
			}

		}

		return false
	}

	return true

}
