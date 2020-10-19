package utils

import (
	"io"
	"os"
)

func SavaFile(fileName string,file io.Reader)(int64,error){
	savaFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 777)
	if err != nil {

		return-1 ,err
	}

	length, err := io.Copy(savaFile, file)
	if err != nil {
		return -1,err
	}
	return length,nil
}