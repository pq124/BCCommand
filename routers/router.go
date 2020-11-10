package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/register",&controllers.RegisterController{})

	beego.Router("/login.html",&controllers.LoginController{})

	beego.Router("/login_sms.html",&controllers.LoginSmsController{})

	beego.Router("/send_sms",&controllers.SendSmsController{})

	beego.Router("/home.html",&controllers.LoginController{})
    //文件上传的功能
	beego.Router("/upload",&controllers.UploadController{})

	beego.Router("/list_record",&controllers.UploadController{})

	beego.Router("/cert_detail.html",&controllers.CertDetailController{})
	//用户新增存证
    //beego.Router("home.html",&controllers.UploadController{})
    //查看认证数据证书页面()
    beego.Router("/user_kyc",&controllers.UserKycController{})

}
