package main

import (
	"context"
	"log"
	"net"

	pb "github.com/wcygan/docker-to-kubernetes-example/generated/go/ping/v1"
	"google.golang.org/grpc"
)

// server is used to implement ping.PingService.
type server struct {
	pb.UnimplementedPingServiceServer
}

// Ping implements ping.PingServiceServer
func (s *server) Ping(_ context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.PingResponse{Message: "Pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
