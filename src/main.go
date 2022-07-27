package main

import (
	"context"
	"log"

	adapters "github.com/sestinj/basin-node/adapters"
	cmd "github.com/sestinj/basin-node/cmd"
	"github.com/sestinj/basin-node/util"
)

func main() {
	ctx := context.Background()

	// Start the BasinNode (libp2p host with associated protocol, stream handler)
	basin, err := StartBasinNode()
	if err != nil {
		log.Fatal("Failed to instantiate the BasinNode: " + err.Error())
	}

	// Create the Router
	_, err = StartKademliaRouter(ctx, basin.Host)
	if err != nil {
		log.Fatal("Failed to instantiate router: ", err.Error())
	}

	// Create new PubSub
	_, err = StartPubSub(ctx, basin.Host)
	if err != nil {
		log.Fatal("Failed to instantiate pubsub: " + err.Error())
	}

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
