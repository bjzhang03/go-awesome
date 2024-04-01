package main

import (
	"log"
	"os"
)

func main() {
	// 定义一个文件
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Get pwd path failed! error := %s", err.Error())
	}
	logpath := pwd + "/configs/logs"
	log.Printf("log path := %s", logpath)
	_, err = os.Stat(logpath)
	// 如果目录不存在则创建一个
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(logpath, os.ModePerm)
		}
	}

	// 生成文件数据
	fileName := logpath + "/ll.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalf("open file error! error := %s", err.Error())
	}
	// 创建一个日志对象
	debugLog := log.New(logFile, "[Debug]", log.LstdFlags)
	debugLog.Println("A debug message here")
	//配置一个日志格式的前缀
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A Info Message here ")
	//配置log的Flag参数
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A different prefix")
}
