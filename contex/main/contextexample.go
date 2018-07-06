package main

import (
	"fmt"
	"time"
	"context"
)

// 与goroutine中的channelexample配合起来研究
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Esc the monitoring,Stoping...")
				return
			default:
				fmt.Println("goroutine monitoring...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("OK,use the context to stop the work.")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
