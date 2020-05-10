// Package main implements a server for Greeter service.
package main

// https://segmentfault.com/a/1190000007933303
import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	pb "github.com/bjzhang03/exlocus-godemo/protobuf/proto"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	tlsport = "127.0.0.1:50052"
)

// server is used to implement helloworld.GreeterServer.
type serverTls struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *serverTls) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Return  " + in.GetName()}, nil
}

func main() {
	fmt.Println("start")
	lis, err := net.Listen("tcp", tlsport)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("./tls/server.pem", "./tls/server.key")
	if err != nil {
		fmt.Printf("Failed to generate credentials %v", err)
	}

	fmt.Println("tlsf finish")
	// 实例化grpc Server, 并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))
	// 注册HelloService
	pb.RegisterGreeterServer(s, &serverTls{})
	log.Println("Listen on " + tlsport + " with TLS")
	fmt.Println("Listen on " + tlsport + " with TLS")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
