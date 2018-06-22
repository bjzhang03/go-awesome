package main

import (
	"os"
	stdlog "log"
	kitlog "github.com/go-kit/kit/log"
)

func main() {
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))
	stdlog.Print("I sure like pie")
}