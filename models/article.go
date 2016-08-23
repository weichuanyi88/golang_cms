package models

import "github.com/astaxie/beego/orm"

type Article struct {
	Id int64
	Title string
	Content string
	AddTime int64
	TypeId int64
}

func init(){
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix("t_", new(Article))
}

/**
按文章ID获取文章详情
 */
func  (this *Article) GetArticleById(id int64)  (Article){
	var article Article
	orm.NewOrm().Raw("SELECT * FROM t_article WHERE id = ?", id).QueryRow(&article)
	return article
}