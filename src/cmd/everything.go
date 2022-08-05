package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sestinj/basin-node/adapters"
	. "github.com/sestinj/basin-node/node"
	server "github.com/sestinj/basin-node/server/go"
	"github.com/sestinj/basin-node/util"
)

func RunHttpServer(ctx context.Context, b *BasinNode, addr string) {
	DefaultApiService := server.NewDefaultApiService()
	DefaultApiController := server.NewDefaultApiController(DefaultApiService)

	router := server.NewRouter(DefaultApiController)

	segs := strings.Split(b.Http, ":")
	port := segs[len(segs)-1]
	fmt.Fprintf(os.Stdout, "Listening at %s...\n", b.Http)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func StartEverything(ctx context.Context, config BasinNodeConfig) {
	// NOTE: The order here matters. For example, db must start before node.

	// Start up the local LevelDB database
	log.Println("Initializing LevelDB...")
	db, err := adapters.StartDB()
	if err != nil {
		log.Fatal(err)
	}

	util.StartLocalOnlyDb(db, "/local/")

	// Start the BasinNode (libp2p host with associated protocol, stream handler)
	log.Println("Launching Basin Node...")
	basin, err := StartBasinNode(config)
	if err != nil {
		log.Fatal("Failed to instantiate the BasinNode: " + err.Error())
	}

	// Create the Router
	log.Println("Starting Router...")
	StartHardcodedRouter(basin.Host.Peerstore().PeerInfo(basin.Host.ID()))
	// _, err = StartKademliaRouter(ctx, basin.Host)
	// if err != nil {
	// 	log.Fatal("Failed to instantiate router: ", err.Error())
	// }

	// Create new PubSub
	log.Println("Creating PubSub...")
	_, err = StartPubSub(ctx, basin.Host)
	if err != nil {
		log.Fatal("Failed to instantiate pubsub: " + err.Error())
	}

	// Start up this node's HTTP API, concurrently with CLI
	log.Println("Serving HTTP API...")
	RunHttpServer(ctx, &basin, config.Http)
}
