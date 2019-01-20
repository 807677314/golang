package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Userinfo struct {
	Id       int       `orm:"pk;auto"`
	User     string    `orm:"on_delete(cascade)`
	UserName string    `orm:"on_delete(cascade)"`
	Password string    `orm:"on_delete(cascade);"`
	DeleteAt time.Time `orm:"type(datetime);null;"form:"-"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"form:"-"`
	Updated  time.Time `orm:"auto_now;type(datetime)"form:"-"`

}

func (u *Userinfo) TableUnique() [][]string {

	return [][]string{
		[]string{"User", "Password"},
		[]string{"UserName"},
	}

}

func init() {
	orm.RegisterModel(new(Userinfo))
}

