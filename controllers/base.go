package controllers

import (
	"github.com/astaxie/beego"
	"golang_cms/models"
	"strings"
	"fmt"
)

type BaseController struct {
	beego.Controller
	page models.Page
}

const (
	CONST_MODULE_ADMIN = "admin" //管理后台页面
	CONST_MODULE_DEFAULT = "default" //默认前端页面
	BIG_PIC_PATH   = "./static/upload/bigpic/"
	SMALL_PIC_PATH = "./static/upload/smallpic/"
	FILE_PATH      = "./static/upload/attachment/"

)

var pathArr []string = []string{"", BIG_PIC_PATH, SMALL_PIC_PATH, FILE_PATH}

func (this *BaseController) Prepare() {
	this.init()
}

func (this *BaseController) init() {
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

func (this *BaseController) setTplName(tplname string) {
	fmt.Println("tplname-------------------------------------->"+tplname)
	fmt.Println("URI------------------------------------------>"+this.Ctx.Request.RequestURI)
	if (strings.HasPrefix(this.Ctx.Request.RequestURI,"/admin")) {
		this.Layout = "admin/admin_layout.html"
		this.TplName = "admin/"+tplname
	}else{
		this.Layout = "common_layout.html"
		this.TplName = tplname
	}
}