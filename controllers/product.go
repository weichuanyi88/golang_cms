package controllers

import (
	"golang_cms/models"
	"golang_cms/utils"
	"strconv"
)

type ProductController struct {
	BaseController
}

var product_model models.Product

func (this *ProductController) GetProductListByType() {
	pno, _ := this.GetInt("pno") //获取当前请求页
	var productList []models.Product
	var conditions string = " order by id desc" //定义日志查询条件,格式为 " and name='zhifeiya' and age=12 "
	var po utils.PageOptions                                                                                                //定义一个分页对象
	po.TableName = "t_product"                                                                                                 //指定分页的表名
	po.EnableFirstLastLink = true                                                                                   //是否显示首页尾页 默认false
	po.EnablePreNexLink = true                                                                                          //是否显示上一页下一页 默认为false
	po.Conditions = conditions                                                                                          // 传递分页条件 默认全表
	po.Currentpage = int(pno)                                                                                           //传递当前页数,默认为1
	po.PageSize = 9                                                                                                                //页面大小  默认为20

	//返回分页信息,
	//第一个:为返回的当前页面数据集合,ResultSet类型
	//第二个:生成的分页链接
	//第三个:返回总记录数
	//第四个:返回总页数
	//var rs orm.RawSeter
	totalItem, _, rs, pagerhtml := utils.GetPagerLinks(&po, this.Ctx)
	rs.QueryRows(&productList)
	productsHtmlBody :=""
	if(productList != nil){
		for i, product := range productList {
			if(i==0){
				productsHtmlBody += "<div class=\"products-grid-lft\">"
			}
			productsHtmlBody += "<div class=\"products-grd\">"
			productsHtmlBody += "<div class=\"p-one simpleCart_shelfItem prd\">"
			productsHtmlBody += "<a href=\"/product/getproductdetail?id="+strconv.FormatInt(product.Id,10)+"\">"
			productsHtmlBody += "<img src=\""+product.Image+"\"+ alt=\"\" class=\"img-responsive\"/>"
			productsHtmlBody += "<div class=\"mask\">"
			productsHtmlBody += "<span>Quick View</span>"
			productsHtmlBody += "</div>"
			productsHtmlBody += "</a>"
			productsHtmlBody += "<h4>"+product.Title+"</h4>"
			productsHtmlBody += "<p><a class=\"item_add\" href=\"#\"><i></i> <span class=\" item_price valsa\">$729</span></a></p>"
			productsHtmlBody += "<div class=\"pro-grd\">"
			productsHtmlBody += "<a href=\"/product/getproductdetail?id="+strconv.FormatInt(product.Id,10)+"\">"+product.Bianhao+"</a>"
			productsHtmlBody += "</div>"
			productsHtmlBody += "</div>"
			productsHtmlBody += "</div>"
			if((i+1)%3==0){
				productsHtmlBody += "<div class=\"clearfix\"></div>"
				productsHtmlBody += "</div>"
				if(i<len(productList)-1){
					productsHtmlBody += "<div class=\"products-grid-lft\">"
				}
			}
			/*if(i==len(productList)-1){
				productsHtmlBody += "</div>"
			}*/

		}
	}

	this.Data["productsHtmlBody"] = productsHtmlBody //把当前页面的数据传递到前台
	this.Data["pagerhtml"] = pagerhtml
	this.Data["totalItem"] = totalItem
	this.Data["PageSize"] = po.PageSize
	this.setTplName("products.html")
}

func (this *ProductController) GetProductDetailById() {
	id, err := this.GetInt64("id")
	productDetail := product_model.GetProductById(id)
	if (err == nil) {
		this.Data["productDetail"] = productDetail
	}

	this.setTplName("productDetail.html")
}