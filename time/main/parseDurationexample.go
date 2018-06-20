package main

import (
	"fmt"
	"time"
)

func main() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("there are %.0f seconds in %v\n", complex.Seconds(), complex)
}
