package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	//创建文件，打开文件
	file, _ := os.Create("osexample.txt")
	file.WriteString("os.example.test")
	defer file.Close()
	fmt.Println(reflect.ValueOf(file).Type())

	//获取环境变量
	fmt.Println(os.Environ())
	fmt.Println(os.Getwd())
	/***
	hello /home/work/software/go，
	$GOROOT替换为环境变量的值，而h没有环境变量可以替换，返回空字符串
	 */
	fmt.Println(os.ExpandEnv("hello $GOROOT"))
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(os.Getpid())
	fmt.Println(os.FindProcess(os.Getpid()))
}
