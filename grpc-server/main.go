package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/alimikegami/go-monorepo/grpc-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
}

func (s *UserService) GetUsers(ctx context.Context, _ *emptypb.Empty) (*proto.Users, error) {
	user1 := &proto.User{
		Id:   1,
		Name: "Alim Ikegami",
	}

	user2 := &proto.User{
		Id:   2,
		Name: "Ikegami",
	}

	return &proto.Users{
		Users: []*proto.User{
			user1,
			user2,
		},
	}, nil
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register our service implementation with the gRPC server
	proto.RegisterUserServiceServer(server, &UserService{})

	fmt.Println("Server started on port 50051")

	// Start serving requests
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
