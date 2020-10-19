package models

import (
	"DataCertPlatform/db_mysql"
	"DataCertPlatform/utils"
)

type  UploadRecord struct {
	Id int
	UserId int
	FileName string
	FileSize int64
	FileCert  string
	FileTitle string
	CertTime  int64
	CertTimeFormat string //仅作为格式化展示使用的字段
}


func (u UploadRecord) SavaRecord()(int64 ,error){
	result ,err := db_mysql.Db.Exec("insert into upload_record(user_id, file_name, file_size, file_cert, file_title, cert_time)"+
		"values(?,?,?,?,?,?)",u.UserId,u.FileName,u.FileSize,u.FileCert,u.FileTitle,u.CertTime)

	if err!=nil {
		return -1,err
	}
	row,err := result.RowsAffected()
	if err!=nil {
		return -1,err
	}
	return row,nil
}

func QueryRecordsByUserId(userId int) ([]UploadRecord,error) {
	rs, err := db_mysql.Db.Query("select id, user_id, file_name, file_size, file_cert, file_title, cert_time  from upload_record where user_id = ?", userId)
	if err != nil {
		return nil, err

	}
		records := make([]UploadRecord, 0)

		for rs.Next() {
			var record UploadRecord
			err := rs.Scan(&record.Id, &record.UserId, &record.FileName, &record.FileSize, &record.FileCert, &record.FileTitle, &record.CertTime)
			if err != nil   {
				return nil, err
			}
			//t:=time.Unix(record.CertTime,0)
			//tStr:=t.Format("2006年01月02日 15:04:05")
			tStr:=utils.TimeFormat(record.CertTime,"2006年01月02日 15:04:05")
			record.CertTimeFormat = tStr
			records = append(records, record)
		}
		return records, nil
	}
