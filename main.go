package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
     block0:=blockchain.CreateGenesisBlock()
     block1:=blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
     //block2:=blockchain.NewBlock(block1.Height+1,block1.Hash,[]byte("a"))
     fmt.Println("block0为:",block0,"block1为:",block1)
	return


	db_mysql.Connect()
	//静态资源文件设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()

}

