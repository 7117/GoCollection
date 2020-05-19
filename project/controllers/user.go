package controllers

//package是namespace
//use是import
//struct是class
import (
	"github.com/astaxie/beego"
	"project/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Method() {
	access := models.Access{Id: 22, Title: "ssss"}
	models.AddAccess(&access)

	c.Ctx.WriteString("write")
}

func (c *UserController) Page(){
	c.Data["website"] = "website"
	c.Data["email"] = "email"
	c.Data["isplay"] = 1

	a := []int{1,2,3,4,5}

	c.Data["a"] = a

	c.TplName = "user.tpl"
}