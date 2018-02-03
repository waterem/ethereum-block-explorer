package ethrpc

import (
	"github.com/onrik/ethrpc"
)

var client *ethrpc.EthRPC

func init() {
}

func InitEthRPCClient(url string) {
	client = ethrpc.NewEthRPC(url)
}

func EthBlockNumber() int {
	blockNumber, err := client.EthBlockNumber()
	if err != nil {
		panic(err)
	}
	return blockNumber
}

func EthGetBlockByNumber(blockNumber int) *ethrpc.Block {
	block, err := client.EthGetBlockByNumber(blockNumber, false)
	if err != nil {
		panic(err)
	}
	return block
}
