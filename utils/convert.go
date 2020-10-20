package utils

import (
	"bytes"
	"encoding/binary"
)

/*
 *将一个int64转化为byte字节切片
 */
func Int64ToByte(num int64)([]byte,error)  {
	//Buffer:缓冲区
	 buff:=new(bytes.Buffer)//通过new实例一个缓存区
	 //buff.Write()通过一些列的Write方法向缓存区写入数据
	 //buff.Bytes()通过Bytes方法从缓存区中获取数据
	 //大端为序排列BigEndian
	 err:=binary.Write(buff,binary.BigEndian,num)
	if err!=nil {
		return nil,err
	}
	//从缓存区中读取数据
	return buff.Bytes(),nil
}

//将字符串转化为字节切片
func StringToBytes(str string)[]byte{
	return []byte(str)
}
