package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/register",&controllers.RegisterController{})

	beego.Router("/login.html",&controllers.LoginController{})

	beego.Router("/home.html",&controllers.LoginController{})
    //文件上传的功能
	beego.Router("/upload",&controllers.UploadController{})

	beego.Router("/list_record",&controllers.UploadController{})
    //用户新增存证
    //beego.Router("home.html",&controllers.UploadController{})









}
