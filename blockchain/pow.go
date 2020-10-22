package blockchain

import (
	"DataCertPlatform/utils"
	"bytes"
	"fmt"
	"math/big"
)

const DIFFICULTY = 2


type ProofOfWork struct {
	Target *big.Int //系统的目标值
	Block  Block    //要找的nonce值对应的区块
}

//
//SHA256(区块A + n )<系统B
//
//实例化一个PoW算法的实例

func NewPoW(block Block) ProofOfWork {
	t := big.NewInt(1)
	t = t.Lsh(t, 255-DIFFICULTY)
	pow := ProofOfWork{
		Target: t,
		Block: block,
	}
	return pow
}

/*
 *run方法用于寻找合适的nonce值
 */
func (p ProofOfWork) Run() ([]byte,int64) {
	var nonce int64
	   nonce = 0
   var blockHash []byte
	for{
		block := p.Block
		heightBytes, _ := utils.Int64ToByte(block.Height)
		timeStampBytes, _ := utils.Int64ToByte(block.TimeStamp)
		versionBytes := utils.StringToBytes(block.Version)

		//已有区块信息和尝试的nonce的拼接信息
		nonceBytes,_ := utils.Int64ToByte(nonce)

		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeStampBytes,
			block.PrevHash,
			block.Data,
			versionBytes,
			nonceBytes,
		}, []byte{})
		//区块和尝试的nonce值拼接后得到的hash值
		blockHash = utils.SHA256HashBlock(blockBytes)

		target := p.Target //目标值
		//指针的定义不会帮你开辟空间
		var hashBig *big.Int  //声明和定义
		hashBig = new(big.Int) //分配内存空间,为变量分配地址

		hashBig = hashBig.SetBytes(blockHash)
		//xx :invalid memory or nil pointer dereference 空指针错误
		fmt.Println("该nonce值为:",nonce)
		if hashBig.Cmp(target) == -1 {
			break
		}
		nonce++
	}
	//将符合规则的nonce值返回
	return blockHash,nonce
}
