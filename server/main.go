// server/main.go
package main

import (
	"log"
	"net"

	pb "github.com/wcygan/docker-to-kubernetes-example/generated/go/ping/v1"
	"github.com/wcygan/docker-to-kubernetes-example/server/internal"
	"google.golang.org/grpc"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server and register services
	s := grpc.NewServer()
	registerServer(s)

	// Start listening for requests
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerServer(s *grpc.Server) {
	pb.RegisterPingServiceServer(s, &internal.PingService{})
}
