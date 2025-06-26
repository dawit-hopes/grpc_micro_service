// Package main serves as the entry point
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dawit_hopes/grpc_micro_service/client/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/dawit_hopes/grpc_micro_service/proto"
)

func main() {
	env := config.NewEnv()

	// connect to grpc server
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", env.Serverhost, env.ServerAddress), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connet: %v", err)
	}

	defer conn.Close()

	// Create the grpc client
	client := pb.NewUserServiceClient(conn)
	ctx, canel := context.WithTimeout(context.Background(), 5*time.Second)
	defer canel()

	// call crate method

	createResp, cErr := client.Create(ctx, &pb.CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	})

	if cErr != nil {
		log.Fatalf("create error: %v", err)
	}

	newData := &pb.UserProfileResponse{
		Id: createResp.GetId(),
		Name: createResp.GetName(),
		Email: createResp.GetEmail(),
	}

	log.Printf("create response: %v", newData)
}
