package controllers

import (
	"golang_cms/models"
)
type ArticleController struct {
	BaseController
}

var article = models.Article{}

func (this *ArticleController)  GetArticleById() {
	id, err := this.GetInt64("id")
	articleDetail := article.GetArticleById(id)
	if (err == nil) {
		this.Data["articleDetail"] = articleDetail
	}
	this.TplName = "articleDetail.html"
}