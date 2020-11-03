package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
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
	  //判断数据库查查询结果
	 if err!=nil {
	 	fmt.Println(err.Error())
	 	fmt.Println("U为",u)
		 l.Ctx.WriteString("登陆失败")
		 return
	 }
	 //3.1增加逻辑:判断用户是否已经实名认证,如果没有实名认证,则跳转到认证页面
	 if strings.TrimSpace( u.Name) ==""|| strings.TrimSpace(u.Card) == "" {
	 	l.Data["Telephone"]=user.Telephone
	 	l.TplName="user_kyc.html"
		 return
	}


	    l.Data["Telephone"] = u.Telephone//动态数据设置
	    l.TplName = "home.html"
 }

