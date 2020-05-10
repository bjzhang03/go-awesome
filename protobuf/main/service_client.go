package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/bjzhang03/exlocus-godemo/protobuf/proto"
	"google.golang.org/grpc"
)

// Package main implements a client for Greeter service.

const (
	address     = "localhost:50051"
	defaultName = "hello world!"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 使用前面的链接创建一个客户端
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	// 如果没有输入的参数则使用默认的参数数据
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
