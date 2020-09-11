package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	//打开文件
	file, err := os.Open("aa.txt")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	// 结束关闭文件流
	defer file.Close()

	gin:=bufio.NewScanner(strings.NewReader("111 \n222"))
	for gin.Scan(){
		fmt.Printf("your input is %s \n",gin.Text())
		if gin.Text()=="stop"{
			break
		}
	}
}
