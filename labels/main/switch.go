package main

import "fmt"

func main() {
SwitchStatement:
	switch 1 {
	case 1:
		fmt.Println(1)
		for i := 0; i < 10; i++ {
			break SwitchStatement
		}
		fmt.Println(2)
	}
	fmt.Println(3)
}
