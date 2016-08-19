package controllers

import (
	"github.com/astaxie/beego"
	"golang_cms/models"
)

type BaseController struct {
	beego.Controller
	page      models.Page
}

func (this *BaseController) Prepare() {
	this.init()
}

func (this *BaseController) init()  {
	this.Layout = "common_layout.html"
	this.initPage()
}

//页面公共信息
func (this *BaseController) initPage() {
	this.page.SiteName = this.getConfig("sitename")
	this.page.Title = this.getConfig("title")
	this.page.Company = this.getConfig("company")
	this.page.Copyright = this.getConfig("copyright")
	this.page.Domain = this.getConfig("domain")
	this.page.Keywords = this.getConfig("keywords")
	this.page.Description = this.getConfig("description")
	this.page.Author = this.getConfig("author")
}

//读取配置
func (this *BaseController) getConfig(key string) string {
	return beego.AppConfig.String(key)
}