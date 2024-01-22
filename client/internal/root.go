package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pb "github.com/wcygan/docker-to-kubernetes-example/generated/go/kv/v1"
)

var ip string
var port int
var conn *grpc.ClientConn
var client pb.KeyValueServiceClient

func init() {
	rootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "localhost", "IP address of the gRPC server")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Port of the gRPC server")
	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(putCmd)
	rootCmd.AddCommand(getCmd)
}

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "A client for docker-to-kubernetes-example",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		conn, err = grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		client = pb.NewKeyValueServiceClient(conn)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		conn.Close()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
