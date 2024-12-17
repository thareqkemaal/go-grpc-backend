package main

import (
	"log"
	"net"

	pb "github.com/thareqkemaal/go-grpc-backend/service"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

		grpcServer := grpc.NewServer()
		pb.RegisterUserServiceServer(grpcServer, &server{})

		log.Println("gRPC server is running on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}