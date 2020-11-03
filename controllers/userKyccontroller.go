package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
)

type UserKycController struct {
	beego.Controller
}

func (u *UserKycController)Get()  {
	u.TplName= "user_kyc.html"
}

func (u *UserKycController)Post()  {
	var  user models.User
	err:=u.ParseForm(&user)
	if err!=nil {
		u.Ctx.WriteString("抱歉,数据解析错误,请重试")
		return
	}
//把用户的实名认证更新到数据库的用户表中
	_,err=user.Updata()
	if err!=nil {
		u.Ctx.WriteString("抱歉链上数据更新失败")
		return
	}
 u.TplName = "home.html"
}
