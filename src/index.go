package main

import (
	"context"
	"log"
	"net/http"

	"github.com/sestinj/basin-node/adapters"
	. "github.com/sestinj/basin-node/node"
	server "github.com/sestinj/basin-node/server/go"
	"github.com/sestinj/basin-node/util"
)

func RunHttpServer(ctx context.Context, b *BasinNode, addr string) {
	DefaultApiService := server.NewDefaultApiService()
	DefaultApiController := server.NewDefaultApiController(DefaultApiService)

	router := server.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8555", router))
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
