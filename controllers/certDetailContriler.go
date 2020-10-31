package controllers

import (
	"github.com/astaxie/beego"
)

type CertDetailController struct {
	beego.Controller
}

func (c *CertDetailController)Get() {
	//1.解析和接收前端页面传递的数据cert_id
	//cert_id := c.GetString("cert_id")
	//2.到区块链上查询区块数据
	//blockchain.CHAIN.QueryBlockByCertId(cert_id)
	//3.跳转证书详情页面
	//c.TplName = "cert_detail.html"
}