package main

import (
	"time"
	"fmt"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "[monitor 1]: ")
	go watch(ctx, "[monitor 2]: ")
	go watch(ctx, "[monitor 3]: ")
	time.Sleep(10 * time.Second)
	fmt.Println("OK,use the context to stop the work.")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Esc the monitoring,Stoping...")
			return
		default:
			fmt.Println(name, "goroutine monitoring...")
			time.Sleep(2 * time.Second)
		}
	}
}
