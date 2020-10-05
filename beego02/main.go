package main

import (
	_ "goPractice/beego02/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

