package main

import (
	"fmt"
	pb "github.com/bjzhang03/go-awesome/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const ipPort = "127.0.0.1:50051"

type streamServer struct {
}

//服务端 单向流
func (ss *streamServer) GetStream(req *pb.MsgRequest, res pb.MsgGreeter_GetStreamServer) error {
	i := 0
	for {
		i++
		res.Send(&pb.MsgReply{Data: fmt.Sprintf("Server time := %v, send by getstream!", time.Now().Unix())})
		time.Sleep(1 * time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

//客户端 单向流
func (ss *streamServer) PutStream(cliStr pb.MsgGreeter_PutStreamServer) error {
	for {
		if tem, err := cliStr.Recv(); err == nil {
			log.Printf("Server time := %v, receive := %s by putstream", time.Now().Unix(), tem.String())
		} else {
			log.Printf("break, err := %s", err.Error())
			break
		}
	}
	return nil
}

//客户端服务端 双向流
func (ss *streamServer) AllStream(allStr pb.MsgGreeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 通过 all stream 接收数据
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Printf("Server time := %v, receive data := %s by all stream!", time.Now().Unix(), data)
		}
		wg.Done()
	}()

	// 通过all stream 发送数据
	go func() {
		for {
			allStr.Send(&pb.MsgReply{Data: fmt.Sprintf("Server time := %v, send msg by all stream", time.Now().Unix())})
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
	return nil
}

func main() {
	//监听端口
	lis, err := net.Listen("tcp", ipPort)
	if err != nil {
		log.Printf("Listen port := %s failed! error := %s", ipPort, err.Error())
		os.Exit(1)
	}
	//创建一个grpc 服务器
	s := grpc.NewServer()
	//注册事件
	pb.RegisterMsgGreeterServer(s, &streamServer{})
	//注意这里这个pro是你定义proto文件生成后的那个go文件中引用的，而后面的这个函数是注册名称，是根据你自己定义的名称生成的，不同的文件该函数名称是不一样的，不过都是register这个单词开头的，这里不能原样照搬
	//处理链接
	s.Serve(lis)
}
