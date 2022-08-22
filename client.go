package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	log.Fatal("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
