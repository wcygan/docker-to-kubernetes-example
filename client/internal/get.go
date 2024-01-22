package internal

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	pb "github.com/wcygan/docker-to-kubernetes-example/generated/go/kv/v1"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		req := &pb.GetRequest{Key: key}
		res, err := client.Get(context.Background(), req)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if res.Found {
			fmt.Printf("Value: %v\n", res.Value)
		} else {
			fmt.Println("Key not found")
		}
	},
}
