package main

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"os"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
