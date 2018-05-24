package main

import (
	"log"
	"os"
	"net"
	pb "bjzhang.com/godemo/grpc/proto"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

func main() {

	log.SetOutput(os.Stdout)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	log.Println("Server......")
	s.Serve(lis)

}

const (
	port = ":50051"
)

type server struct{}

func (s *server) Login(ctx context.Context, usr *pb.UserInfo) (*pb.FuncResponse, error) {
	log.Println("Server...... Login() UserInfo:", usr)

	usr.Id = 200
	strId := "200" //strconv.Itoa(usr.Id)

	return &pb.FuncResponse{Reply: strId}, nil

}

func (s *server) Logout(ctx context.Context, uid *pb.UserID) (*pb.FuncResponse, error) {
	log.Println("Server...... Logout() UserID:", uid)
	return &pb.FuncResponse{Reply: "Logout Successed."}, nil
}
