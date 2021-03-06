package models

import (
	"DataCertPlatform/db_mysql"
	"DataCertPlatform/utils"
	"fmt"
)

type User struct {
	Id        int    `form:"id"`
	Telephone string `form:"telephone"`
	Password  string `form:"password"`
	Name      string `form:"name"` //名字
	Card      string `form:"card"` //身份证号
	Sex       string `form:"sex"`  //性别
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

func (u User) UpdateUser() (int64, error) {
	fmt.Println("电话号码为:", u.Telephone)
	rs, err := db_mysql.Db.Exec("update user set name = ?, card = ?, sex = ? where telephone = ?", u.Name, u.Card, u.Sex, u.Telephone)
	if err != nil {
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (u User) AddUser() (int64, error) {
	fmt.Println(u.Telephone, u.Password)

	//md5Hash := md5.New()
	//md5Hash.Write([]byte(u.Password))
	//passwordBytes := md5Hash.Sum(nil)
	//u.Password = hex.EncodeToString(passwordBytes)

	u.Password = utils.Md5HashString(u.Password)
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
	u.Password = utils.Md5HashString(u.Password)
	row := db_mysql.Db.QueryRow("select telephone, name, card from user where telephone = ? and password = ?",
		u.Telephone, u.Password)
	err := row.Scan(&u.Telephone, &u.Name, &u.Card)
	if err != nil {
		return nil, err
	}
	fmt.Println("数据库读取到的用户数据：", u.Telephone)
	return &u, nil
}

func (u User) QueryUserByPhone() (*User, error) {
	row := db_mysql.Db.QueryRow("select id,name ,card,telephone from user where telephone = ?", u.Telephone)
	var user User
	err := row.Scan(&user.Id,&user.Name,&user.Card,&user.Telephone)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
