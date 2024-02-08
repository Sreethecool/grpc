package main

import (
	"grpc/internal/db"
	"grpc/internal/proto"
	"grpc/internal/blog"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Initialize database repository
	postRepo := db.NewPostRepository()

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register API server with gRPC server
	proto.RegisterBlogServiceServer(grpcServer, blog.NewGRPCServer(postRepo))

	// Start gRPC server
	port := "50051"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("gRPC server listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}