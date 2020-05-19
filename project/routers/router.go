package routers

import (
	"project/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/main/method", &controllers.MainController{},"*:Method")
	beego.Router("/main/deal", &controllers.MainController{},"*:Deal")
	beego.Router("/main/tocookie", &controllers.MainController{},"*:Tocookie")
	beego.Router("/main/tosession", &controllers.MainController{},"*:Tosession")
	beego.Router("/test/method", &controllers.TestController{},"*:Method")
	beego.Router("/test/deal", &controllers.TestController{},"*:Deal")
}
