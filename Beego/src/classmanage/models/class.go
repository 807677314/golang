package models

import (
	"github.com/astaxie/beego/validation"
	"time"

	"github.com/astaxie/beego/orm"
)

type Class struct {
	Id        int       `orm:"pk;auto"`
	ClassName string    `valid:"Required"orm:"default('')"`
	ClassDesc string    `valid:"Required"orm:"default('')"`
	DeleteAt  time.Time `orm:"type(datetime);null;form:"-""`
	CreateAt  time.Time `orm:"type(datetime);auto_now_add"form:"-"`
	UpdateAt  time.Time `orm:"type(datetime);auto_now"form:"-"`
}

func (c *Class) TableIndex() [][]string {

	return [][]string{
		[]string{"ClassDesc"},
	}

}

func (c *Class) TableUnique() [][]string {

	return [][]string{
		[]string{"ClassName"},
	}

}

func (c *Class) Valid(v *validation.Validation) {

	class := Class{
		ClassName: c.ClassName,
	}

	o := orm.NewOrm()

	err := o.Read(&class, "ClassName")

	if nil == err {
		v.SetError("ClassName", "ClassName can not repeat")
	}

}

func init() {
	orm.RegisterModel(new(Class))
}
