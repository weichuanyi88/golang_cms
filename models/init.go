package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

//页面公共信息
type Page struct {
	SiteName    string //网站名称
	Title       string //页面标题
	Company     string //公司名称
	Domain      string //域名
	Copyright   string //版权
	Keywords    string //Seo关键词
	Description string //Seo描述
	Author      string //作者
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1234@tcp(47.88.24.94:3306)/d_go_cms?charset=utf8")
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}