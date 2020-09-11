package main

import (
	"log"
	"os"
)

func init()  {
	log.SetOutput(os.Stdout)
}

func main()  {
	log.Println("111")
}
