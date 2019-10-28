package main

import (
	"log"
	"os"
	"google.golang.org/grpc"
	"strconv"
	pb "github.com/bjzhang03/exlocus-godemo/grpc/proto"
	"golang.org/x/net/context"
)

const (
	address = "localhost:50051"
)

func main() {
	log.SetOutput(os.Stdout)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	var status pb.UserStatus
	status = pb.UserStatus_ONLINE

	userInfo := &pb.UserInfo{
		Id:     30,         //proto.Int32(10),
		Name:   "grpcTest111", //proto.String("XCL-gRPC"),
		Status: status,
	}

	r, err := c.Login(context.Background(), userInfo)
	if err != nil {
		log.Fatalf("Login Failed!  %v", err)
	}
	log.Println("Login():", r)

	uid, err := strconv.ParseInt(r.Reply, 10, 32)
	if err != nil {
		log.Fatalf("Error:  %v", err)
	}

	userID := &pb.UserID{int32(uid)}
	out, err := c.Logout(context.Background(), userID)
	log.Println("Logout():", out)
}
