package main

import (
	"time"
	"os"
	"os/signal"
	"syscall"
	"runtime"
	"fmt"
)

func main() {
	setupSigusr1Trap()
	go a1()
	m11()
}
func m11() {
	m22()
}
func m22() {
	m33()
}
func m33() {
	time.Sleep(time.Second)
}

func a1() {
	time.Sleep(time.Second)
}
func setupSigusr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			DumpStacks()
		}
	}()
}
func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}