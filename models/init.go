package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

//var (
//	ormer orm.Ormer
//)

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

func Init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1234@/d_go_cms?charset=utf8")
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	//orm.NewOrm() = orm.NewOrm()
	//orm.NewOrm().Using("default")
}