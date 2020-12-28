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
	logpath := pwd + "/configs/os"
	log.Printf("log path := %s", logpath)
	_, err = os.Stat(logpath)
	// 如果目录不存在则创建一个
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(logpath, os.ModePerm)
		}
	}
	f, err := os.OpenFile(logpath+"/notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
