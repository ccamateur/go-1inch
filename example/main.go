package main

import (
	"context"
	"fmt"
	"log"

	go1inch "github.com/jon4hz/go-1inch"
)

func main() {
	client := go1inch.NewClient()
	res, _, err := client.ApproveCalldata(context.Background(), "eth", "0x6b175474e89094c44da98b954eedeac495271d0f", &go1inch.ApproveCalldataOpts{
		Amount: "5000000",
	})
	if err != nil {
		log.Fatal("Error while getting an approve calldata ", err)
	}
	fmt.Println(res)

	fmt.Println(client.Tokens(context.Background(), "matic"))
}
