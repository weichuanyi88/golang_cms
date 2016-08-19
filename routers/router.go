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

	product := &controllers.ProductController{}
	beego.Router("/product/getlist", product, "get,post:GetProductListByType")
	beego.Router("/product/getproductdetail", product, "get,post:GetProductDetailById")

}
