package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := cli.NewApp()
	app.Name = "explorer-cli"
	app.Usage = "get ethereum blockchain data and import into mysql database"
	app.Action = func(c *cli.Context) error {
		start()
		return nil
	}

	app.Run(os.Args)
}

func start() {

	db, err := sql.Open("mysql", "admin:12dlql*41@(database:3306)/explorer")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM index_blocks")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(columns[0])
}
