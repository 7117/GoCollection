package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type User struct{
	Username string
	Password string
}


func (c *MainController) Deal() {

	content := `<html><form methos="post" action="http://127.0.0.1:8080/main/method">
	<input type="text" name="Username">
	<input type="text" name="Password">
	<input type="submit" name="name" value="提交">
</form></html>`;
	c.Ctx.WriteString(content)
}

func (c *MainController) Method() {
	u := User{}
	if err:=c.ParseForm(&u) ; err != nil{
	}

	c.Ctx.WriteString("Username:" + u.Username + " Password:" + u.Password)
}

func (c *MainController) Tocookie() {
	
	c.Ctx.SetCookie("name","info", 1000, "/")
	value := c.Ctx.GetCookie("name");
	c.Ctx.WriteString(value)
}

func (c *MainController) Tosession() {
	c.SetSession("name","sunxiao")

	var count interface{}
	count = c.GetSession("name")
	fmt.Println(count)
}