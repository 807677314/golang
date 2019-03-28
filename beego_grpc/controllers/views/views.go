package views

import (
	"beego_grpc/models/official_content_models"
	"time"
)

type CMSCategoryView struct {
	Id       int               `json:"id"`
	Name     string            `json:"name"`
	Lang     string            `json:"lang"`
	Contents []*CMSContentView `json:"contents"`
}

func (this *CMSCategoryView) ConvertDown(entity *official_content_models.CMSCategory) {
	this.Id = entity.Id
	this.Name = entity.Name
	this.Lang = entity.Lang

	for _, content := range entity.Contents {
		contentView := &CMSContentView{}
		contentView.ConvertDown(content)
		this.Contents = append(this.Contents, contentView)
	}
}

func (this *CMSCategoryView) ConvertUp(entity *official_content_models.CMSCategory) {
	entity.Id = this.Id
	entity.Name = this.Name
	entity.Lang = this.Lang

	for _, contentView := range this.Contents {
		content := &official_content_models.CMSContent{}
		contentView.ConvertDown(content)
		entity.Contents = append(entity.Contents, content)
	}
}

type CMSContentView struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	SubTitle string    `json:"subTitle"`
	Content  string    `json:"content"`
	Created  time.Time `json:"created"`
}

func (this *CMSContentView) ConvertDown(entity *official_content_models.CMSContent) {
	this.Id = entity.Id
	this.Title = entity.Title
	this.SubTitle = entity.SubTitle
	this.Content = entity.Content
	this.Created = entity.Created
}

func (this *CMSContentView) ConvertUp(entity *official_content_models.CMSContent) {
	entity.Id = this.Id
	entity.Title = this.Title
	entity.SubTitle = this.SubTitle
	entity.Content = this.Content
	entity.Created = this.Created
}
