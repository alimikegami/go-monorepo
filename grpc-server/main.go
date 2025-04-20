package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/alimikegami/go-monorepo/grpc-server/config"
	"github.com/alimikegami/go-monorepo/grpc-server/db"
	"github.com/alimikegami/go-monorepo/grpc-server/entity"
	"github.com/alimikegami/go-monorepo/grpc-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
	Db *gorm.DB
}

func (s *UserService) GetUsers(ctx context.Context, _ *emptypb.Empty) (*proto.Users, error) {
	var data []entity.User
	err := s.Db.Find(&data).Error
	if err != nil {
		return nil, err
	}

	var users []*proto.User

	for _, user := range data {
		users = append(users, &proto.User{
			Id:   int64(user.ID),
			Name: user.Name,
		})
	}

	return &proto.Users{
		Users: users,
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *proto.User) (*proto.User, error) {
	user := &entity.User{
		Name: req.Name,
	}
	res := s.Db.WithContext(ctx).Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &proto.User{
		Id:   int64(user.ID),
		Name: user.Name,
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

	config, _ := config.LoadConfig()

	db, _ := db.InitDB(config)

	// Register our service implementation with the gRPC server
	proto.RegisterUserServiceServer(server, &UserService{Db: db})

	fmt.Println("Server started on port 50051")

	// Start serving requests
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
