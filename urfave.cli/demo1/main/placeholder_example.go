package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"log"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}


	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
