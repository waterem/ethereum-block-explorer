package main

import (
	"log"
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

	log.Println(lastBlock)
	log.Println(currentBlock)

	for lastBlock <= currentBlock {
		functions.CreateBlock(lastBlock)

		lastBlock++
		log.Println(lastBlock)

		break
	}
}
