package functions

import (
	"database/sql"
	"log"

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
	blockHashId := CreateTransaction(block.Hash)
	log.Println(blockHashId)
}

func CreateTransaction(blockHash string) int64 {

	rows, err := db.Query("SELECT id FROM index_blocks WHERE `hash`=? LIMIT 1", blockHash)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
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
