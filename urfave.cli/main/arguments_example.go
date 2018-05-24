package main

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"os"
	"log"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q", c.Args().Get(0))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
