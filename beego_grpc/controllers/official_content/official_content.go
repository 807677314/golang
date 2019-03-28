package official_content

import (
	"beego_grpc/tools/error"
	"beego_grpc/controllers/views"
	"beego_grpc/models/official_content_models"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)

const (
	Y = "Y"
	N = "N"
)

type CMSContentController struct {
	beego.Controller
}

func (this *CMSContentController) URLMapping() {
	this.Mapping("getCategoryList", this.GetCategoryList)
	this.Mapping("addCategory", this.AddCategory)
	this.Mapping("getContent", this.GetContent)
	this.Mapping("addContent", this.AddContent)
	this.Mapping("removeContent", this.RemoveContent)
	this.Mapping("getContentList", this.GetContentList)
}

// @Title 获取内容分类目录列表
// @Description 获取内容分类目录列表
// @Param   lang  		string	true 	"语言类别"
// @Param   content		string	false 	"是否包含content 取值 [Y, N]"
// @Success 200 {object}   controllers.CMSCategoryView
// @Failure 400
// @Failure 404

// @router /getCategoryList [get]
func (this *CMSContentController) GetCategoryList() {
	lang := this.GetString("lang")
	includeContent := this.GetString("content")

	categories := *official_content_models.GetCategoryAll(lang)
	var result []*views.CMSCategoryView
	contents := *official_content_models.GetContentListByCategories(categories...)
	for _, category := range categories {
		categoryView := &views.CMSCategoryView{}
		categoryView.ConvertDown(&category)

		result = append(result, categoryView)
	}

	if includeContent == Y {
		for _, categoryView := range result {
			for _, content := range contents {
				if categoryView.Id == content.Category.Id {
					contentView := &views.CMSContentView{}
					contentView.ConvertDown(&content)
					categoryView.Contents = append(categoryView.Contents, contentView)
				}
			}
		}
	}

	//this.JsonSuccess(&result)
	this.Data["result"] = &result

	this.ServeJSON()

}

// @Title 		查询内容目录列表
// @Description 查询内容目录列表: 传入语言类型, 文章类别进行查询, 结果按时间降序
// @Param   pageSize	query		string	false 	"单页数量"
// @Param   pageIndex	query		string	false 	"当前页码"
// @Param   lang  		FromData	string	false 	"语言类别"
// @Param   categoryId	FromData	string	false 	"文章类型ID"
// @Param   title		FromData	string	false 	"文章标题模糊搜索"
// @Success 200 {object}   controllers.CMSCategoryView
// @Failure 400
// @Failure 404

// @router /getContentList [post]
func (this *CMSContentController) GetContentList() {
	type ContentParams struct {
		CategoryId int
		Lang       string
		Title      string
		SubTitle   string
	}

	pageSize, _ := this.GetInt("pageSize", 10)
	pageIndex, _ := this.GetInt("pageIndex", 1)
	if pageIndex <= 0 {
		pageIndex = 1
	}

	var ob ContentParams
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error("Param is not json", string(this.Ctx.Input.RequestBody))
		//this.JsonResult(&ERROR_PARAM_NOT_JSON)
		return
	}

	content := &official_content_models.CMSContent{
		Category: &official_content_models.CMSCategory{Id: ob.CategoryId},
		Title:    ob.Title,
		Lang:     ob.Lang,
		SubTitle: ob.SubTitle,
	}

	contents := *official_content_models.GetCatontentList(content, pageIndex, pageSize)

	log.Println(contents)
	this.Data["content"] = &content

	this.ServeJSON()
}

// @Title 根据ID获取内容详情
// @Description 根据ID获取内容详情
// @Param   name	formData   string    true      "分类标题"
// @Param   lang 	formData   string    true      "语言类型"
// @Success 200 {object}	controllers.Result        ""
// @Failure 400
// @Failure 404

