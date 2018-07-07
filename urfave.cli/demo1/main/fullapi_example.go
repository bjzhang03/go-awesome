package main

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"io/ioutil"
	"io"
	"time"
	"os"
	"errors"
	"flag"
)

func init() {
	cli.AppHelpTemplate += "\nAppHelpTemplate\n"
	cli.CommandHelpTemplate += "\nCommandHelpTemplate\n"
	cli.SubcommandHelpTemplate += "\nSubcommandHelpTemplate\n"

	cli.HelpFlag = cli.BoolFlag{Name: "helpflag"}
	cli.BashCompletionFlag = cli.BoolFlag{Name: "bashcompletionflag", Hidden: true}
	cli.VersionFlag = cli.BoolFlag{Name: "version"}

	//go run fullapi_example.go help调用这一段的函数
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		fmt.Fprintf(w, "helpprinter test\n")
	}

	//go run fullapi_example.go --version 调用这一段的函数
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "version = %s\n", c.App.Version)
	}
	cli.OsExiter = func(c int) {
		fmt.Fprintf(cli.ErrWriter, "refusing to exit %d\n", c)
	}
	cli.ErrWriter = ioutil.Discard
	cli.FlagStringer = func(fl cli.Flag) string {
		return fmt.Sprintf("\t\t%s", fl.GetName())
	}
}

type hexWriter struct{}

func (w *hexWriter) Write(p []byte) (int, error) {
	for _, b := range p {
		fmt.Printf("%x", b)
	}
	fmt.Printf("\n")

	return len(p), nil
}

type genericType struct {
	s string
}

func (g *genericType) Set(value string) error {
	g.s = value
	return nil
}

func (g *genericType) String() string {
	return g.s
}

var stringflag string

