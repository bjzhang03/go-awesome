package main

import "fmt"

func main() {
OuterLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i =%v, j =%v\n", i, j)
			break OuterLoop
		}
	}
}
