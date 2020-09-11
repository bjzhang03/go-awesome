package main

import (
	"bufio"
	"fmt"
	"os"
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

	gin:=bufio.NewScanner(file)
	for gin.Scan(){
		fmt.Printf("your input is %s \n",gin.Text())
		if gin.Text()=="stop"{
			break
		}
	}
}

