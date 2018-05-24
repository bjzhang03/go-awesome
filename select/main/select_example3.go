package main

import (
	"fmt"
)

//随机执行的顺序
//如果同时有多个case满足条件，通过一个伪随机的算法决定哪一个case将会被执行。
func main() {
	chanCap := 5
	ch7 := make(chan int, chanCap)

	for i := 0; i < chanCap; i++ {
		select {
		case ch7 <- 1:
		case ch7 <- 2:
		case ch7 <- 3:
		}
	}

	for i := 0; i < chanCap; i++ {
		fmt.Printf("%v\n", <-ch7)
	}
}
