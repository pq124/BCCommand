package models

import "DataCertPlatform/db_mysql"

type  UploadRecord struct {
	Id int
	UserId int
	FileName string
	FileSize int
	FileCert  string
	FileTitle string
	CertTime  int
}
func (u UploadRecord) SavaRecord(){
	db_mysql.Db.Exec("insert into")
}
