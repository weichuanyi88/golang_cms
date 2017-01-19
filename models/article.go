package models

import "github.com/astaxie/beego/orm"

type Article struct {
	Id int64  `form:"Id"`
	Title string `form:"Title"`
	Content string
	Addtime int64
	Typeid int64
	Image string
	Keywords string
	Description string
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

/**
按文章ID删除文章
 */
func  (this *Article) DeleteArticleById(id int64) (int64, error) {
	return  orm.NewOrm().Delete(&Article{Id: id})
}

/**
添加文章
 */
func  (this *Article) AddArticle(articleObj Article)  (int64, error){
	return orm.NewOrm().Insert(&articleObj)
}

/**
修改文章
 */
func  (this *Article) UpdateArticle(articleObj Article)  (int64, error){
	return orm.NewOrm().Update(&articleObj)
}

/**
获取文章列表
 */
func (this *Article) GetArticleList() ([]Article){
	var articleList []Article
	orm.NewOrm().Raw("SELECT id, title,typeid,addTime FROM t_article").QueryRows(&articleList)
	return articleList
}