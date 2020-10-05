package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type ModelController struct {
	beego.Controller
}

type Access struct {
	Id       int
	Username string
	Password string
}

func (c *ModelController) Test() {
	orm.RegisterModel(new(Access))
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test1?charset=utf8", 30)

	o := orm.NewOrm()

	user := Access{Username: "zs", Password: "123456"}

	//增加
	//id, err := o.Insert(&user)

	//更新
	//user.Id = 1;
	//user.Username = "ss"
	//id, err := o.Update(&user)

	//c.Ctx.WriteString(string(id))

	//读取
	user.Id = 1;
	o.Read(&user)

	//if err != nil {
	//	panic(err)
	//}

	c.Ctx.WriteString(fmt.Sprintf("info:%v", user))
}
