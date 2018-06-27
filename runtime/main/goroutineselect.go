package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}

}