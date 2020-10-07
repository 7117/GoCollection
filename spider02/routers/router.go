package routers

import (
	"goPractice/spider02/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.SpiderController{},"*:CrawlMovie")
}
