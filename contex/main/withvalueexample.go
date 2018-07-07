package main

import (
	"time"
	"fmt"
	"context"
)

var key string = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	valueCtx := context.WithValue(ctx, key, "[monitor 1]")
	go watch2(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("OK,Send the signal to kill the monitor!")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func watch2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "Esc the monitor,Stoping...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine monitoring...")
			time.Sleep(2 * time.Second)
		}
	}
}
