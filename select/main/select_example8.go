package main

import (
	"fmt"
	"time"
)

func main() {
	startSign := make(chan int)
	for i := 0; i < 10; i++ {
		go func(num int) {
			fmt.Println("Start work : ", num)
			time.Sleep(time.Duration(i+1) * time.Second)
			startSign <- num
		}(i)
	}

	works := []int{}
	for {
		select {
		case val := <-startSign:
			fmt.Println("Finished the work : ", val)
			works = append(works, val)
		default:
			fmt.Println("Sleep 1s to wait for the works finished!")
			time.Sleep(time.Second)
		}
		if len(works) == 10 {
			fmt.Println("All works finished!")
			break
		}
	}

}
