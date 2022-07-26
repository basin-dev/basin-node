package main

import (
	"context"

	adapters "github.com/sestinj/basin-node/adapters"
	cmd "github.com/sestinj/basin-node/cmd"
	// kademlia "github.com/libp2p/go-libp2p-kad-dht"
)

func main() {
	ctx := context.Background()

	// Start up the local LevelDB database
	adapters.StartDB()

	// Start up this node's HTTP API, concurrently with CLI
	go RunHttpServer()

	// Run the CLI
	cmd.Execute()
}
