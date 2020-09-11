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
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		for isPrefix{
			newline, nispre, nerr := reader.ReadLine()
			if nerr != nil {
				fmt.Println(err.Error())
				break
			}
			line= append(line, newline...)
			isPrefix=nispre
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
