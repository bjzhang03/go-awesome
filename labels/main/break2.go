package main

import "fmt"

func main() {
	a := 1
Loop:
	for j := 0; j < 3; j++ {
		for a < 10 {
			fmt.Println(a)
			if a == 5 {
				break Loop
			}
			a++
		}
	}
	fmt.Printf("a = %d", a)
}
