package main

import (
	"context"
	"fmt"
	pb "github.com/bjzhang03/go-awesome/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	ipAddress = "127.0.0.1:50051"
)

func main() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(ipAddress, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pb.NewMsgGreeterClient(conn) //这个pro与服务端的同理，也是来源于proto编译生成的那个go文件内部调用
	//调用服务端推送流
	reqstreamData := &pb.MsgRequest{Data: fmt.Sprintf("Client time := %v, getstream", time.Now().Unix())}
	res, _ := c.GetStream(context.Background(), reqstreamData)
	for {
		aa, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Client receive data := %s", aa.Data)
	}
	//客户端 推送 流
	putRes, _ := c.PutStream(context.Background())
	i := 1
	for {
		i++
		putRes.Send(&pb.MsgRequest{Data: fmt.Sprintf("Client time := %v, putstream", time.Now().Unix())})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	//服务端 客户端 双向流
	allStr, _ := c.AllStream(context.Background())
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Printf("Client receive data := %s", data.String())
		}
	}()

	go func() {
		for {
			allStr.Send(&pb.MsgRequest{Data: fmt.Sprintf("Client time := %v, allstream", time.Now().Unix())})
			time.Sleep(time.Second)
		}
	}()
	select {}
}
