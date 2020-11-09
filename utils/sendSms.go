package utils

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
	"math/rand"
	"strings"
	"time"
)

type  SmsCode struct {
	Code string `form:"code"`
}

type SmsResult struct {
	BizId string
	Code string
	Message string
	RequestId string
}

const SMS_TLP_REGISTER = "SMS_205393604" //注册业务的短信模板
const SMS_TLP_LOGIN = "SMS_205398654"    //用户登录的短信模板
const SMS_TLP_KYC = ""      //实名认证的短信模板



//该函数用于发送一条短信
//code
//telephone :电话,接收验证码的号码
func SendSms(telephone string,code string,templateType string)(*SmsResult,error){
	config := beego.AppConfig
	accessKey :=config.String("accessKey")
	accessKeySecrt :=config.String("accessKeySecrt")
	client ,err :=dysmsapi.NewClientWithAccessKey("cn-hangzhou",accessKey,accessKeySecrt)
	if err!=nil {
		return nil ,err
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = telephone
	request.SignName = "线上餐厅"
	request.TemplateCode = templateType
	smsCode := SmsCode{Code:code,}
	smsbytes,_:=json.Marshal(smsCode)
	request.TemplateParam = string(smsbytes)

	response,err:=client.SendSms(request)
	if err!=nil {
		return nil ,err
	}
	SmsResult:=SmsResult{
		BizId:    response.BizId ,
		Code:      response.Code,
		Message:   response.Message,
		RequestId: response.RequestId,
	}
	return &SmsResult,nil
}
//生成随机数
func GenRandCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}
