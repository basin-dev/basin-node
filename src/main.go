package main

import (
	"context"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	adapters "github.com/sestinj/basin-node/adapters"
	cmd "github.com/sestinj/basin-node/cmd"
	"github.com/sestinj/basin-node/util"
	// kademlia "github.com/libp2p/go-libp2p-kad-dht"
)

func main() {
	ctx := context.Background()

	// Create listener on port
	host, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		println("host err")
		panic(err)
	}

	// Create the Router
	_, err = StartKademliaRouter(ctx, host)
	if err != nil {
		log.Fatal("Failed to instantiate router: ", err.Error())
	}

	// Create new PubSub
	_, err = StartPubSub(ctx, host)
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
