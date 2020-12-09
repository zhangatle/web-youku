package routers

import (
	"github.com/astaxie/beego"
	"web-youku/controllers"
)

func init() {
	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("mini/login", &controllers.UserController{}, "get:MiniLogin")
}
