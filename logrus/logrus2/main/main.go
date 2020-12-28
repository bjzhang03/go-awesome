package main

import (
	"github.com/bjzhang03/go-awesome/logrus/logrus2/logt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	// 定义一个文件
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Get pwd path failed! error := %s", err.Error())
	}
	logpath := pwd + "/configs/logrus"
	log.Printf("log path := %s", logpath)
	_, err = os.Stat(logpath)
	// 如果目录不存在则创建一个
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(logpath, os.ModePerm)
		}
	}
	// 生成文件数据
	fileName := logpath + "/golog.log"

	//创建一个hook，将日志存储路径输入进去
	hook := logt.NewHook(fileName)
	//加载hook之前打印日志
	logrus.WithField("file", "golog.log").Info("New logrus hook err.")
	logrus.AddHook(hook)
	//加载hook之后打印日志
	logrus.SetReportCaller(true)
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
