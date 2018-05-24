package main

import (
	"fmt"
	"time"
)

//https://blog.csdn.net/liuxinmingcode/article/details/49507991
//实现多个Goroutine之间的同步
func main(){
	unbufChan := make(chan int)
	//unbufChan := make(chan int, 1) 有缓冲容量

	//启用一个Goroutine接收元素值操作
	go func(){
		fmt.Println("Sleep a second...")
		time.Sleep(time.Second)//休息1s
		num := <- unbufChan //接收unbufChan通道元素值
		fmt.Printf("Received a integer %d.\n", num)
	}()

	num := 1
	fmt.Printf("Send integer %d...\n", num)
	//发送元素值
	unbufChan <- num
	fmt.Println("Done.")
}