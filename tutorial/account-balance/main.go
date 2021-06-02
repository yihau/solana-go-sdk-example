package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
)

func main() {
	c := client.NewClient("http://localhost:8899")
	balance, err := c.GetBalance(context.Background(), "6ygSF4M5AXDgyzwBneQyM3wnadUcZ9tX9Y3Y58fRmXNc")
	if err != nil {
		log.Fatalln("get balance error", err)
	}
	fmt.Println(balance)
}
