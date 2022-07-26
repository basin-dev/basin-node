package main

import (
	"context"
	"log"

	adapters "github.com/sestinj/basin-node/adapters"
	cmd "github.com/sestinj/basin-node/cmd"
	"github.com/sestinj/basin-node/util"
	// kademlia "github.com/libp2p/go-libp2p-kad-dht"
)

func main() {
	ctx := context.Background()

	// Start up the local LevelDB database
	db, err := adapters.StartDB()
	if err != nil {
		log.Fatal(err)
	}

	util.StartLocalOnlyDb(db, "/local/")

	// Start up this node's HTTP API, concurrently with CLI
	go RunHttpServer(ctx)

	// Run the CLI
	cmd.Execute()
}
