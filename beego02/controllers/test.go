package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string

}

func (this *TestController) Test() {
	id := this.GetString("id")
	id_two :=this.Input().Get("id")
	this.Ctx.WriteString(id);
	this.Ctx.WriteString(id_two);
}

func (this *TestController) Other() {
	this.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/post" method="post">
			<input type="text" name="Username">
			<input type="text" name="Password">
			<input type="submit" value="提交">
		</form></html>`)
}

func (this *TestController) Post(){
	u := User{}

	if err:=this.ParseForm(&u);err == nil{
		this.Ctx.WriteString("name:"+ u.Username + "password:" + u.Password)
	}


}
