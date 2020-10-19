package blockchain

import "time"

/*
 *定义区块结构体,用于表示区块
 */
type Block struct {
     Height int64 //区块的高度,第几个区块
     TimeStamp int64 // 区块产生的时间戳
     PrevHash []byte //前一个区块的Hash
     Data  []byte //数据字段
     Hash  []byte //当前区块的Hash值
     Version string //版本号
}

func NewBlock(height int64,prevhash []byte,data []byte) Block {
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevhash,
		Data:      data,
		Version:   "0x01",
	}
	
	 return block
}
//创建创世区块
func CreateGenesisBlock()Block  {
	genesisBlock:=NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},nil)
	return genesisBlock
}