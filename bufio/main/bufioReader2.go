package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("aa.txt")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	// 结束关闭文件流
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadSlice('\n')
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("the line:%s\n", line)
	}
	//line, _ := reader.ReadSlice('\n')
	//fmt.Printf("the line:%s\n", line)
	//// 这里可以换上任意的 bufio 的 Read/Write 操作
	//n, _ := reader.ReadSlice('\n')
	//fmt.Printf("the line:%s\n", line)
	//fmt.Println(string(n))
}
