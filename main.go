package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
)

func main() {
	 /*
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
	      block0Bytes := block0.Serialize()
	      fmt.Println("创世区块gob序列化:",block0Bytes)
	      deblock0,err :=blockchain.DeSerialize(block0Bytes)
	 	if err!=nil {
	 		fmt.Println(err.Error())
	 		return
	 	}
	 	fmt.Println("反序列化后区块的高度是:",deblock0.Height)
	 	return

	      //序列化
	      //序列化:将数据从内存中形式转换为可以持久化储存在硬盘上或者在网络上传输的形式,称之为序列化
	      //反序列化:将数据从文件中或者网络中读取，然后转化到计算机内存中的过程称之为反序列化
	      //序列化和反序列化Marshal,unMarshal
	      //只有进行序列化才能进行传输
	      //xml<>
	      //gob包
	      blockJson,_ :=json.Marshal(block0)
	      fmt.Println("JsonBlock1:",string(blockJson))

	 */
	//bc:=blockchain.NewBlockChain()//封装
	//fmt.Printf("最新区块的哈希值:%x\n 0",bc.LastHash)
	//block1,err :=bc.AddBlock([]byte("用户药保存到区块中的数据"))
	//if err!=nil {
	///	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Printf("区块高度:n",block1.)
	//return

	 blockchain.NewBlockChain()


	db_mysql.Connect()
	//静态资源文件设置
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
     
}



