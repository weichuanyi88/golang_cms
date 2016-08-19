package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int64
	Username string
	Password string
	Nickname string
}

func init() {
	// 需要在init中注册定义的model
	//orm.RegisterModel(new(User))
	orm.RegisterModelWithPrefix("t_", new(User))
}

func  (this *User) AddUser(m *User)  (int64, error){
	return ormer.Insert(m)
}

func  (this *User) GetUser(id int64)  (User){
	user := User{Id: id}
	ormer.Read(&user)
	return user
}