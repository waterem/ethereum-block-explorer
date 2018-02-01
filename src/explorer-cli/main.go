package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"

	"./config"
	"./ethrpc"
	"./file"
	"./functions"
)

func main() {
	app := cli.NewApp()
	app.Name = "explorer-cli"
	app.Usage = "ethereum blockchain explorer cli"
	app.Action = func(c *cli.Context) error {
		start()
		return nil
	}
	app.Run(os.Args)
}

func start() {
	lastBlock, _ := strconv.Atoi(file.ReadLastBlock(config.Get("lastBlockFile")))
	lastBlock++ // TODO

	currentBlock := ethrpc.EthBlockNumber()

	fmt.Println(lastBlock)
	fmt.Println(currentBlock)

	for lastBlock <= currentBlock {
		functions.CreateBlock(lastBlock)

		lastBlock++
		fmt.Println(lastBlock)

		break
	}
}
