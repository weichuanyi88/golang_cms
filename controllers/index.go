package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	companyNews := news.GetTopNewsByTypeId(1)
	customerNews := news.GetTopNewsByTypeId(2)
	industryNews := news.GetTopNewsByTypeId(3)
	technologyNews := news.GetTopNewsByTypeId(4)
	this.Data["companyNews"] = companyNews
	this.Data["customerNews"] = customerNews
	this.Data["industryNews"] = industryNews
	this.Data["technologyNews"] = technologyNews

	this.Data["page"] = this.page
	this.TplName = "index.html"
}
