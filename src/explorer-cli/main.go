package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"./config"
	"./file"
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
	lastBlock := file.ReadLastBlock(config.Get("lastBlockFile"))
	currentBlock := file.ReadLastBlock(config.Get("lastBlockFile"))
	fmt.Println(lastBlock)
	fmt.Println(currentBlock)
}
