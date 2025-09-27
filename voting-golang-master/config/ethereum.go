package config

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient *ethclient.Client

func init() {
	ethClient, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545")

	if err != nil {
		panic(err)
	}
	EthClient = ethClient
	log.Println("Connected to Ethereum!")

}
