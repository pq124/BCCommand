package main

import (
	_ "DataCertPlatform/routers"
	"HelloBeego/db_myseq"
	"github.com/astaxie/beego"
)

func main() {

	db_myseq.Connect()

	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()

}

