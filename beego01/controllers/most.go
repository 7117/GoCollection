package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

//由于model这个名字叫 UserInfo 那么操作的表其实 user_info
type Access struct{ 
	Id int64
	Title string
}

type MostController struct {
	beego.Controller
}

func (c *MostController) Get() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test1?charset=utf8", 30)
	orm.RegisterModel(new(Access))

	o := orm.NewOrm()
	//下面是插入
	// user := Access{Id:6, Title:"123456"}
	// aaa,_ := o.Insert(&user)
	// c.Ctx.WriteString(fmt.Sprintf("user info:%v", aaa))

	//下面是更新
	// user := Access{Id:6, Title:"dddd"}
	// user.Id = 6
	// o.Update(&user)
	
	//下面是读取
	usera := Access{Id:6}
	o.Read(&usera)

	c.Ctx.WriteString(fmt.Sprintf("user info:%v", usera))
}


// 原始sql
func (c *MostController) Deal() {

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test1?charset=utf8", 30)
	orm.RegisterModel(new(Access))
	o := orm.NewOrm()

	//下面是原生读取
	// var users []Access
	// o.Raw("select * from access").QueryRows(&users)
	// c.Ctx.WriteString(fmt.Sprintf("%v", users))


	// 查询构建器
	var aaa []Access
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("title").From("access").Where("id=?")
	sql := qb.String()
	fmt.Print(sql)
	o.Raw(sql,"1").QueryRows(&aaa)

	c.Ctx.WriteString(fmt.Sprintf("%v", aaa))


}

