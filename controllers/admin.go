package controllers

type AdminController struct {
	BaseController
}

func (this *AdminController) Admin(){

	this.setTplName("admin_index.html")
}
