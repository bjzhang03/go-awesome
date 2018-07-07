package main

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"os"
)

var (
	Revision = "1.0.0"
)

//go run appversion_example.go --version
func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s revision=%s\n", c.App.Version, Revision)
	}

	app := cli.NewApp()
	app.Name = "party"
	app.Version = "19.99.0"
	app.Run(os.Args)
}
