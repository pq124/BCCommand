package controllers

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/models"
	"DataCertPlatform/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CertDetailController struct {
	beego.Controller
}

func (c *CertDetailController)Get() {
	//1.解析和接收前端页面传递的数据cert_id
	cert_id := c.GetString("cert_id")
	//2.到区块链上查询区块数据
	block,err:=blockchain.CHAIN.QueryBlockByCertId(cert_id)
	if err!=nil {
		c.Ctx.WriteString("抱歉,查询链上数据遇到错误,请重新尝试")
		return
	}
	if block == nil {
      c.Ctx.WriteString("抱歉，未查到链上数据")
		return
	}
	fmt.Println("查询到的区块高度",block.Height)

	certRecoed,err:=models.DeserializeCertRecord(block.Data)
	certRecoed.CertIdHex = strings.ToUpper(string(certRecoed.CertId))
	certRecoed.CertHashHex = string(certRecoed.CertHash)
	certRecoed.CertTimeFormat = utils.TimeFormat(certRecoed.CertTime,utils.TIME_FORMAT_ONE)

	c.Data["CerRecord"]=certRecoed
	//3.跳转证书详情页面
	c.TplName = "cert_detail.html"
}