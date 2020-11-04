package utils
/*
import (
	"encoding/json"
	"github.com/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
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



//该函数用于发送一条短信
//code
//telephone :电话,接收验证码的号码
func SendSms(telephone string,code string,templateType string)(*SmsResult,error){
	config := beego.AppConfig
	accessKey :=config.String("sms_access_key")
	accessKeySecrt :=config.String("sms_acess_secret")
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
func GenrandCode(width int)string {

}
 */