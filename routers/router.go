package routers

import (
	"golang_cms/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})

	user := &controllers.UserController{}
	beego.Router("/user/getuser", user, "get,post:GetUser")

	news := &controllers.NewsController{}
	beego.Router("/news/getnewsdetail", news, "get,post:GetNewsById")
	beego.Router("/news/getnewslist", news, "get,post:GetNewsList")
	beego.Router("/news/getnewslistbytypeid", news, "get,post:GetNewsListByTypeId")
	beego.Router("/admin/news/list", news, "get,post:GetNewsList4Admin")
	beego.Router("/admin/news/linktoadd", news, "get,post:LinkToAddNews")
	beego.Router("/admin/newstype/getsubtype4select", news, "get,post:GetSubType4Select")


	product := &controllers.ProductController{}
	beego.Router("/product/getlist", product, "get,post:GetProductListByType")
	beego.Router("/product/getproductdetail", product, "get,post:GetProductDetailById")

	article := &controllers.ArticleController{}
	beego.Router("/article/getarticledetail", article, "get,post:GetArticleById")
	beego.Router("/admin/article/add", article, "get,post:AddArticle")
	beego.Router("/admin/article/linktoadd", article, "get,post:LinkToAddArticle")
	beego.Router("/admin/article/pagelist", article, "get,post:GetArticleList")
	beego.Router("/admin/article/linktoedit", article, "get,post:LinkToEditArticle")
	beego.Router("/admin/article/update", article, "get,post:UpdateArticle")
	beego.Router("/admin/article/delete",article, "*:DeleteArticleById")
	beego.Router("/admin/article/upload",article, "*:Upload")

	admin := &controllers.AdminController{}
	beego.Router("/admin", admin, "get,post:Admin")

}
