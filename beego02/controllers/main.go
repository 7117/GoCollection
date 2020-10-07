package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Test() {

	c.SetSession("nasssme","aaaa");
	aa := c.GetSession("nasssme")
	fmt.Println(aa)


}

func (c *MainController) Other() {

	c.Data["aa"]= "aa";
	c.Data["bb"]= "bb";
	c.TplName = "maincontroller/test.tpl"
}

