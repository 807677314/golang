package models

import (
	"github.com/astaxie/beego/validation"
	"time"

	"github.com/astaxie/beego/orm"
)

type Classroom struct {
	Id              int       `orm:"pk;auto"`
	ClassroomName   string    `valid:"Required"orm:"default('')"`
	ClassroomAdress string    `valid:"Required"orm:"default('')"`
	DeleteAt        time.Time `orm:"type(datetime);null;"form:"-"`
	CreateAt        time.Time `orm:"type(datetime);auto_now_add"form:"-"`
	UpdateAt        time.Time `orm:"type(datetime);auto_now"form:"-"`
}

func (c *Classroom) TableUnique() [][]string {

	return [][]string{
		[]string{"ClassroomAdress", "ClassroomName"},
	}

}

func (c *Classroom) Valid(v *validation.Validation) {

	classroom := Classroom{
		ClassroomName: c.ClassroomName,
	}

	o := orm.NewOrm()

	err := o.Read(&classroom, "ClassroomName")

	if nil == err {
		v.SetError("ClassroomName", "ClassroomName can not repeat")
	}

	classroom = Classroom{
		ClassroomAdress: c.ClassroomAdress,
	}

	err = o.Read(&classroom, "ClassroomAdress")

	if nil == err {
		v.SetError("ClassroomAdress", "ClassroomAdress can not repeat")
	}

}

func init() {

	orm.RegisterModel(new(Classroom))
}
