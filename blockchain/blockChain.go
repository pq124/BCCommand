package blockchain

import "github.com/bolt"

const BLOCKCHAIN  ="blockchain.db"
const BLOCK_NAME  ="blocks"
const LASTHASH  ="lasthash"

type BlockChain struct {
	LastHash []byte  //表示区块链中最新区块的哈希,用于查找最新的区块内容
	BoltDb  *bolt.DB //区块链中操作区块数据文件的数据库操作对象
}
/*
 *创建一个区块链
 */
func NewBlockChain()BlockChain  {
	//创世区块
	genesis := CreateGenesisBlock()
	db,err:=bolt.Open(BLOCKCHAIN,0600,nil)
	if err!=nil {
		panic(err.Error())
	}

	bc := BlockChain{
		LastHash: genesis.Hash,
		BoltDb:  db ,
	}
	//把创世区块保存到数据库文件中去
	db.Update(func(tx *bolt.Tx) error {
		bucket,err:=tx.CreateBucket([]byte(BLOCK_NAME))
		if err!=nil {
			panic(err.Error())
		}
		//序列化
		genesisBytes:=genesis.Serialize()
		//把创世区块存入桶中
		bucket.Put(genesis.Hash,genesisBytes)
		//
		bucket.Put([]byte(LASTHASH),genesis.Hash)
		return nil
	})
	return bc
}



func (bc BlockChain) AddBlock(data []byte) {
	//1.从文件中读取到最新的区块
	db := bc.BoltDb
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BLOCK_NAME))
		if bucket==nil {
			panic("读取区块链数据失败")
		}
		lastHash := bucket.Get([]byte(LASTHASH))
		lastBlockBytes :=bucket.Get(lastHash)

		lastBlock,_=DeSerialize(lastBlockBytes)
		return nil
	})


	//新建一个区块
	newBlock:=NewBlock(lastBlock.Height+1,lastBlock.Hash,data)
	//把新区块存到文件中
	db.Update()
}