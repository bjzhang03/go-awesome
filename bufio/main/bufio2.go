package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	gin:=bufio.NewScanner(os.Stdin)
	for gin.Scan(){
		fmt.Println(gin.Text())
		if gin.Text()=="stop"{
			break
		}
	}
}
