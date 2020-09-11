package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There ware errors reading,exiting program.")
		return
	}
	fmt.Printf("Your name is %s", input)

	input2, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There ware errors reading,exiting program.")
		return
	}
	fmt.Printf("Your name is %s", input2)
	//对unix:使用“\n”作为定界符，而window使用"\r\n"为定界符
	//Version1
	/*
	   switch input {
	   case "Philip\r\n":
	       fmt.Println("Welcome Philip!")
	   case "Chris\r\n":
	       fmt.Println("Welcome Chris!")
	   case "Ivo\r\n":
	       fmt.Println("Welcome Ivo!")
	   default:
	       fmt.Println("You are not welcome here! Goodbye!")
	   }
	*/
	//version2
	/*
	   switch input {
	   case "Philip\r\n":
	       fallthrough
	   case "Ivo\r\n":
	       fallthrough
	   case "Chris\r\n":
	       fmt.Printf("Welcome %s\n", input)
	   default:
	       fmt.Printf("You are not welcome here! Goodbye!")
	   }
	*/
	//version3
	switch input {
	case "Philip\r\n", "Ivo\r\n", "Chris\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}
}
