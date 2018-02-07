package functions

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/onrik/ethrpc"

	"../ethrepository"
)

var db, err = sql.Open("mysql", "admin:12dlql*41@(database:3306)/explorer")

func init() {
	if err != nil {
		panic(err.Error())
	}
}

func CreateBlock(blockNumber int) {
	block := ethrepository.EthGetBlockByNumber(blockNumber)

	blockHashId := createIndexBlock(block.Hash)
	parentHashId := createIndexBlock(block.ParentHash)

	if rowExists("SELECT number FROM blocks WHERE number=? LIMIT 1", blockNumber) {
		_, err = db.Exec("UPDATE blocks SET hash_id=?, parent_hash_id=?, timestamp=? where number=?", blockHashId, parentHashId, block.Timestamp, blockNumber)
		if err != nil {
			panic(err.Error())
		}
	} else {
		result, err := db.Exec("INSERT INTO blocks (number, hash_id, parent_hash_id, timestamp) values (?,?,?,?)", blockNumber, blockHashId, parentHashId, block.Timestamp)
		if err != nil {
			panic(err.Error())
		}
		_, err = result.LastInsertId()
		if err != nil {
			panic(err.Error())
		}
	}

	for _, value := range block.Transactions {

		if value.Input == "0x" {
			continue
		}

		// create token
		if strings.Contains(value.Input, "0x18160ddd") &&
			strings.Contains(value.Input, "0x70a08231") &&
			strings.Contains(value.Input, "0xdd62ed3e") &&
			strings.Contains(value.Input, "0xa9059cbb") &&
			strings.Contains(value.Input, "0x095ea7b3") &&
			strings.Contains(value.Input, "0x23b872dd") {

			transactionReceipt := ethrepository.EthGetTransactionReceipt(value.Hash)

			name := ethrepository.EthCall(ethrpc.T{From: value.From, To: transactionReceipt.ContractAddress, Data: "0x06fdde03"})
			symbol := ethrepository.EthCall(ethrpc.T{From: value.From, To: transactionReceipt.ContractAddress, Data: "0x95d89b41"})

			if name != "0x" && name != "0x0000000000000000000000000000000000000000000000000000000000000001" &&
				name != "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000" &&
				symbol != "0x" && symbol != "0x0000000000000000000000000000000000000000000000000000000000000001" &&
				symbol != "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000" {

				totalSuply := ethrepository.EthCall(ethrpc.T{From: value.From, To: transactionReceipt.ContractAddress, Data: "0x18160ddd"})
				decimals := ethrepository.EthCall(ethrpc.T{From: value.From, To: transactionReceipt.ContractAddress, Data: "0x313ce567"})

				createToken(name, symbol, totalSuply, decimals, transactionReceipt.ContractAddress)

				continue
			}
		}

		// transfer
		if strings.Contains(value.Input, "0xa9059cbb") {

			method, address, balance := value.Input[0:len("0xa9059cbb")], value.Input[len("0xa9059cbb"):64], value.Input[len("0xa9059cbb")+64:]
			fmt.Println(method, address, balance)

			//i, _ := strconv.ParseInt(balance, 16, 64)
			//fmt.Println(i)

			// get decimals from DB

			// decode input into address and value
			// index_addresses
			// transfer
			// update balance
		}
	}
}

func createToken(name string, symbol string, totalSuply string, decimals string, address string) {
	nameStr := hexToStr(name)
	symbolStr := hexToStr(symbol)
	totalSuplyStr := hexToDecimal(totalSuply)
	decimalsStr := hexToDecimal(decimals)

	addressId := createIndexAddress(address)

	if rowExists("SELECT id FROM tokens WHERE address_id=? LIMIT 1", addressId) {
		_, err = db.Exec("UPDATE tokens SET name=?, symbol=?, total_suply=?, decimals=? where address_id=?", nameStr, symbolStr, totalSuplyStr, decimalsStr, addressId)
		if err != nil {
			panic(err.Error())
		}
	} else {
		result, err := db.Exec("INSERT INTO tokens (name, symbol, total_suply, decimals, address_id) values (?,?,?,?,?)", nameStr, symbolStr, totalSuplyStr, decimalsStr, addressId)
		if err != nil {
			panic(err.Error())
		}
		_, err = result.LastInsertId()
		if err != nil {
			panic(err.Error())
		}
	}
}

func hexToStr(h string) string {
	str := strings.Replace(h, "0x", "", -1)
	bs, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(bs))
}

func hexToDecimal(h string) int64 {
	r, err := strconv.ParseInt(h, 0, 64)
	if err != nil {
		panic(err)
	}
	return r
}

func createIndexBlock(blockHash string) int64 {

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
		rows.Close()
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

func createIndexAddress(address string) int64 {

	rows, err := db.Query("SELECT id FROM index_addresses WHERE `address`=? LIMIT 1", address)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		rows.Close()
		return id
	}

	result, err := db.Exec("INSERT INTO index_addresses (`address`) values (?)", address)
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return id
}

func rowExists(query string, args ...interface{}) bool {

	query = fmt.Sprintf("SELECT exists (%s)", query)

	var exists bool
	err := db.QueryRow(query, args...).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}

	return exists
}
