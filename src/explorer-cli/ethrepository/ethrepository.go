package ethrepository

import (
	"github.com/onrik/ethrpc"
)

var client *ethrpc.EthRPC

func init() {
}

func InitEthRPCClient(url string) {
	client = ethrpc.NewEthRPC(url)
	//client.Debug = true
}

func EthBlockNumber() int {
	blockNumber, err := client.EthBlockNumber()
	if err != nil {
		panic(err)
	}
	return blockNumber
}

func EthGetBlockByNumber(blockNumber int) *ethrpc.Block {
	block, err := client.EthGetBlockByNumber(blockNumber, true)
	if err != nil {
		panic(err)
	}
	return block
}

func EthGetTransactionReceipt(hash string) *ethrpc.TransactionReceipt {
	transactionReceipt, err := client.EthGetTransactionReceipt(hash)
	if err != nil {
		panic(err)
	}
	return transactionReceipt
}

func EthCall(params ethrpc.T) string {
	result, err := client.EthCall(params, "latest")
	if err != nil {
		panic(err)
	}
	return result
}
