package main

import "fmt"

func main() {
	var a int = 10
LABEL:
	for a < 20 {
		if a == 15 {
			a = a + 1
			continue LABEL
		}
		fmt.Println(a)
		a++
	}
}
