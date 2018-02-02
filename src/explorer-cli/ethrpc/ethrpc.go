package ethrpc

import (
	"github.com/onrik/ethrpc"
)

func init() {
}

func EthBlockNumber() int {
	client := ethrpc.NewEthRPC("")

	blockNumber, err := client.EthBlockNumber()
	if err != nil {
		panic(err)
	}

	return blockNumber
}

func EthGetBlockByNumber(blockNumber int) *ethrpc.Block {
	client := ethrpc.NewEthRPC("")

	block, err := client.EthGetBlockByNumber(blockNumber, false)
	if err != nil {
		panic(err)
	}

	return block
}
