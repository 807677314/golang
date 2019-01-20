package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Teacher struct {
	Id          int       `orm:"pk;auto"`
	TeacherName string    `valid:"Required"orm:"default()"`
	Gender      string    `valid:"Required"orm:"default(select)"`
	DeleteAt    time.Time `orm:"type(datetime);null;"`
	CreateAt    time.Time `orm:"type(datetime);auto_now_add"form:"-"`
	UpdateAt    time.Time `orm:"type(datetime);auto_now"form:"-"`
}

func (c *Teacher) TableIndex() [][]string {

	return [][]string{
		[]string{"TeacherName"},
	}

}

func init() {

	orm.RegisterModel(new(Teacher))
}
