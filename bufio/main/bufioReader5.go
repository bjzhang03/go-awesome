package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		result,rerr:=strconv.Atoi(string(line))
		if rerr!=nil{
			fmt.Println("Not Number")
		}else {
			fmt.Printf("Number is %d \n",result)
		}
		fmt.Printf("the line:%s\n", line)

		strline:=string(line)

		for i:=0;i< len(strline);i++{
			if unicode.IsDigit(rune(strline[i])){
				fmt.Print(string(rune(strline[i])))
			}
		}
		fmt.Println()
	}
}
