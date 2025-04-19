package main

import (
	"context"
	"log"
	"time"

	// Correct import path corresponding to the 'go_package' option
	proto "github.com/alimikegami/go-monorepo/grpc-server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb" // Needed for GetUsers request
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := proto.NewUserServiceClient(conn) // Use the 'proto' alias

	// Set a deadline for the RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the GetUsers method using an empty request message
	request := &emptypb.Empty{}
	response, err := client.GetUsers(ctx, request)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	log.Printf("Response from server: %+v", response.Users)
}
