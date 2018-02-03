package functions

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"../ethrpc"
)

var db, err = sql.Open("mysql", "admin:12dlql*41@(database:3306)/explorer")

func init() {
	if err != nil {
		panic(err.Error())
	}
}

func CreateBlock(blockNumber int) {
	block := ethrpc.EthGetBlockByNumber(blockNumber)
	blockHashId := CreateIndexBlock(block.Hash)
	parentHashId := CreateIndexBlock(block.ParentHash)

	rows, err := db.Query("SELECT number FROM blocks WHERE number=? LIMIT 1", blockNumber)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			panic(err.Error())
		}

		result, err := db.Exec("UPDATE blocks SET hash_id=?, parent_hash_id=?, timestamp=?, hash=? where number=?", blockHashId, parentHashId, block.Timestamp, block.Hash, blockNumber)
		if err != nil {
			panic(err.Error())
		}

		id, err = result.LastInsertId()
		if err != nil {
			panic(err.Error())
		}
	}

	result, err := db.Exec("INSERT INTO blocks (number, hash_id, parent_hash_id, timestamp) values (?,?,?,?)", blockNumber, blockHashId, parentHashId, block.Timestamp)
	if err != nil {
		panic(err.Error())
	}

	_, err = result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	// TODO check each transaction
}

func CreateIndexBlock(blockHash string) int64 {

	rows, err := db.Query("SELECT id FROM index_blocks WHERE `hash`=? LIMIT 1", blockHash)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		return id
	}

	result, err := db.Exec("INSERT INTO index_blocks (`hash`) values (?)", blockHash)
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return id
}
