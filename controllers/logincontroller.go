package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get()  {
    l.TplName ="login.html"
}

func (l *LoginController)Post() {
    var user models.User
	err := l.ParseForm(&user)
		if err != nil {
			l.Ctx.WriteString("解析失败")
			return
	}
	//3、判断数据库查询结果
	  u,err :=  user.QueryData()
	 if err!=nil {
	 	fmt.Println(err.Error())
		 l.Ctx.WriteString("登陆失败")
		 return
	 }
	    l.Data["Telephone"] = u.Telephone//动态数据设置
	    l.TplName = "home.html"
 }

