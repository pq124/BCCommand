package blockchain

import (
	"errors"
	"github.com/bolt"
	"math/big"
)

const BLOCKCHAIN = "blockchain.db"
const BUCKET_NAME = "blocks"
const LASTHASH = "lasthash"
var CHAIN *BlockChain

type BlockChain struct {
	LastHash []byte   //表示区块链中最新区块的哈希,用于查找最新的区块内容
	BoltDb   *bolt.DB //区块链中操作区块数据文件的数据库操作对象
}

/*
 *创建一个区块链
 */
func NewBlockChain() *BlockChain {
	var bc *BlockChain
	//1.先打开文件
	db, err := bolt.Open(BLOCKCHAIN, 0600, nil)
	//查看chain.db文件
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil { //没有桶,要创建新桶子
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}

		lastHash := bucket.Get([]byte(LASTHASH))
		if len(lastHash) == 0 { //桶中没有lasthash记录,需要新建创世区块，并保存
			//创世区块
			genesis := CreateGenesisBlock()
			//区块序列化以后的数据
			gensisBytes := genesis.Serialize()
			//创世区块保存到blitdb中
			bucket.Put(genesis.Hash, gensisBytes)
			// 更新只想最新区块的lasthash值
			bucket.Put([]byte(LASTHASH), genesis.Hash)
			bc = &BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
		} else { //桶中已经有lasthash记录,不再需要创世区块，只需要读取即可
			lastHash := bucket.Get([]byte(LASTHASH))
			bc = &BlockChain{
				LastHash: lastHash,
				BoltDb:   db,
			}

		}
		return nil
	})
	return bc

}
//
//该放用于遍历区块链chain.db文件，并将所有的区块查出，并返回
//
func(bc BlockChain)QueryAllBlocks()([]*Block,error){

	blocks :=make([]*Block,0)
	db:=bc.BoltDb
	var err error
	//从chain.db文件查询所有的区块
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			 err = errors.New("查询区块数据失败!")
			 return err
		}
		//bucket存在
		eachHash :=bc.LastHash
		eachBig :=new(big.Int)
		zeroBig :=big.NewInt(0)//默认值零的大整数
		for  {
			//根据hash值获取对应的区块
			eachBlockBytes := bucket.Get(eachHash)

			//反序列化操作
			eachBlock,_:=DeSerialize(eachBlockBytes)

			//将遍历到每一个区块放入到切片容器中
			blocks =append(blocks,eachBlock)

			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig)==0 {//找到了创世区块
				break//跳出循环
			}
			//不满足条件,没有找到创世区块
			eachHash = eachBlock.PrevHash
		}
		return nil
	})
	return blocks,err
}


//该方法用于完成更具用户输入的区块高度查询对应的区块信息
func (bc BlockChain)QueryBlockByHeight(height int64)(*Block,error){

	if height < 0 {
		return nil,nil
	}
	db:=bc.BoltDb

	var errs error
	var eachBlock *Block
    db.View(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			//panic("读取区块数据失败!")
			errs = errors.New("读取区块数据失败!")
			return errs
		}
		eachHash:=bc.LastHash
		for   {
			//获取到最后一个区块的hash
			eachBlockBytrs :=bucket.Get(eachHash)
			//反序列化操作
			eachBlock,err := DeSerialize(eachBlockBytrs)
			if err!=nil {
               return errs
			}
			if eachBlock.Height<height {
				break
			}
			if  eachBlock.Height == height{
				break
			}
			eachHash = eachBlock.PrevHash
		}
            return nil
	})
	return eachBlock,errs
}


//保存数据在区块链中:先生成一个新区块,然后将新区块添加到区块链中去
func (bc *BlockChain) AddBlock(data []byte)(Block ,error) {
	//1.从文件中读取到最新的区块
	db := bc.BoltDb
	var lastBlock *Block

	//error的自定义
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("读取区块链数据失败")
			err = errors.New("读取区块链失败")
			return err
		}
		//lastHash := bucket.Get([]byte(LASTHASH))
		lastBlockBytes := bucket.Get(bc.LastHash)

		lastBlock ,_ =DeSerialize(lastBlockBytes)
		return nil
	})
	//新建一个区块
	newBlock := NewBlock(lastBlock.Height+1, lastBlock.Hash, data)
	//把新区块存到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		newBlockBytes := newBlock.Serialize()
		//把新创建的区块存入到boltdb数据库中
		bucket.Put(newBlock.Hash, newBlockBytes)
		//更新LASTHASH对应的值,更新为最新区块的区块的hash值
		bucket.Put([]byte(LASTHASH), newBlock.Hash)
		return nil
		//将区块链实例的LASTHASH更新成最新区块
		bc.LastHash = newBlock.Hash
		return nil
	})
	return newBlock, err
}
//该方法用于更具用户输入的认证号查询到对应的区块信息
/*func(bc BlockChain) QueryBlockByCertId(cert_id string)*Block{
	blocks :=make([]*Block,0)
	db :=bc.BoltDb
	var  err  error
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {//判断桶是否存在
			err =errors.New("查询链上数据发生的错误")
			return  err
		}

		eachHash :=bc.LastHash
		eachBig :=new(big.Int)
        zeroBig :=big.NewInt(0)

		for  {
           eachBlockBytes :=bucket.Get(eachHash)
           eachBlock,err:=DeSerialize(eachBlockBytes)
			if err!=nil {
				break
			}
			// 将遍历到得区块中得数据跟用户提供得认证号比较
			if string(eachBlock.Data)==cert_id {
				blocks :=eachBlock
				break
			}
			eachHash = eachBlock.PrevHash
			if  {

			}
		}
		return  nil

	})
	return block,nil
}
 */


