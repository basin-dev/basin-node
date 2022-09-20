package main

import (
	"fmt"
	"log"

	"github.com/onrik/ethrpc"
)

var (
	EthClient *ethrpc.EthRPC
)

func StartEthClient(gatewayUrl string) (*ethrpc.EthRPC, error) {
	client := ethrpc.New(gatewayUrl)

	_, err := client.Web3ClientVersion()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Ethereum node: %w\n", err)
	}

	EthClient = client
	return client, nil
}

func SendTx() {
	txid, err := EthClient.EthSendTransaction(ethrpc.T{
		From:  "0x6247cf0412c6462da2a51d05139e2a3c6c630f0a",
		To:    "0xcfa202c4268749fbb5136f2b68f7402984ed444b",
		Value: ethrpc.Eth1(),
	})
	if err != nil {
		log.Fatal(err)
	}
}
