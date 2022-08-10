package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/sestinj/basin-node/adapters"
	. "github.com/sestinj/basin-node/node"
	server "github.com/sestinj/basin-node/server/go"
)

func RunHttpServer(ctx context.Context, b *BasinNode, addr string) {
	DefaultApiService := server.NewDefaultApiService()
	DefaultApiController := server.NewDefaultApiController(DefaultApiService)

	router := server.NewRouter(DefaultApiController)

	segs := strings.Split(b.Http, "://")
	portHost := segs[len(segs)-1]
	fmt.Fprintf(os.Stdout, "Listening at %s...\n", b.Http)
	log.Fatal(http.ListenAndServe(portHost, router))
}

func StartEverything(ctx context.Context, config BasinNodeConfig) {
	// NOTE: The order here matters. For example, db must start before node.

	// Start up the local LevelDB database
	log.Println("Initializing LevelDB...")
	_, err := adapters.StartDB(config.Http)
	if err != nil {
		log.Fatal(err)
	}

	// Start the BasinNode (libp2p host with associated protocol, stream handler)
	log.Println("Launching Basin Node...")
	basin, err := StartBasinNode(config)
	if err != nil {
		log.Fatal("Failed to instantiate the BasinNode: " + err.Error())
	}

	// Create the Router
	log.Println("Starting Router...")
	info := basin.Host.Peerstore().PeerInfo(basin.Host.ID())
	StartHardcodedRouter(info)
	// _, err = StartKademliaRouter(ctx, basin.Host)
	// if err != nil {
	// 	log.Fatal("Failed to instantiate router: ", err.Error())
	// }

	// Setup Discovery
	log.Println("Setting up mDNS discovery...")
	err = setupDiscovery(basin.Host)
	if err != nil {
		log.Fatal("Failed to start mDNS discovery: ", err.Error())
	}

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

// DiscoveryServiceTag is used in our mDNS advertisements to discover other peers.
const DiscoveryServiceTag = "basin-pubsub"

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	fmt.Printf("discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
	HostRouter.Peer = pi
}

func setupDiscovery(h host.Host) error {
	// setup mDNS discovery to find local peers
	s := mdns.NewMdnsService(h, DiscoveryServiceTag, &discoveryNotifee{h: h})
	return s.Start()
}
