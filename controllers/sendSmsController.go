package controllers

import (
	"DataCertPlatform/models"
	"DataCertPlatform/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type SendSmsController struct {
	beego.Controller
}

/**
 * 发送短信验证码的功能
 */
func (s *SendSmsController) Post() {
	var smsLogin models.SmsLogin
	err := s.ParseForm(&smsLogin)
	if err != nil {
		s.Ctx.WriteString("发送验证码数据解析失败，请重试")
		return
	}
	fmt.Println("开始发送验证码")
	telephone := smsLogin.Telephone
	code := utils.GenRandCode(6) //返回一个6位的随机数
	result, err := utils.SendSms(telephone, code, utils.SMS_TLP_REGISTER)
	if err != nil {
		s.Ctx.WriteString("发送验证码失败，请重试！")
		return
	}
	if len(result.BizId) == 0 {
		fmt.Println("BizId的值:",result.BizId)
		s.Ctx.WriteString(result.Message)
		return
	}

	//验证码发送成功
	smsRecord := models.SmsRecord{
		BizId:     result.BizId,
		Telephone: telephone,
		Code:      code,
		Status:    result.Code,
		Message:   result.Message,
		TimeStemp: time.Now().Unix(),
	}
    _,err=smsRecord.SaveSmsRecord()
	if err!=nil {
		s.Ctx.WriteString("抱歉,获取验证码失败,请重试!")
		return
	}
	s.Data["Telephone"]=smsLogin.Telephone
	s.Data["BizId"]=smsRecord.BizId
	s.TplName = "login_sms.html"
}
