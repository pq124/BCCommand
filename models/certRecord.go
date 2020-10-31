package models

import (
	"bytes"
	"encoding/gob"
)

type CertRecord struct {
	CertId []byte//认证id,本质是一个md5值
	CertHash []byte //存证文件得hash值,本质是一个sha256值
	CertName string //认证人得名称
	Phone   string //联系方式
	CertCard string //身份证号
	FileName string //认证文件得名称
	FileSize int64 //文件的大小
	CertTime int64 //认证的时间


}

//序列化操作
func (c CertRecord)Serialize()([]byte,error){
	buff:=new(bytes.Buffer)
	err:=gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(),err
}

func DeserializaCertRecord(data []byte)(*CertRecord,error){
	var  certRecord *CertRecord
	err :=gob.NewDecoder(bytes.NewReader(data)).Decode(&certRecord)
	return certRecord,err
}