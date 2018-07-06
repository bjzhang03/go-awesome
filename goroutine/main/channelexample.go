package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	// 启动一个协程然后调用匿名函数直接进行运行
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Esc the monitoring,Stoping...")
				return
			default:
				fmt.Println("goroutine monitoring...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("OK,use the channel to stop the work.")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
