package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"strconv"
)

type HttpController struct {
	beego.Controller
}


func (this *HttpController) Http() {

	req := httplib.Get("http://www.baidu.com")
	str,_ := req.String()

	this.Ctx.WriteString("info");
	this.Ctx.WriteString(str);
}

func (this *HttpController) Context() {
	this.Ctx.WriteString("<br>")
	this.Ctx.WriteString(this.Ctx.Input.IP())
	this.Ctx.WriteString("<br>")
	this.Ctx.WriteString(strconv.Itoa(this.Ctx.Input.Port()))
	this.Ctx.WriteString("<br>")


	m := make(map[string]float64)
	m["a"] = 1.1;
	m["b"] = 2.2;
	this.Ctx.Output.JSON(m,false,false)
}