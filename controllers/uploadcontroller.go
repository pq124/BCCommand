package controllers

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/models"
	"DataCertPlatform/utils"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"time"
)

type UploadController struct {
	beego.Controller
}


    //该Post方法用于处理客户端提交的认证文件
func (u *UploadController) Post() {
	//1.解析用户上传的文件解析
	//用户上传的自定义文件名
	telephone := u.Ctx.Request.PostFormValue("telephone")
	title := u.Ctx.Request.PostFormValue("upload_title") //获取用户输入的标题
	fmt.Println(title)

	//用户上传的文件
	file, head, err := u.GetFile("pengqiang")
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("数据解析失败,请重试")
		return
	}
	defer file.Close() //延迟执行 invalid memory or nil pointer derfence 空指针错误

	//调用工具保存在本地
	savaFilePath := "static/upload/" + head.Filename
	_, err = utils.SavaFile(savaFilePath,file)
	if err!=nil {
		u.Ctx.WriteString("抱歉,文件数据认证失败")
		return
	}
	//使用io包提供的方法保存文件
	//savaFile, err := os.OpenFile(savaFilePath, os.O_CREATE|os.O_RDWR, 777)
	//if err != nil {
		//u.Ctx.WriteString("抱歉,电子数据认证失败,请重试!")
	//}

	//_, err = io.Copy(savaFile, file)
	//if err != nil {
	//	u.Ctx.WriteString("抱歉,电子数据认证失败,请重新尝试!")
	//	return
	//}

	//计算SHA256
	//hash256 := sha256.New()
	//fileBytes, _ := ioutil.ReadAll(file)
	//hash256.Write(fileBytes)
	//hashBytes := hash256.Sum(nil)
	//fmt.Println(hex.EncodeToString(hashBytes))
	fileHash,err:=utils.SHA256HashReader(file)
	fmt.Println(fileHash)

	fmt.Println(telephone)
	user1,err := models.User{Telephone:telephone}.QueryUserByPhone()
	//user1,err := user.QueryUserByPhone()
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉电子数据认证失败,请重试")
		return
	}

	//md5Hash := md5.New()
	//fileMd5Bytes, err := ioutil.ReadAll(savaFile)
	//md5Hash.Write(fileMd5Bytes)
	//bytes := md5Hash.Sum(nil)
	savaFile,err:=os.Open(savaFilePath)
	md5String,err:=utils.MD5HashReader(savaFile)
	if err!=nil {
		u.Ctx.WriteString("抱歉电子数据认证失败.")
		return
	}

	record := models.UploadRecord{
		UserId:    user1.Id,
		FileName:  head.Filename,
		FileSize:  head.Size,
		FileCert: md5String,
		FileTitle: title,
		CertTime:  time.Now().Unix(),
	}

	_, err = record.SavaRecord()
	if err != nil {
		u.Ctx.WriteString("抱歉,数据存储失败,请重试")
		return
	}

   user:=&models.User{
   Telephone:telephone,
   }
     user,_= user.QueryUserByPhone()
     CertRecord :=models.CertRecord{
		CertId:   []byte(md5String),
		CertHash: []byte(fileHash),
		CertName: user.Name,
		CertCard: user.Card,
		Telephone:user.Telephone,
		FileName: head.Filename,
		FileSize: head.Size,
		CertTime: time.Now().Unix(),
	}
	 certBytes,_:=CertRecord.Serialize()
	//将用户上传的文件的md5值和sha256值保存到区块链上
	block,err:=blockchain.CHAIN.AddBlock([]byte(certBytes))
	if err!=nil {
		u.Ctx.WriteString("抱歉数据上链失败"+err.Error())
		return
	}
	fmt.Println("恭喜,已经保存到区块两种，区块的高度是:",block.Height)


	//先查询用户id
    records,err:=models.QueryRecordsByUserId(user1.Id)
	if err!=nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("查询用户失败,请重试")
		return
	}
	u.Data["Records"] = records
	u.TplName ="list_record.html"
	//
	//u.Ctx.WriteString("恭喜,已经接收到上传文件!")

	/*储存数据库操作

		var upload models.UploadRecord
		err = u.ParseForm(&upload)
		if err!=nil {
			u.Ctx.WriteString("文件解析失败，请重新尝试")
			return
		}
		_,err = upload.SavaRecord()
		if err!=nil {
			fmt.Println(err.Error())
			u.Ctx.WriteString("文件储存失败")
			return
		}
	     u.Ctx.WriteString("文件储存成功")
	}

	*/
	//方法二 上传文件
	/*
	fileNameSlice := strings.Split(head.Filename,".")
	fileType := fileNameSlice[1]
	fmt.Println(fileNameSlice)
	fmt.Println(":",fileType)

	isJpg := strings.HasSuffix(head.Filename,".jpg")
	isPng := strings.HasSuffix(head.Filename,".png")
	if !isJpg && !isPng{
	//文件类型不支持
	u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
	return
	}

	//if fileType != " jpg" || fileType != "png" {
	//	//文件类型不支持
	//	u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
	//	return
	//}

	//文件的大小 200kb
	config := beego.AppConfig
	fileSize,err := config.Int64("file_size")

	if head.Size / 1024 > fileSize {
	u.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
	return
	}
	fmt.Println("标题:",title)
	fmt.Println("上传的文件名称:",head.Filename)
	fmt.Println("上传的文件大小:",head.Size)


	//fromFile:文件
	//toFile:要保存的文件路径
	//premission权限
	//权限的组成:
	//a+b+c
	//a:文件所有者对文件的操作权限   读4 写2 执行1
	//b:文件所有者所在组的用户的操作权限  读4 写2 执行1
	//c:其他用户的操作权限  读4 写2 执行1

	//创建upload文件
	saveDir := "static/upload"
	//1.打开文件
	_,err =os.Open(saveDir)

	//flag:文件的操作项

	//os.OpenFile("文件名",os.O_CREATE|os.O_RDWR)
	if err!=nil {
	//2 创建文件
	err =os.Mkdir("static/upload",777)
	if err!=nil {
	u.Ctx.WriteString("抱歉,文件认证遇到错误，请重试")
	return
	}
	}



	//ToFile :要保存的文件路径
	savaName:=saveDir+"/"+head.Filename
	fmt.Println("要保存的文件名",savaName)

	err = u.SaveToFile("pengqiang",savaName)
	if err!=nil {
	fmt.Println(err.Error())
	u.Ctx.WriteString("抱歉文件认证失败，请重试")
	return
	}


	fmt.Println("上传的文件:",file)
	u.Ctx.WriteString("获取到上传文件")

	}

	*/
}
/*
 *func (u *UploadController) Get()  {
	u.Data["Telephone"] =u.T
	u.TplName="home.html"
}
 */
