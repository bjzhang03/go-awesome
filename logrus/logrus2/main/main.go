package main

import (
	"github.com/bjzhang03/exlocus-godemo/logrus/logrus2/logt"
	"github.com/sirupsen/logrus"
)

func main() {
	//创建一个hook，将日志存储路径输入进去
	hook := logt.NewHook("golog.log")
	//加载hook之前打印日志
	logrus.WithField("file", "golog.log").Info("New logrus hook err.")
	logrus.AddHook(hook)
	//加载hook之后打印日志
	logrus.SetReportCaller(true)
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
