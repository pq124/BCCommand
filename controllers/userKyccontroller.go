package controllers

import (
	"DataCertPlatform/models"
	"fmt"
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
	fmt.Println("userKyc里的phone：",user.Telephone)
//把用户的实名认证更新到数据库的用户表中
	_ ,err = user.UpdateUser()
	if err!=nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉链上数据更新失败")
		return
	}
	//fmt.Println("updata后的user：",user)
 u.TplName = "home.html"
}
