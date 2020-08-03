package main

import (
	"log"

	pb "github.com/bjzhang03/exlocus-godemo/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	tlsAuthClientAddress = "127.0.0.1:50052"
	OpenTls              = true
)

type customerCredential struct {
}

func (c customerCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "I am key",
	}, nil
}

func (c customerCredential) RequireTransportSecurity() bool {
	if OpenTls {
		return true
	}
	return false
}

func main() {

	var err error
	var opts []grpc.DialOption

	if OpenTls {
		/*tls的链接*/
		creds, err := credentials.NewClientTLSFromFile("./tls/server.pem", "localhost")
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	/*指定自定义认证*/
	opts = append(opts, grpc.WithPerRPCCredentials(new(customerCredential)))

	conn, err := grpc.Dial(tlsAuthClientAddress, opts...)

	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	/*初始化客户端*/
	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "tlsAuthTest"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
