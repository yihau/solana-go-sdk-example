package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/types"
)

func main() {
	c := client.NewClient("http://localhost:8899")

	newAccount := types.NewAccount()
	fmt.Println(newAccount.PublicKey.ToBase58())
	fmt.Println(newAccount.PrivateKey)

	txhash, err := c.RequestAirdrop(
		context.Background(),
		newAccount.PublicKey.ToBase58(),
		1e9,
	)
	if err != nil {
		log.Fatalf("request airdrop error, err: %v", err)
	}

	fmt.Println("txhash:", txhash)
}
