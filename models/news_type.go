package models

import "github.com/astaxie/beego/orm"

type NewsType struct {
	Id         int64
	Name       string
	OrderIndex int64
	Parentid   int64
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix("t_", new(NewsType))

}

