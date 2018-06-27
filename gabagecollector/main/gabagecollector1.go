package main

import (
	"time"
	"fmt"
	"runtime/debug"
	"runtime"
)

const (
	windowSize = 200000
	msgCount   = 1000000
)

type (
	message []byte
	buffer [windowSize]message
)

var worst time.Duration

func mkMessage(n int) message {
	m := make(message, 1024)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

func pushMsg(b *buffer, highID int) {
	start := time.Now()
	m := mkMessage(highID)
	(*b)[highID%windowSize] = m
	elapsed := time.Since(start)
	if elapsed > worst {
		worst = elapsed
	}
}

func main() {
	var b buffer
	// 主动调用垃圾回收机制的两种方案
	runtime.GC()
	for i := 0; i < msgCount; i++ {
		pushMsg(&b, i)
	}
	// 主动调用垃圾回收机制的两种方案
	debug.FreeOSMemory()
	fmt.Println("Worst push time: ", worst)
}
