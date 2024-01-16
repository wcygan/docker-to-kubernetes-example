package main

import (
	"context"
	"log"
	"time"

	pb "github.com/wcygan/docker-to-kubernetes-example/generated/go/ping/v1"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Ping(ctx, &pb.PingRequest{Message: "Ping"})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}
