package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
     block0:=blockchain.CreateGenesisBlock()
     //block1:=blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
     //block2:=blockchain.NewBlock(block1.Height+1,block1.Hash,[]byte("a"))
     block1:=blockchain.NewBlock(
     	block0.Height+1,
     	block0.Hash,
		 []byte{})
	fmt.Println(block0)
	fmt.Printf("block0hash为:%x\n",block0.Hash)
     fmt.Printf("block1的PrevHash: %x\n",block1.PrevHash)
     //序列化
     //序列化:将数据从内存中形式转换为可以持久化储存在硬盘上或者在网络上传输的形式,称之为序列化
     //反序列化:将数据从文件中或者网络中读取，然后转化到计算机内存中的过程称之为反序列化
     //序列化和反序列化Marshal,unMarshal
     //只有进行序列化才能进行传输
     blockJson,_ :=json.Marshal(block0)
     fmt.Println("JsonBlock1:",string(blockJson))


	db_mysql.Connect()
	//静态资源文件设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()

}


