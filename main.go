package main

import (
	_ "golang_cms/routers"
	"github.com/astaxie/beego"
	_ "golang_cms/models"
	"fmt"
)

func main() {
	//models.Init()
	//beego.SetStaticPath("/admin", "views/admin")
	fmt.Println(100-0.9)
	beego.Run()
}

