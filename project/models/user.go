package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Access struct {
	Id    int64
	Title string
}
//定義的全局的變量
var o orm.Ormer

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test1?charset=utf8", 30)
	orm.RegisterModel(new(Access))
	o = orm.NewOrm()
}

func AddAccess(access *Access) (int64, error) {
	id, err := o.Insert(access)
	return id, err

}
