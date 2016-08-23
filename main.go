package main

import (
	_ "golang_cms/routers"
	"github.com/astaxie/beego"
	_ "golang_cms/models"
)

func main() {
	//models.Init()
	beego.Run()
}

