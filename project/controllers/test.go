package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Method() {
	c.Data["website"] = "beego"
	c.TplName = "testcontroller/method.tpl"
}

func (c *TestController) Deal() {
	c.Ctx.WriteString("背景颜色")
}