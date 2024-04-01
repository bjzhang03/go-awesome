package main

import (
	"context"
	"log"

	pb "github.com/bjzhang03/go-awesome/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Package main implements a client for Greeter service.

const (
	tlsaddress     = "localhost:50052"
	tlsdefaultName = "hello world!"
)

func main() {
	// Set up a connection to the server.
	// TLS连接
	log.Println("client")
	creds, err := credentials.NewClientTLSFromFile("./tls/server.pem", "localhost")
	log.Println("cred")
	if err != nil {
		log.Printf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(tlsaddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	log.Println("conn")
	// 使用前面的链接创建一个客户端
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := tlsdefaultName
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
