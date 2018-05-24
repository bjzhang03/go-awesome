package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"log"
	"fmt"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "language for the greeting",
			EnvVar: "LEGACY_COMPAT_LANG,APP_LANG,LANG",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("error happpened!")
		log.Fatal(err)
	}
}