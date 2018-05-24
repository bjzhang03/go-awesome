package main

import (
	"fmt"
	"time"
)

//select与非缓冲通道
//与操作缓冲通道的select相比，它被阻塞的概率一般会大很多。只有存在可配对的操作的时候，传递元素值的动作才能真正的开始。

/***
*发送操作间隔1s,接收操作间隔2s
*分别向unbufChan通道发送小于10和大于等于10的整数，这样更容易从打印结果分辨出配对的时候哪一个case被选中了。下列案例两个case是被随机选择的。
 */
func main() {
	unbufChan := make(chan int)
	sign := make(chan byte, 2)

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case unbufChan <- i:
			case unbufChan <- i + 10:
			default:
				fmt.Println("default!")
			}
			time.Sleep(time.Second)
		}
		close(unbufChan)
		fmt.Println("The channel is closed.")
		sign <- 0
	}()
	go func() {
	loop:
		for {
			select {
			case e, ok := <-unbufChan:
				if !ok {
					fmt.Println("Closed channel.")
					break loop
				}
				fmt.Printf("e: %d\n", e)
				time.Sleep(2 * time.Second)
			}
		}
		sign <- 1
	}()
	<-sign
	<-sign
}
