package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

//处理用户注册
func (r *RegisterController)Post()  {
	//解析用户端提交的请求数据
	var user models.User

	err := r.ParseForm(&user)
	if err != nil {

		r.Ctx.WriteString("解析失败")
		return
	}

	//将解析到的数据保存到数据库中
      _ , err = user.AddUser()
	if err!=nil {
		r.Ctx.WriteString("抱歉用户注册失败")
		return
	}
	r.TplName="login.html"
}