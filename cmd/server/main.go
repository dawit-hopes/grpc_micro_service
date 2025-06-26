// Package main implements the main entry point for the server application.
package main

import (
	"log"
	"net"

	"github.com/dawit_hopes/grpc_micro_service/internal/config"
	interfaces "github.com/dawit_hopes/grpc_micro_service/pkg/v1"
	"github.com/dawit_hopes/grpc_micro_service/pkg/v1/handler"
	"github.com/dawit_hopes/grpc_micro_service/pkg/v1/repository"
	"github.com/dawit_hopes/grpc_micro_service/pkg/v1/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func main() {
	env := config.NewEnv()
	client, err := config.NewMongoClient(env)

	if err != nil {
		log.Fatalf("Error connecting to mongo client %v", err.Error())
	}

	// add a listent address
	lis, lErr := net.Listen("tcp", ":8080")
	if lErr != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", lErr)
	}

	// start the grpc server
	grpcServer := grpc.NewServer()

	// get the user usecase
	userUsecase := initUserServer(client)

	// set up the handler
	handler.NewServer(grpcServer, userUsecase)

	// start the server
	log.Fatal(grpcServer.Serve(lis))
}

func initUserServer(db *mongo.Client) interfaces.UseCaseInterface {
	mongoCollection := db.Database("grpc_clen").Collection("user")
	userRepo := repository.NewUserRepository(mongoCollection)

	return usecase.NewUseCase(userRepo)
}
