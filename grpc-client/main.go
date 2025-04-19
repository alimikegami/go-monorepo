package main

import (
	"context"
	"log"
	"time"

	"github.com/alimikegami/go-monorepo/grpc-client/github.com/alimikegami/go-monorepo/grpc-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := proto.NewUserServiceClient(conn)

	// Set a deadline for the RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the SayHello method
	response, err := client.GetUsers(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	log.Printf("Response from server: %+v", response.Users)
}
