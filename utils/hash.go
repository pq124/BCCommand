package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

/*
 *对一个字符串数据进行hash计算
 */
//Hash计算
func Md5HashString(data string)(string)  {
	   md5Hash := md5.New()
	   md5Hash.Write([]byte(data))
	   bytes := md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
/*
 *读取io流中的数据,并对数据进行hash计算,返回sha256 hash值
 */
func SHA256HashReader(reader io.Reader) (string,error) {
	sha256Hash:=sha256.New()
	readerBytes ,err:=ioutil.ReadAll(reader)
	if err!=nil {
		return "",err
	}
	sha256Hash.Write(readerBytes)
	hashBytes:=sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}

//io:input output
func MD5HashReader(reader io.Reader)(string,error)  {
	md5Hash:=md5.New()
	readerbytes,err :=ioutil.ReadAll(reader)
	fmt.Println("读取到的文件:",readerbytes)
	if err!=nil {
		return "",err
	}
	md5Hash.Write(readerbytes)
	hashBytes :=md5Hash.Sum(nil)
	  return hex.EncodeToString(hashBytes),nil
}

/*
 *对区块数据进行SHA256哈希计算
 */
func SHA256HashBlock(blockBytes []byte) []byte {
 //1.将block结构体数据转化为[]byte字节切片
 //2.将转换后的[]byte字节切片输入Write方法
	sha256Hash := sha256.New()
	sha256Hash.Write(blockBytes)
	hash:=sha256Hash.Sum(nil)
	return hash
}