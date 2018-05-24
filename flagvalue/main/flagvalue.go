package main

import (
	"flag"
	"time"
	"fmt"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

//go run flagvalue.go -period 50s
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()

}
