package main

import (
	"log"
	"net"

	"grpc-service-poc/internal/generated/model"
	"grpc-service-poc/internal/service"

	"google.golang.org/grpc"
)

func main() {
	// Start listening on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server with custom options
	grpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(1024),
	)

	// Register the EntityService server implementation with the gRPC server
	model.RegisterEntityServiceServer(grpcServer, service.NewEntityServiceServer())

	// Log the server address and start serving gRPC requests
	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
