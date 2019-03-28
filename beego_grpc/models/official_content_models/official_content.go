package official_content_models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type CMSCategory struct {
	Id       int           `orm:"column(id)"`
	Name     string        `orm:"column(name)"`
	Lang     string        `orm:"column(lang)"`
	DelFlag  string        `orm:"column(del_flag);type(char);size(1);default(N)"`
	Created  time.Time     `orm:"auto_now_add;null;type(datetime)"`
	Updated  time.Time     `orm:"auto_now;null;type(datetime)"`
	Contents []*CMSContent `orm:"reverse(many)"`
}

func (this *CMSCategory) TableName() string {
	return "cms_category"
}

// 多字段唯一键
func (this *CMSCategory) TableUnique() [][]string {
	return [][]string{
		{"name", "lang", "del_flag"},
	}
}

type CMSContent struct {
	Id       int          `orm:"column(id)"`
	Lang     string       `orm:"-"`
	Title    string       `orm:"column(title)"`
	SubTitle string       `orm:"column(sub_title)"`
	FrontPic string       `orm:"column(front_pic);size(1024)"`
	Content  string       `orm:"column(content);type(text)"`
	DelFlag  string       `orm:"column(del_flag);type(char);size(1);default(N)"`
	Created  time.Time    `orm:"auto_now_add;null;type(datetime)"`
	Updated  time.Time    `orm:"auto_now;null;type(datetime)"`
	Category *CMSCategory `orm:"rel(fk);null;on_delete(set_null)"`
}

func (this *CMSContent) TableName() string {
	return "cms_content"
}

// 删除对象
func RemoveEntity(item interface{}) error {
	if _, err := orm.NewOrm().Delete(item); err != nil {
		return err
	} else {
		return nil
	}
}

// 查询Entity详情
func GetEntity(item interface{}, cols ...string) error {
	return orm.NewOrm().Read(item, cols...)
}

// 插入Content对象
func Insert(entity interface{}) int64 {
	num, err := orm.NewOrm().InsertOrUpdate(entity)
	if err != nil {
		return 0
	}

	return num
}

// 查询所有分类列表
func GetCategoryAll(lang string) *[]CMSCategory {
	cmsList := &[]CMSCategory{}

	querySel := orm.NewOrm().QueryTable(&CMSCategory{}).Filter("del_flag", "N")
	if len(lang) > 0 {
		querySel = querySel.Filter("lang", lang)
	}
	_, err := querySel.All(cmsList)

	if err != nil {
		return &[]CMSCategory{}
	}

	return cmsList
}

// 查询所有分类列表
func GetCatontentList(content *CMSContent, pageIndex, pageSize int) *[]CMSContent {
	cmsList := &[]CMSContent{}

	querySel := orm.NewOrm().QueryTable(&CMSContent{}).Filter("del_flag", "N")
	if content.Category != nil && content.Category.Id > 0 {
		querySel = querySel.Filter("category_id", content.Category.Id)
	}

	if len(content.Title) > 0 {
		querySel = querySel.Filter("title__icontains", content.Title)
	}

	if len(content.Lang) > 0 {
		querySel = querySel.Filter("category__lang", content.Lang).RelatedSel()
	}

	offset := (pageIndex - 1) * pageSize
	_, err := querySel.OrderBy("created").Limit(pageSize, offset).All(cmsList)

	if err != nil {
		return &[]CMSContent{}
	}

	return cmsList
}

// 根据分类获取列表
func GetContentListByCategories(categories ...CMSCategory) *[]CMSContent {
	var ids []int
	for _, item := range categories {
		ids = append(ids, item.Id)
	}

	return GetContentListByCategoryIds(ids...)
}

// 根据分类获取列表
func GetContentListByCategoryIds(ids ...int) *[]CMSContent {
	cmsList := &[]CMSContent{}

	if len(ids) <= 0 {
		return &[]CMSContent{}
	}

	_, err := orm.NewOrm().QueryTable(&CMSContent{}).
		Filter("category__in", ids).
		RelatedSel().
		All(cmsList)

	if err != nil {
		return &[]CMSContent{}
	}

	return cmsList
}

func init() {
	orm.RegisterModel(new(CMSContent), new(CMSCategory))
}
