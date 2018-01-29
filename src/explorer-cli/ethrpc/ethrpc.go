package ethrpc

import (
	"fmt"
	"log"

	"github.com/onrik/ethrpc"
)

func init() {
}

func Example() {
	client := ethrpc.NewEthRPC("")

	version, err := client.EthBlockNumber()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)
}
