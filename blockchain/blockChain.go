package blockchain

import "github.com/bolt"

const BLOCKCHAIN  ="blockchain.db"
const BUCKET_NAME  ="blocks"
const LASTHASH  ="lasthash"

type BlockChain struct {
	LastHash []byte  //表示区块链中最新区块的哈希,用于查找最新的区块内容
	BoltDb  *bolt.DB //区块链中操作区块数据文件的数据库操作对象
}
/*
 *创建一个区块链
 */
func NewBlockChain()BlockChain  {
      var bc BlockChain
      //1.先打卡文件
	db,err:=bolt.Open(BLOCKCHAIN,0600,nil)

    //查看chain.db文件
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket ==nil {//没有桶,要创建新桶子
			bucket,err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err!=nil {
				panic(err.Error())
			}
		}
		//
		lastHash := bucket.Get([]byte(LASTHASH))
		if len(lastHash)==0 {//桶中没有lasthash记录,需要新建创世区块，并保存
			//创世区块
			genesis := CreateGenesisBlock()
			//区块序列化以后的数据
			gensisBytes :=genesis.Serialize()
			//创世区块保存到blitdb中
			bucket.Put(genesis.Hash,gensisBytes)
			// 更新只想最新区块的lasthash值
			bucket.Put([]byte(LASTHASH),genesis.Hash)
			bc = BlockChain{
				LastHash: genesis.Hash,
				BoltDb:  db ,
			}
		}else {//桶中已经有lasthash记录,不再需要创世区块，只需要读取即可
			lastHash := bucket.Get([]byte(LASTHASH))
			bc =BlockChain{
				LastHash:lastHash ,
				BoltDb: db ,
			}

		}
		return nil
	})
      return bc



	//把创世区块保存到数据库文件中去
	/*
	 *
	 */
	//db.Update(func(tx *bolt.Tx) error {
		//bucket,err:=tx.CreateBucket([]byte(BUCKET_NAME))
		//if err!=nil {
		//	panic(err.Error())
		//}
		//序列化
		//genesisBytes:=genesis.Serialize()
		//把创世区块存入桶中
		//bucket.Put(genesis.Hash,genesisBytes)
		//
		//bucket.Put([]byte(LASTHASH),genesis.Hash)
		//return nil
	//})
	//return bc
}



func (bc BlockChain) AddBlock(data []byte) {
	//1.从文件中读取到最新的区块
	db := bc.BoltDb
	var lastBlock *Block
		db.View(func(tx *bolt.Tx) error {
			bucket:=tx.Bucket([]byte(BUCKET_NAME))
			if bucket==nil {
				panic("读取区块链数据失败")
			}
			//lastHash := bucket.Get([]byte(LASTHASH))
			lastBlockBytes :=bucket.Get(bc.LastHash)

			lastBlock,_=DeSerialize(lastBlockBytes)
			return nil
	})

	//新建一个区块
	newBlock:=NewBlock(lastBlock.Height+1,lastBlock.Hash,data)
	//把新区块存到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		newBlockBytes :=newBlock.Serialize()
		//把新创建的区块存入到boltdb数据库中
		bucket.Put(newBlock.Hash,newBlockBytes)
		//更新LASTHASH对应的值,更新为最新区块的区块的hash值
		bucket.Put([]byte(LASTHASH),newBlock.Hash)
		return nil
		 //将区块链实例的LASTHASH
		bc.LastHash=newBlock.Hash
		return nil
	})

}