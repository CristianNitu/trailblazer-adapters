package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

func processTransactionIndexer(client *ethclient.Client, processor adapters.BlockProcessor, blockNumber int64) error {
	blockNumberBig := big.NewInt(blockNumber)
	block, err := client.BlockByNumber(context.Background(), blockNumberBig)
	if err != nil {
		log.Fatalf("Failed to fetch the block: %v", err)
		return err
	}

	senders, err := processor.ProcessBlock(context.Background(), block, client)
	if err != nil {
		log.Fatalf("Failed to process the block: %v", err)
		return err
	}

	fmt.Printf("Senders: %v\n", senders)
	return nil
}
