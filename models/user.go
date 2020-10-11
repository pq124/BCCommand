package models

import (
	"DataCertPlatform/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct{
	Id int  `form:"id"`
	Telephone string
	Password string
}


func (u User) SaveUser()(int64,error){
	//1.密码脱敏处理
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)
	u.Password =hex.EncodeToString(passwordBytes)
	//2.执行数据库操作
	row,err :=db_mysql.Db.Exec("insert into user (telephone,password)+values(?,?),u.Telephone,u.Password")
	if err != nil {
		return -1,err
	}
     numbers,err :=row.RowsAffected()
	if err != nil {
		return-1,err
	}
	return numbers,nil
}