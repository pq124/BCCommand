package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beegio.me"
	//c.Data["Email"] = "astaxie@gmal.com"
	c.TplName = "register.html"
}