func main() {
	app := cli.NewApp()
	app.Name = "urfaveclitest"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "cliexample",
			Email: "bjzhang03@foxmail.com",
		},
	}

	app.Copyright = "(c) 2018 bjzhang"
	app.HelpName = "cli demo"
	app.Usage = "urfave cli demo"
	app.UsageText = "urfave cli demo API"
	app.ArgsUsage = "[args and such]"
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "subcommand",
			Aliases:     []string{"suc"},
			Category:    "motion",
			Usage:       "subcommand test",
			UsageText:   "subcommand - test the cli subcommand function",
			Description: "subcommand test example",
			ArgsUsage:   "[sub]",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "forever, forevvarr"},
			},
			//go run fullapi_example.go subcommand test
			//Subcommands: cli.Commands{
			//	cli.Command{
			//		Name:   "test",
			//		Action: testAction,
			//	},
			//},
			//SkipFlagParsing: false,
			//HideHelp:        false,
			//Hidden:          false,
			//HelpName:        "command",
			//BashComplete: func(c *cli.Context) {
			//	fmt.Fprintf(c.App.Writer, "bashcomplete\n")
			//},
			//go run fullapi_example.go subcommand
			Before: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "operation before subcommand\n")
				return nil
			},
			//go run fullapi_example.go subcommand
			After: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "operation after subcommand\n")
				return nil
			},
			//go run fullapi_example.go subcommand
			Action: func(c *cli.Context) error {
				c.Command.FullName()
				c.Command.HasName("test")
				c.Command.Names()
				c.Command.VisibleFlags()
				fmt.Fprintf(c.App.Writer, "operation in subcommand\n")
				if c.Bool("forever") {
					c.Command.Run(c)
				}
				return nil
			},
			//当使用出错的时候调用的函数
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Println(err)
				fmt.Fprintf(c.App.Writer, "use subcommand wrong!\n")
				return err
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "boolflag"},
		//cli.BoolTFlag{Name: "booltflag"},
		//cli.DurationFlag{Name: "durationflag", Value: time.Second * 3},
		//cli.Float64Flag{Name: "float64flag"},
		//cli.GenericFlag{Name: "genericflag", Value: &genericType{}},
		//cli.Int64Flag{Name: "int64flag"},
		//cli.Int64SliceFlag{Name: "int64sliceflag"},
		//cli.IntFlag{Name: "intflag"},
		//cli.IntSliceFlag{Name: "intsliceflag"},
		//go run fullapi_example.go --stringflag 111
		cli.StringFlag{Name: "stringflag", Value: "123", Destination: &stringflag,},
		//cli.StringSliceFlag{Name: "stringsliceflag"},
		//cli.UintFlag{Name: "uintflag"},
		//cli.Uint64Flag{Name: "uint64flag"},
	}
	app.EnableBashCompletion = true
	app.HideHelp = false
	app.HideVersion = false
	app.BashComplete = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "bashcomplete")
	}
	//在app之前执行的操作
	app.Before = func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "operation before the app\n")
		return nil
	}
	//在app之后执行的操作
	app.After = func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "operation after the app\n")
		return nil
	}
	//命令找不到的时候运行的函数
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "command not found\n", command)
	}
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		if isSubcommand {
			return err
		}

		fmt.Fprintf(c.App.Writer, "OnUsageError: %#v\n", err)
		return nil
	}

	app.Action = func(c *cli.Context) error {
		cli.DefaultAppComplete(c)
		cli.HandleExitCoder(errors.New("handleExitCoder"))
		cli.ShowAppHelp(c)
		cli.ShowCommandCompletions(c, "urfaveclitest")
		fmt.Println("-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1-1")

		cli.ShowCommandHelp(c, "subcommand")
		cli.ShowCompletions(c)
		cli.ShowSubcommandHelp(c)
		cli.ShowVersion(c)

		fmt.Println("-------------------------------------------------------------------------------")

		categories := c.App.Categories()
		categories.AddCommand("addcommandtest", cli.Command{
			Name: "addcommandtest",
		})


		////处理flag的函数
		//if stringflag == "111" {
		//	fmt.Println("111")
		//} else {
		//	fmt.Println("222")
		//}

		for _, category := range c.App.Categories() {
			fmt.Fprintf(c.App.Writer, "category.Name = : %s\n\n", category.Name)
			fmt.Fprintf(c.App.Writer, "category.Commands =: %#v\n\n", category.Commands)
			fmt.Fprintf(c.App.Writer, "category.VisibleCommands() =: %#v\n\n", category.VisibleCommands())
		}
		fmt.Printf("c.App.Command(\"subcommand\") =: %#v\n\n\n", c.App.Command("subcommand"))
		fmt.Println("stringflag", c.String("stringflag"))
		fmt.Println("boolflag", c.Bool("boolflag"))


		// go run fullapi_example.go --boolflag
		if c.Bool("boolflag") {
			//fmt.Println("1111\n\n\n\n")
			c.App.Run([]string{"urfaveclitest", "subcommand", "test"})
		}
		fmt.Println("**************************************************************************************")


		//if c.Bool("stringflag") {
		//	fmt.Println("111111\n\n\n\n")
		//	c.App.RunAsSubcommand(c)
		//}
		c.App.Setup()
		fmt.Printf("c.App.VisibleCategories() =: %#v\n\n", c.App.VisibleCategories())
		fmt.Printf("c.App.VisibleCommands() =: %#v\n\n", c.App.VisibleCommands())
		fmt.Printf("c.App.VisibleFlags() =: %#v\n\n", c.App.VisibleFlags())

		fmt.Printf("%#v\n", c.Args().First())
		if len(c.Args()) > 0 {
			fmt.Printf("%#v\n", c.Args()[1])
		}
		fmt.Printf("c.Args().Present() =: %#v\n\n", c.Args().Present())
		fmt.Printf("c.Args().Tail() =: %#v\n\n", c.Args().Tail())

		fmt.Println("-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")


		set := flag.NewFlagSet("contrive", 0)
		nc := cli.NewContext(c.App, set, c)
		fmt.Printf("%#v\n", nc.Args())
		fmt.Printf("%#v\n", nc.Bool("bool"))
		fmt.Printf("%#v\n", nc.BoolT("BoolT"))
		fmt.Printf("%#v\n", nc.Duration("Duration"))
		fmt.Printf("%#v\n", nc.Float64("Float64"))
		fmt.Printf("%#v\n", nc.Generic("Generic"))
		fmt.Printf("%#v\n", nc.Int64("Int64"))
		fmt.Printf("%#v\n", nc.Int64Slice("Int64Slice"))
		fmt.Printf("%#v\n", nc.Int("Int"))
		fmt.Printf("%#v\n", nc.IntSlice("IntSlice"))
		fmt.Printf("%#v\n", nc.String("String"))
		fmt.Printf("%#v\n", nc.StringSlice("StringSlice"))
		fmt.Printf("%#v\n", nc.Uint("Uint"))
		fmt.Printf("%#v\n", nc.Uint64("Uint64"))
		fmt.Printf("%#v\n", nc.GlobalBool("GlobalBool"))
		fmt.Printf("%#v\n", nc.GlobalBoolT("GlobalBoolT"))
		fmt.Printf("%#v\n", nc.GlobalDuration("GlobalDuration"))
		fmt.Printf("%#v\n", nc.GlobalFloat64("GlobalFloat64"))
		fmt.Printf("%#v\n", nc.GlobalGeneric("GlobalGeneric"))
		fmt.Printf("%#v\n", nc.GlobalInt("GlobalInt"))
		fmt.Printf("%#v\n", nc.GlobalIntSlice("GlobalIntSlice"))
		fmt.Printf("%#v\n", nc.GlobalString("GlobalString"))
		fmt.Printf("%#v\n", nc.GlobalStringSlice("GlobalStringSlice"))
		fmt.Printf("%#v\n", nc.FlagNames())
		fmt.Printf("%#v\n", nc.GlobalFlagNames())
		fmt.Printf("%#v\n", nc.GlobalIsSet("GlobalIsSet"))
		fmt.Printf("%#v\n", nc.GlobalSet("GlobalSet", "GlobalSet"))
		fmt.Printf("%#v\n", nc.NArg())
		fmt.Printf("%#v\n", nc.NumFlags())
		fmt.Printf("%#v\n", nc.Parent())

		nc.Set("Set", "Set")

		ec := cli.NewExitError("NewExitError", 86)
		fmt.Fprintf(c.App.Writer, "%d\n", ec.ExitCode())
		fmt.Printf("Success\n")
		return nil
	}
	//获取环境变量的操作
	if os.Getenv("TEST") != "" {
		app.Writer = &hexWriter{}
		app.ErrWriter = &hexWriter{}
	}

	app.Metadata = map[string]interface{}{
		"numbertest": "many",
		"booltest":   false,
		"othertest":  19.99,
	}

	// ignore error so we don't exit non-zero and break gfmrun README example tests
	_ = app.Run(os.Args)
}

//go run fullapi_example.go subcommand test
func testAction(c *cli.Context) error {
	fmt.Fprintf(c.App.Writer, "testAction\n")
	return nil
}
