package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type News struct {
	Id int64
	Title string
	Content string
	Typeid int64
	Image string
	Addtime int64
	ParenttypeId int64
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
	orm.NewOrm().Raw("SELECT id, title,content,typeId,image,addTime,parentTypeId FROM t_news WHERE typeId = ? limit 10", id).QueryRows(&newsList)
	return newsList
}

/**
资讯公告获取最新20条
 */
func  (this *News) GetNewsListOrderByAddTimeDesc()  ([]News){
	var newsList []News
	orm.NewOrm().Raw("SELECT id, title,content,typeId,image,addTime,parentTypeId FROM t_news limit 20").QueryRows(&newsList)
	return newsList
}

/**
按新闻ID获取新闻详情
 */
func  (this *News) GetNewsById(id int64)  (News){
	var newsDetail News
	orm.NewOrm().Raw("SELECT * FROM t_news WHERE id = ?", id).QueryRow(&newsDetail)
	return newsDetail
}