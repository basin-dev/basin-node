package node

import (
	"context"
	"encoding/json"

	"github.com/ipfs/go-ipns"
	"github.com/sestinj/basin-node/log"

	libp2p_host "github.com/libp2p/go-libp2p-core/host"
	libp2p_peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	kademlia "github.com/libp2p/go-libp2p-kad-dht"
	record "github.com/libp2p/go-libp2p-record"
)

var (
	HostRouter KademliaRouter
)

/*
A Router is responsible for resolving a Basin URL to a host on the network, that can then be connected to.
*/
type Router interface {
	ResolvePeer(ctx context.Context, url string) (libp2p_peer.AddrInfo, error)
	RegisterUrl(ctx context.Context, url string) error
}

/* A stupid simple router for early testing. Can do this because routers are just a plugin. */
type HardcodedRouter struct {
	Peer libp2p_peer.AddrInfo
}

func (h HardcodedRouter) ResolvePeer(ctx context.Context, url string) (libp2p_peer.AddrInfo, error) {
	return h.Peer, nil
}

func (h HardcodedRouter) RegisterUrl(ctx context.Context, url string) error {
	return nil
}

/* PLAN FOR NOW is to simply keep a global Kademlia DHT, and have the producer node write their info into the value for whatever url they are registering. */
type KademliaRouter struct {
	dht  *kademlia.IpfsDHT
	host libp2p_host.Host
}

const DHT_NAMESPACE = "/basin"

/* Responsible for converting a Basin URL to info about a peer that be connected to */
func (k KademliaRouter) ResolvePeer(ctx context.Context, url string) (libp2p_peer.AddrInfo, error) {
	data, err := k.dht.GetValue(ctx, DHT_NAMESPACE+"/"+url, routing.Expired)

	pi := new(libp2p_peer.AddrInfo)
	err = json.Unmarshal(data, pi)
	if err != nil {
		return *pi, err
	}
	return *pi, err
}

/* This function is specific to the "KademliaRouter", our first simple version. It writes itself as the peer to contact for a Basin URL. This should be called when a new schema is registered by this node. */
func (k KademliaRouter) RegisterUrl(ctx context.Context, url string) error {
	// TODO[FEATURE][1]: Should be signing the records with the private key of the DID in the url
	pi := libp2p_peer.AddrInfo{ID: k.host.ID(), Addrs: k.host.Addrs()}
	val, err := json.Marshal(pi)
	if err != nil {
		return err
	}

	err = k.dht.PutValue(ctx, DHT_NAMESPACE+"/"+url, val)

	if err != nil {
		log.Warning.Printf("Error writing key value pair ('%s', '%s') to Kademlia DHT: %w\n", DHT_NAMESPACE+"/"+url, string(val), err)
	}

	return err
}

/* Instantiate a Router instance */
func StartKademliaRouter(ctx context.Context, h libp2p_host.Host) (Router, error) {
	dht, err := kademlia.New(ctx, h, kademlia.ProtocolPrefix(DHT_NAMESPACE), kademlia.Mode(kademlia.ModeServer)) // FIXME: ModeServer is just for testing purposes
	if err != nil {
		return nil, err
	}

	dht.Validator = record.NamespacedValidator{
		"pk":    record.PublicKeyValidator{},
		"ipns":  ipns.Validator{KeyBook: h.Peerstore()},
		"basin": CustomValidator{},
	}

	HostRouter = KademliaRouter{dht: dht, host: h}
	return HostRouter, nil
}

type CustomValidator struct{}

// Validate validates the given record, returning an error if it's
// invalid (e.g., expired, signed by the wrong key, etc.).
func (v CustomValidator) Validate(key string, value []byte) error {
	return nil
}

// Select selects the best record from the set of records (e.g., the
// newest).
//
// Decisions made by select should be stable.
func (v CustomValidator) Select(key string, values [][]byte) (int, error) {
	return 0, nil
}
