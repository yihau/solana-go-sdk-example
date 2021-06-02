package main

import (
	"fmt"

	"github.com/portto/solana-go-sdk/types"
)

func main() {
	newAccount := types.NewAccount()
	fmt.Println(newAccount.PublicKey.ToBase58())
	fmt.Println(newAccount.PrivateKey)

	newAccount2 := types.AccountFromPrivateKeyBytes(newAccount.PrivateKey)
	fmt.Println(newAccount2.PublicKey.ToBase58())
}
