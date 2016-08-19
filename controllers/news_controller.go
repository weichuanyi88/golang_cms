package controllers

import (
	"golang_cms/models"
	"golang_cms/utils"
)

type NewsController struct {
	BaseController
}

var news = models.News{}

func (this *NewsController)  GetNewsById() {
	id, err := this.GetInt64("id")
	newsDetail := news.GetNewsById(id)
	if (err == nil) {
		this.Data["newsDetail"] = newsDetail
	}
	this.TplName = "newsDetail.html"
}

func (this *NewsController)  GetNewsList() {
	pno, _ := this.GetInt("pno") //获取当前请求页
	var newsList []models.News
	var conditions string = " order by id desc" //定义日志查询条件,格式为 " and name='zhifeiya' and age=12 "
	var po utils.PageOptions                                                                                                //定义一个分页对象
	po.TableName = "t_news"                                                                                                 //指定分页的表名
	po.EnableFirstLastLink = true                                                                                   //是否显示首页尾页 默认false
	po.EnablePreNexLink = true                                                                                          //是否显示上一页下一页 默认为false
	po.Conditions = conditions                                                                                          // 传递分页条件 默认全表
	po.Currentpage = int(pno)                                                                                           //传递当前页数,默认为1
	po.PageSize = 2                                                                                                                //页面大小  默认为20

	//返回分页信息,
	//第一个:为返回的当前页面数据集合,ResultSet类型
	//第二个:生成的分页链接
	//第三个:返回总记录数
	//第四个:返回总页数
	//var rs orm.RawSeter
	totalItem, _, rs, pagerhtml := utils.GetPagerLinks(&po, this.Ctx)
	rs.QueryRows(&newsList)
	this.Data["newsList"] = newsList //把当前页面的数据传递到前台
	this.Data["pagerhtml"] = pagerhtml
	this.Data["totalItem"] = totalItem
	this.Data["PageSize"] = po.PageSize

	//newsList := news.GetNewsListOrderByAddTimeDesc()
	//this.Data["newsList"] = newsList
	//this.Data["page"] = this.page

	this.TplName = "newsList.html"
}