package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Overall struct {
	Id        int        `orm:"pk;auto"`
	Class     *Class     `orm:"null;rel(fk)"`
	Classroom *Classroom `orm:"null;rel(fk)"`
	Teacher   *Teacher   `orm:"null;rel(fk)"`
	DeleteAt  time.Time  `orm:"type(datetime);null;"form:"-"`
	UpdateAt  time.Time  `orm:"type(datetime);auto_now"form:"-"`
}

func init() {
	orm.RegisterModel(new(Overall))
}
