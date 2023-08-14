package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/b70264eed22742869f075f750e213bea")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0x68ae23370f5ac690ad47eea221324bd015fffdee9c327774b29635c214fce47a
			fmt.Println(block.Number().Uint64())   // 17832647
			fmt.Println(block.Time())              // 1691043647
			fmt.Println(block.Nonce())             // 0
			fmt.Println(len(block.Transactions())) // 115
		}
	}
}