// @router /addCategory [post]
func (this *CMSContentController) AddCategory() {
	type ContentParams struct {
		Name string
		Lang string
	}

	var ob ContentParams
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error("Param is not json", string(this.Ctx.Input.RequestBody))
		//this.JsonResult(&ERROR_PARAM_NOT_JSON)
		return
	}

	// 需要指定category的值
	category := &official_content_models.CMSCategory{Name: ob.Name, Lang: ob.Lang, DelFlag: "N"}
	if err := official_content_models.GetEntity(category, "Name", "Lang", "DelFlag"); err == nil {
		beego.Error("record exist: %v", category)
		//this.JsonResult(&ERROR_RECORD_EXIST)
		return
	}

	if official_content_models.Insert(category) <= 0 {
		beego.Error("content insert db failed; ", err.Error())
		//this.JsonResult(&ERROR_SYS_WRONG)
		return
	}

	this.Data["SUCCESS"] = (&error.SUCCESS)
	this.ServeJSON()
}

// @Title 根据ID获取内容详情
// @Description 根据ID获取内容详情
// @Param	id	query	int		true	"内容的ID"
// @Success 200 {object}   controllers.CMSContentView
// @Failure 400
// @Failure 404

// @router /getContent [get]
func (this *CMSContentController) GetContent() {
	id, err := this.GetInt("id")
	if err != nil {
		//this.JsonResult(&ERROR_PARAM_INVALID)
		return
	}

	var result views.CMSContentView
	content := &official_content_models.CMSContent{Id: id}
	if err := official_content_models.GetEntity(content); err == nil {
		result.ConvertDown(content)
	}

	this.Data["result"] = &result

	this.ServeJSON()
}

// @Title 新增Content内容
// @Description 新增Content内容
// @Param   title     formData   string    true      "内容标题"
// @Param   subTitle  formData   string    true      "副标题"
// @Param	frontPic  formData	 string	   true		 "封面连接"
// @Param   lang      formData   string    true      "语言类型"
// @Param   category  formData   id		   true      "分类类别"
// @Param   content   formData   string    true      "内容富文本信息"
// @Success 200 {object}   controllers.Result        ""
// @Failure 400
// @Failure 404

// @router /addContent [post]
func (this *CMSContentController) AddContent() {
	type ContentParams struct {
		Id       int
		Title    string
		SubTitle string
		FrontPic string
		Content  string
		Category int
	}

	var ob ContentParams
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error("param is not json", string(this.Ctx.Input.RequestBody))
		//this.JsonResult(&ERROR_PARAM_NOT_JSON)
		return
	}

	// 需要指定category的值
	category := &official_content_models.CMSCategory{Id: ob.Category}
	if err := official_content_models.GetEntity(category); err != nil {
		beego.Error("must exact the category value")
		//this.JsonResult(&ERROR_PARAM_INVALID)
		return
	}

	content := &official_content_models.CMSContent{Title: ob.Title, SubTitle: ob.SubTitle, Content: ob.Content, FrontPic: ob.FrontPic, Category: &official_content_models.CMSCategory{Id: ob.Category}, DelFlag: "N"}
	if official_content_models.Insert(content) <= 0 {
		//this.JsonResult(&ERROR_SYS_WRONG)
		return
	}

	this.Data["SUCCESS"] = (&error.SUCCESS)
	this.ServeJSON()
}

// @Title 新增Content内容
// @Description 新增Content内容
// @Param   contentId	query	string	true		"内容Id"
// @Success 200 {object}	controllers.Result		""
// @Failure 400
// @Failure 404

// @router /removeContent [post]
func (this *CMSContentController) RemoveContent() {
	contentId, _ := this.GetInt("contentId", -1)
	if contentId < 0 {
		//this.JsonResult(&ERROR_PARAM_INVALID)
		return
	}

	if err := official_content_models.RemoveEntity(&official_content_models.CMSContent{Id: contentId}); err != nil {
		//this.JsonResult(&ERROR_SYS_WRONG)
		return
	}

	this.Data["SUCCESS"] = (&error.SUCCESS)
	this.ServeJSON()
}
