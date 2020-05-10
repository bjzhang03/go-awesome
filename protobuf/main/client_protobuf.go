package main

import (
	"bufio"
	stProto "github.com/bjzhang03/exlocus-godemo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	strIP := "localhost:6600"
	var conn net.Conn
	var err error

	//连接服务器
	for conn, err = net.Dial("tcp", strIP); err != nil; conn, err = net.Dial("tcp", strIP) {
		log.Printf("connect to %s failed", strIP)
		time.Sleep(time.Second)
		log.Printf("reconnect...")
	}
	log.Printf("connect to %s success!", strIP)
	defer conn.Close()

	//发送消息
	cnt := 0
	// 从输入获取数据
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		cnt++
		// 根据输入构造结构体对象,然后直接发送出去
		msgs := strings.Split(sender.Text(), ",")
		stSend := &stProto.UserInfo{
			Message: msgs,
			Length:  *proto.Int(len(msgs)),
			Cnt:     *proto.Int(cnt),
			Other:   strIP,
		}

		//protobuf编码
		pData, err := proto.Marshal(stSend)
		if err != nil {
			panic(err)
		}

		//发送
		conn.Write(pData)
		for _, val := range msgs {
			if val == "stop" {
				return
			}
		}

	}
}
