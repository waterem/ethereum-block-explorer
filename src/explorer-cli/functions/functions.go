package functions

import (
	"../ethrpc"
)

func init() {
}

func CreateBlock(blockNumber int) {
	ethrpc.EthGetBlockByNumber(blockNumber)
}

func CreateTransaction(blockHash string) {

}
