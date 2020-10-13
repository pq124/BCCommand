package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register",&controllers.RegisterController{})
	beego.Router("/login.html",&controllers.LoginController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/home",&controllers.LoginController{})
	beego.Router("/home.html",&controllers.LoginController{})







}
