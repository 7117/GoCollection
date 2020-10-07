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

	s := make([]int,10)
	s[0]=0;
	s[1]=1;
	s[2]=2;
	s[3]=3;

	c.Data["aa"]= "aa";
	c.Data["bb"]= "bb";
	c.Data["ss"] = s ;
	c.TplName = "maincontroller/test.tpl"
}

