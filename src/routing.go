package main

import (
	"context"

	libp2p_host "github.com/libp2p/go-libp2p-core/host"
	kademlia "github.com/libp2p/go-libp2p-kad-dht"
)

/*
A Router is responsible for resolving a Basin URL to a host on the network, that can then be connected to.
*/
type Router interface {
	ResolvePeer(url string) (libp2p_host.Host, error)
}

/* PLAN FOR NOW is to simply keep a global Kademlia DHT, and have the producer node write their info into the value for whatever url they are registering. */
type KademliaRouter struct {
	dht *kademlia.IpfsDHT
}

/* Instantiate a new KademliaRouter */
func (k KademliaRouter) New(ctx context.Context, h libp2p_host.Host) (KademliaRouter, error) {
	dht, err := kademlia.New(ctx, h)

	return KademliaRouter{dht}, err
}

/* Responsible for converting a Basin URL to a  */
func (k KademliaRouter) ResolvePeer(url string) (libp2p_host.Host, error) {
	k.dht.
}
