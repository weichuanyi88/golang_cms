package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id      int64
	TypeId  int64
	Image   string
	Title   string
	BianHao string
	GuiGe   string
	Content string
	Addtime int64
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix("t_", new(Product))
}

func (this *Product) getProductListByType(typeId int64) ([]Product) {
	var productList []Product
	return productList

}

/**
获取商品详情页内容
 */
func (this *Product) GetProductById(id int64) (Product){
	var productDetail Product
	orm.NewOrm().Raw("SELECT * FROM t_product WHERE id = ?", id).QueryRow(&productDetail)
	return productDetail
}