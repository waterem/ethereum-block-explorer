package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func start() {
	loadEnv()

	ethrpc.InitEthRPCClient(os.Getenv("ETHEREUM_RPC"))

	lastBlock, _ := strconv.Atoi(file.ReadLastBlock(config.Get("lastBlockFile")))
	lastBlock++

	if lastBlock < 46147 {
		lastBlock = 46147
	}

	currentBlock := ethrpc.EthBlockNumber()

	log.Println("lastBlock=?, currentBlock=?", lastBlock, currentBlock)

	for lastBlock <= currentBlock {
		functions.CreateBlock(lastBlock)

		lastBlock++
		log.Println(lastBlock)

		// TODO
		if lastBlock > 46157 {
			break
		}
	}
}
