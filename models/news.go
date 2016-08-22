package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type News struct {
	Id int64
	Title string
	Content string
	TypeId int64
	Image string
	AddTime int64
	ParentTypeId int64
}


func init() {
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix("t_", new(News))
}


/**
首页按新闻分类分别获取10条
 */
func  (this *News) GetTopNewsByTypeId(id int64)  ([]News){
	var newsList []News
	//ormer := orm.NewOrm()
	//ormer.Using("default")
	num,err := ormer.Raw("SELECT id, title,content,typeId,image,addTime,parentTypeId FROM t_news WHERE typeId = ? limit 10", id).QueryRows(&newsList)
	if(err ==nil){
		fmt.Println(num)
	}
	return newsList
}

/**
资讯公告获取最新20条
 */
func  (this *News) GetNewsListOrderByAddTimeDesc()  ([]News){
	var newsList []News
	//ormer := orm.NewOrm()
	//ormer.Using("default")
	num,err := ormer.Raw("SELECT id, title,content,typeId,image,addTime,parentTypeId FROM t_news limit 20").QueryRows(&newsList)
	if(err ==nil){
		fmt.Println(num)
	}
	return newsList
}

/**
按新闻ID获取新闻详情
 */
func  (this *News) GetNewsById(id int64)  (News){
	var newsDetail News
	//ormer := orm.NewOrm()
	//ormer.Using("default")
	ormer.Raw("SELECT * FROM t_news WHERE id = ?", id).QueryRow(&newsDetail)
	return newsDetail
}