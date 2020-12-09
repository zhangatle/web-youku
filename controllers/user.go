package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// 用户登录页面
func (c *UserController) Login() {
	fmt.Println(1)
	c.TplName = "login.html"
}

// 登录页mini
func (c *UserController) MiniLogin() {
	c.TplName = "mini_login.html"
}
