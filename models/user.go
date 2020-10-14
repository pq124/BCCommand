package models

import (
	"DataCertPlatform/db_mysql"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	Id        int    `form:"id"`
	Telephone string `form:"telephone"`
	Password  string `form:"password"`
}

/*
func (u User) AddUser() (int64,error){
	//脱敏
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes) //把脱敏的密码的md5值重新赋值为密码进行存储

	rs, err := db_mysql.Db.Exec("insert into user(phone,password) values(?,?)",
		u.Phone, u.Password)
	//错误早发现早解决
	if err != nil {//保存数据遇到错误
		return -1,err
	}
	id, err :=rs.RowsAffected()
	if err != nil {//保存数据遇到错误
		return -1,err
	}
	//id值代表的是此次数据操作影响的行数,id是一个整数int64类型
	return id,nil
}}
*/

func (u User) AddUser() (int64, error) {
	fmt.Println(u.Telephone, u.Password)
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordBytes)
	result, err := db_mysql.Db.Exec("insert into user (telephone, password)"+"values(?,?)", u.Telephone, u.Password)
	if err != nil {
		return -1, err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return row, nil

}

/*func (u User) QueryUser() (*User,error){

	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes) //把脱敏的密码的md5值重新赋值为密码进行存储

	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
		u.Phone, u.Password)

	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,nil
}
*/

func (u User) QueryData() (*User, error) {
	fmt.Println(u.Telephone, u.Password)
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes)

	row := db_mysql.Db.QueryRow("select telephone from user where telephone = ? and password = ?",
		u.Telephone, u.Password)

	err := row.Scan(&u.Telephone)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
