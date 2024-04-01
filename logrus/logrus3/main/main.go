package main

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})
	log.AddHook(lfHook)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
}

//切割日志和清理过期日志
func ConfigLocalFilesystemLogger1(filePath string) {
	writer, err := rotatelogs.New(
		filePath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(filePath),           // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Second*60*3),     // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Second*60), // 日志切割时间间隔
	)
	if err != nil {
		log.Fatal("Init log failed, err:", err)
	}
	log.SetOutput(writer)
	log.SetLevel(log.InfoLevel)
}

func main() {
	// 定义一个文件
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Get pwd path failed! error := %s", err.Error())
	}
	logpath := pwd + "/configs/tmp/logrus"
	log.Printf("log path := %s", logpath)
	_, err = os.Stat(logpath)
	// 如果目录不存在则创建一个
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(logpath, os.ModePerm)
		}
	}
	//ConfigLocalFilesystemLogger1("/home/bjzhang03/gocode/src/github.com/bjzhang03/exlocus-godemo/logrus/logrus3/main/log")
	ConfigLocalFilesystemLogger(logpath+"/data/",
		"log", time.Second*60*3, time.Second*60)
	for {
		log.Info(111)
		time.Sleep(500 * time.Millisecond)
	}
}
