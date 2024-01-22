package internal

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	pb "github.com/wcygan/docker-to-kubernetes-example/generated/go/kv/v1"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Put a key-value pair",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		req := &pb.PutRequest{Key: key, Value: value}
		res, err := client.Put(context.Background(), req)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Success: %v\n", res.Success)
	},
}
