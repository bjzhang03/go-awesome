package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	// 定义一个文件
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Get pwd path failed! error := %s", err.Error())
	}
	logpath := pwd + "/configs/tmp/os"
	log.Printf("log path := %s", logpath)
	_, err = os.Stat(logpath)
	// 如果目录不存在则创建一个
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(logpath, os.ModePerm)
		}
	}
	// 生成文件数据
	fileName := logpath + "/osexample.txt"
	//创建文件，打开文件
	file, _ := os.Create(fileName)
	file.WriteString("os.example.test\n")
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
