package main

import (
	"context"
	"encoding/json"

	libp2p_host "github.com/libp2p/go-libp2p-core/host"
	libp2p_peer "github.com/libp2p/go-libp2p-core/peer"
	kademlia "github.com/libp2p/go-libp2p-kad-dht"
)

var (
	HostRouter KademliaRouter
)

/*
A Router is responsible for resolving a Basin URL to a host on the network, that can then be connected to.
*/
type Router interface {
	ResolvePeer(ctx context.Context, url string) (libp2p_peer.AddrInfo, error)
}

/* PLAN FOR NOW is to simply keep a global Kademlia DHT, and have the producer node write their info into the value for whatever url they are registering. */
type KademliaRouter struct {
	dht  *kademlia.IpfsDHT
	host libp2p_host.Host
}

/* Responsible for converting a Basin URL to info about a peer that be connected to */
func (k KademliaRouter) ResolvePeer(ctx context.Context, url string) (libp2p_peer.AddrInfo, error) {
	data, err := k.dht.GetValue(ctx, url, nil)

	addrInfo := new(libp2p_peer.AddrInfo)
	err = json.Unmarshal(data, addrInfo)
	if err != nil {
		return *addrInfo, err
	}
	return *addrInfo, err
}

/* This function is specific to the "KademliaRouter", our first simple version. It writes itself as the peer to contact for a Basin URL. This should be called when a new schema is registered by this node. */
func (k KademliaRouter) RegisterUrl(ctx context.Context, url string) error {
	addrInfo := libp2p_peer.AddrInfo{ID: k.host.ID(), Addrs: k.host.Addrs()}
	val, err := json.Marshal(addrInfo)
	if err != nil {
		return err
	}

	err = k.dht.PutValue(ctx, url, val)

	return err
}

/* Instantiate a Router instance */
func StartKademliaRouter(ctx context.Context, h libp2p_host.Host) (Router, error) {
	dht, err := kademlia.New(ctx, h)
	if err != nil {
		return nil, err
	}

	HostRouter = KademliaRouter{dht: dht, host: h}
	return HostRouter, nil
}
