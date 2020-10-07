package routers

import (
	"goPractice/beego02/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Test")
    beego.Router("/test", &controllers.MainController{},"get:Test")
    beego.Router("/other", &controllers.MainController{},"get:Other")

	beego.Router("two", &controllers.TestController{},"get:Test")
	beego.Router("/test/two", &controllers.TestController{},"get:Test")
	beego.Router("/other/two", &controllers.TestController{},"get:Other")
	beego.Router("/post", &controllers.TestController{},"*:Post")

    beego.Router("/login", &controllers.TestController{},"*:Login")

    beego.Router("/testmodel", &controllers.ModelController{},"*:Test")

    beego.Router("/insert", &controllers.ModelController{},"*:Insert")
}
