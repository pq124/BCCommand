package blockchain

import "math/big"

type ProofOfWork struct {
   Target *big.Int //系统的目标值
   Block Block//要找的nonce值对应的区块
}

//
//SHA256(区块A + n )<系统B
//
//实例化一个PoW算法的实例
func NewPoW(block Block) ProofOfWork{
	t:=big.NewInt(1)

	t =t.Lsh(t,255)

	pow:= ProofOfWork{
		Target: t,
		Block:block,

	}
	return pow
}
