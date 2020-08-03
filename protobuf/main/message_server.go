package main

import (
	stProto "github.com/bjzhang03/exlocus-godemo/protobuf/proto"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "localhost:6600")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		log.Printf("new connect to %s", conn.RemoteAddr().String())
		go readMessage(conn)
	}
}

func readMessage(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096, 4096)
	for {
		//读消息
		cnt, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}

		stReceive := &stProto.UserInfo{}
		pData := buf[:cnt]

		//protobuf解码
		err = proto.Unmarshal(pData, stReceive)
		if err != nil {
			panic(err)
		}

		log.Printf("receive message success! ipaddress := %s, info := %s", conn.RemoteAddr(), stReceive.String())
		for _, val := range stReceive.Message {
			if val == "stop" {
				os.Exit(0)
			}
		}
	}
}
