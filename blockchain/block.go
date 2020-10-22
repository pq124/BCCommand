package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"
)

/*
 *定义区块结构体,用于表示区块
 */
type Block struct {
	Height    int64  //区块的高度,第几个区块
	TimeStamp int64  // 区块产生的时间戳
	PrevHash  []byte //前一个区块的Hash
	Data      []byte //数据字段
	Hash      []byte //当前区块的Hash值
	Version   string //版本号
	Nonce     int64  //区块对应的nonce值
}

func NewBlock(height int64, prevhash []byte, data []byte) Block {
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevhash,
		Data:      data,
		Version:   "0x01",
	}
	//找到nonce值
	pow := NewPoW(block)
	hash,nonce := pow.Run()
	block.Nonce = nonce
	block.Hash= hash

	//heightBytes, _ := utils.Int64ToByte(block.Height)
	//timeStampBytes, _ := utils.Int64ToByte(block.TimeStamp)
	//versionBytes := utils.StringToBytes(block.Version)
	//nonceBytes, _ := utils.Int64ToByte(block.Nonce)

	//var blockBytes []byte

	//blockBytes = bytes.Join([][]byte{
	//	heightBytes,
	//	timeStampBytes,
	//	block.PrevHash,
	//	block.Data,
	//	versionBytes,
	//	nonceBytes,
	// []byte{})
	//计算Hash值
	//block.Hash = utils.SHA256HashBlock(blockBytes)

	//挖矿竞争,获得记账权
	return block
}

//创建创世区块
func CreateGenesisBlock() Block {
	genesisBlock := NewBlock(0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil)
	return genesisBlock
}

func (b Block)Serialize() ([]byte){
	buff := new(bytes.Buffer)
	encoder := gob.NewEncoder(buff)
	encoder.Encode(b)//将区块B放入到序列化编码器中
    return buff.Bytes()
}

func DeSerialize (data []byte)(*Block,error){
	var  block Block
	decode := gob.NewDecoder(bytes.NewReader(data))
	err:=decode.Decode(block)
	if err!=nil {
		return nil,err
	}
	return &block,nil
}