package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"

	"./config"
	"./ethrepository"
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

	ethrepository.InitEthRPCClient(os.Getenv("ETHEREUM_RPC"))

	lastBlock, _ := strconv.Atoi(file.ReadLastBlock(config.Get("lastBlockFile")))
	lastBlock++

	// 1397553
	if lastBlock < 1397553 {
		lastBlock = 1397553
	}

	currentBlock := ethrepository.EthBlockNumber()

	log.Println("lastBlock=?, currentBlock=?", lastBlock, currentBlock)

	for lastBlock <= currentBlock {
		log.Println(lastBlock)
		functions.CreateBlock(lastBlock)

		lastBlock++

		// 1500000
		if lastBlock > 1500000 {
			break
		}
	}
}
