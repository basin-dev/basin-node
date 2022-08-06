package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func StartEthClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		return nil, err
	}
	return client, nil
}
