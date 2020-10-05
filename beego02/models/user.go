package models

import "github.com/astaxie/beego/orm"

type Access struct {
	Id       int
	Username string
	Password string
}

var o orm.Ormer

func init() {
	orm.Debug = true
	orm.RegisterModel(new(Access))
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test1?charset=utf8", 30)
	o = orm.NewOrm()

}

func AddUser(user_ibfo *Access) (int64, error) {
	id, err := o.Insert(user_ibfo)
	return id, err;
}
