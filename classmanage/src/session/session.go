package session

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Read(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	rand.Seed(time.Now().Unix())

	if rand.Intn(100) > 50 {
		GC()
	}
	content := map[string]interface{}{}

	key := initsessionID(w, r)

	var unmarshalcontent string

	readcontent, err := ioutil.ReadFile("./temp/" + key)

	if nil == err {

		unmarshalcontent = string(readcontent)

	}

	json.Unmarshal([]byte(unmarshalcontent), &content)

	return content

}

func Write(w http.ResponseWriter, r *http.Request, sessionDate map[string]interface{}) error {

	key := initsessionID(w, r)

	st, err := json.Marshal(sessionDate)

	if nil == err {
		err = ioutil.WriteFile("./temp/"+key, st, 0755)
	}

	return err

}

func GC() {

	infos, err := ioutil.ReadDir("./temp/")

	if nil != err {
		log.Println(err)
		return
	}

	for _, value := range infos {

		if time.Now().Unix()-value.ModTime().Unix() >= 60 {
			os.Remove("./temp/" + value.Name())
		}
	}

}

func initsessionID(w http.ResponseWriter, r *http.Request) string {

	var key string

	cookie, err := r.Cookie("key")

	if nil != err {

		rand.Seed(time.Now().Unix())
		key = fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(rand.Int()))))

		uid_cookies := &http.Cookie{
			Name:  "key",
			Value: key,
			Path:  "/",
		}

		http.SetCookie(w, uid_cookies)

	} else {

		key = cookie.Value

	}

	return key

}
