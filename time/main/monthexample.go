package main

import (
	"fmt"
	"time"
)

func main() {
	_, month, day := time.Now().Date()
	if month == time.June && day == 20 {
		fmt.Println("Happy Go day!")
	}
}