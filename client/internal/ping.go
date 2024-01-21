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

var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the server",
	RunE:  runPing,
}

var ip string
var port int

func init() {
	PingCmd.Flags().StringVarP(&ip, "ip", "i", "localhost", "IP address of the gRPC server")
	PingCmd.Flags().IntVarP(&port, "port", "p", 50051, "Port of the gRPC server")
	rootCmd.AddCommand(PingCmd)
}

func runPing(cmd *cobra.Command, args []string) error {
	address := fmt.Sprintf("%s:%d", ip, port)

	_, err := fmt.Println("Connecting to server at: ", address)
	if err != nil {
		log.Fatalf("could not print: %v", err)
	}

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
