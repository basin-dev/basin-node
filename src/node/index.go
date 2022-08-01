package node

import (
	"context"
	"log"

	"github.com/sestinj/basin-node/adapters"
	"github.com/sestinj/basin-node/util"
)

type BasinNodeConfig struct {
	Http string
	Did  string
	Pw   string
}

func (c *BasinNodeConfig) SetDefaults() {
	if c.Http == "" {
		c.Http = "127.0.0.1:3000"
	}
}

func StartEverything(ctx context.Context, config BasinNodeConfig) {
	// Start the BasinNode (libp2p host with associated protocol, stream handler)
	basin, err := StartBasinNode(config)
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
	RunHttpServer(ctx, &basin, config.Http)
}
