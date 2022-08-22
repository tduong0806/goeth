package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	rawTx := "f86e018504a817c80082520894ad7e5e885be0d4ba122e40e219c93f245c32ec4c880de0b6b3a764000080822d46a0dac14614c868de59ce2c7a8830473b8a78ef8b1250a9557a05057c5f5fec09b7a07354c21c3839da39379edbc7797ae66888d7547a9a17fad9416043f27b658d0d"

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
