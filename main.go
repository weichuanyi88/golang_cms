package main

import (
	_ "golang_cms/routers"
	"github.com/astaxie/beego"
	"golang_cms/models"
)

func main() {
	models.Init()
	beego.Run()

}

