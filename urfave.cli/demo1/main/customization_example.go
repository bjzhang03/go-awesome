package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
)

//go run customization_example.go wat --compgen
func main() {
	cli.BashCompletionFlag = cli.BoolFlag{
		Name:   "compgen",
		Hidden: false,
	}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name: "wat",
		},
	}
	app.Run(os.Args)
}