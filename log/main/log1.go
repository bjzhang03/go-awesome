package main

import (
	"log"
	"os"
)

// 将日志输出到控制台
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("console output!")
}
