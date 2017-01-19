package controllers

import (
	"golang_cms/models"
	"golang_cms/utils"
	"strconv"
	"fmt"
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
	this.setTplName("newsDetail.html")
}

func (this *NewsController)  GetNewsList() {
	pno, _ := this.GetInt("pno") //获取当前请求页
	var newsList []models.News
	var conditions string = " order by id desc" //定义日志查询条件,格式为 " and name='zhifeiya' and age=12 "
	var po utils.PageOptions //定义一个分页对象
	po.TableName = "t_news" //指定分页的表名
	po.EnableFirstLastLink = true //是否显示首页尾页 默认false
	po.EnablePreNexLink = true //是否显示上一页下一页 默认为false
	po.Conditions = conditions //传递分页条件 默认全表
	po.Currentpage = int(pno) //传递当前页数,默认为1
	po.PageSize = 2 //页面大小  默认为20

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
	this.setTplName("newsList.html")
}

func (this *NewsController)  GetNewsList4Admin() {
	pno, _ := this.GetInt("pno") //获取当前请求页
	var newsList []models.News
	var conditions string = " order by id desc" //定义查询条件,格式为 " and name='zhifeiya' and age=12 "
	var po utils.PageOptions //定义一个分页对象
	po.TableName = "t_news" //指定分页的表名
	po.EnableFirstLastLink = true //是否显示首页尾页 默认false
	po.EnablePreNexLink = true //是否显示上一页下一页 默认为false
	po.Conditions = conditions //传递分页条件 默认全表
	po.Currentpage = int(pno) //传递当前页数,默认为1
	po.PageSize = 2 //页面大小  默认为20

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
	this.setTplName("admin_newsList.html")
}

func (this *NewsController)  GetNewsListByTypeId() {
	pno, _ := this.GetInt("pno") //获取当前请求页
	typeId, _ := this.GetInt64("typeid") //获取当前请求页
	var newsList []models.News
	var conditions string = " and typeId=" + strconv.FormatInt(typeId, 10) + " order by id desc" //定义日志查询条件,格式为 " and name='zhifeiya' and age=12 "
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
	this.setTplName("newsList.html")
}

func (this *NewsController)  LinkToAddNews() {
	this.setTplName("admin_addNews.html")
}

func (this *NewsController) GetSubType4Select() {
	fmt.Println("#############################")
	parentTypeId, _ := this.GetInt64("parentTypeId") //获取当前请求页
	if (parentTypeId > 0) {
		fmt.Println("进来拉")
		fmt.Println(parentTypeId)
		 selectOptionHtml := "<select  class=\"form-control\" id=\"form-field-select-1\" name=\"Typeid\" >"+
		"<option  value=\"\">&nbsp;</option>"+
		 "<option value=\"1\">a</option>"+
		"<option value=\"2\">b</option>"+
		"<option value=\"3\">c</option>"+
		"</select>"
		this.Ctx.WriteString(selectOptionHtml)
	}else{
		fmt.Println("mei 进来拉")
		this.Ctx.WriteString("")

	}

}