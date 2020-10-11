package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

//处理用户登录
func (r *RegisterController)Post()  {
	//解析用户端提交的请求数据
	var user models.User

	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("不好意思，出错了")
		return
	}

	//将解析到的数据保存到数据库中

	//2.保存用户信息到数据库
	_,err =user.SaveUser()




	//将处理结果返回给客户端浏览器
	//如果成功，跳转登陆界面
	//如果失败，则提示失败信息


	//template 模板


}