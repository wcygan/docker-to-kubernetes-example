// server/main.go
package main

import (
	"github.com/wcygan/docker-to-kubernetes-example/generated/go/kv/v1"
	"github.com/wcygan/docker-to-kubernetes-example/server/internal/kv"
	"github.com/wcygan/docker-to-kubernetes-example/server/internal/ping"
	"log"
	"net"

	"github.com/wcygan/docker-to-kubernetes-example/generated/go/ping/v1"
	"google.golang.org/grpc"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Register services with the server
	err = registerServer(s)
	if err != nil {
		log.Fatalf("failed to register server: %v", err)
	}

	// Start listening for requests
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerServer(s *grpc.Server) error {
	pingv1.RegisterPingServiceServer(s, &ping.PingService{})
	kvv1.RegisterKeyValueServiceServer(s, kv.NewKeyValueService(kv.NewRedisClient()))
	return nil
}
