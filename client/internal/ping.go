package internal

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	pb "github.com/wcygan/docker-to-kubernetes-example/generated/go/ping/v1"
	"google.golang.org/grpc"
	"log"
	"time"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the server",
	RunE:  runPing,
}

func runPing(cmd *cobra.Command, args []string) error {
	address := fmt.Sprintf("%s:%d", ip, port)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
	return nil
}
