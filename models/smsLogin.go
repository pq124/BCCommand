package models

type SmsLogin struct {
	Telephone string `form:"telephone"`
	Code  string `form:"code"`
}
