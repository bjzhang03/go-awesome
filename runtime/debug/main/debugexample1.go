package main

import (
	"fmt"
	"runtime/debug"
)
// 打印堆栈信息的例子
func test1() {
	test2()
}

func test2() {
	test3()
}

func test3() {
	fmt.Printf("%s", debug.Stack())
	debug.PrintStack()
}

func main() {
	test1()
}