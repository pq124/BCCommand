package models

import "DataCertPlatform/db_mysql"

type SmsRecord struct {
	BizId     string
	Telephone string
	Code      string
	Status    string
	Message   string
	TimeStemp int64
}
//向数据空当中插入验证码记录,该记录由阿里云第三方返回
func (s SmsRecord) SaveSmsRecord()(int64,error){
	rs,err:=db_mysql.Db.Exec("insert into sms_record(biz_id,telephone,code,status,message,timestemp)values (?,?,?,?,?,?)",
		s.BizId,s.Telephone,s.Code,s.Status,s.Message,s.TimeStemp)
	if err!=nil {
		return -1,err
	}
	return rs.RowsAffected()
}
