package main

import (
	"fmt"
	"time"
)

func main() {
	ns, _ := time.ParseDuration("1000ns")
	fmt.Printf("one microsecond has %d nanoseconds.", ns.Nanoseconds())
}