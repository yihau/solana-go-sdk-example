package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
)

func main() {
	c := client.NewClient("http://localhost:8899")
	resp, err := c.GetVersion(context.Background())
	if err != nil {
		log.Fatalf("get version error, err: %v", err)
	}
	fmt.Println("solana version:", resp.SolanaCore)
}
