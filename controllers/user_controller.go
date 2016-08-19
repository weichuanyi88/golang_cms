package controllers

import (
	"github.com/astaxie/beego"
	"golang_cms/models"
	"fmt"
)

type UserController struct {
	beego.Controller
}

var user = &models.User{}

func (this *UserController)  GetUser() {
	this.Layout = "common_layout.html"
	//http://127.0.0.1:8080/user/getuser?id=4
	fmt.Println("进入controller")
	id,err := this.GetInt64("id")
	u := user.GetUser(id)
	if(err ==nil){
		this.Data["user"] = u
	}else{
		u.Username="default"
		u.Nickname="default"
	}

	this.TplName = "gotest/user.html"
	//this.Render()
}
