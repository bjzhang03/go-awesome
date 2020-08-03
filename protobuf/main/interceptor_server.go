package main

import (
	"log"
	"net"

	pb "github.com/bjzhang03/exlocus-godemo/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

const (
	tlsAuthServerAddress = "127.0.0.1:50052"
)

type helloService struct {
}

var HelloService = helloService{}

// SayHello implements helloworld.GreeterServer
func (s *helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Return  " + in.GetName()}, nil
}

func auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "无Token认证消息!")
	}
	var (
		appid  string
		appkey string
	)
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}
	if appid != "101010" || appkey != "I am key" {
		return grpc.Errorf(codes.Unauthenticated, "Token认证消息无效: appid = %s, appkey = %s", appid, appkey)
	}
	return nil
}

func main() {
	log.Println("start")
	lis, err := net.Listen("tcp", tlsAuthServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("./tls/server.pem", "./tls/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}
	opts = append(opts, grpc.Creds(creds))

	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = auth(ctx)
		if err != nil {
			return
		}
		return handler(ctx, req)
	}

	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	log.Println("tls auth finish")
	// 实例化grpc Server, 并开启TLS认证
	s := grpc.NewServer(opts...)
	// 注册HelloService
	pb.RegisterGreeterServer(s, &HelloService)
	log.Println("Listen on " + tlsAuthServerAddress + " with TLS Auth")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
