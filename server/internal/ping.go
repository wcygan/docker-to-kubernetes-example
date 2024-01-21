package internal

import (
	"context"
	"log"

	pb "github.com/wcygan/docker-to-kubernetes-example/generated/go/ping/v1"
)

type PingService struct {
	pb.UnimplementedPingServiceServer
}

func (s *PingService) Ping(_ context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.PingResponse{Message: "Pong"}, nil
}
