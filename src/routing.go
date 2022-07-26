package main

import (
	"context"

	libp2p_host "github.com/libp2p/go-libp2p-core/host"
	kademlia "github.com/libp2p/go-libp2p-kad-dht"
)

/*
A Router is responsible for resolving Basin URLs to IP addresses.
*/
type Router interface {
	ResolveIp(url string) uint32
}

type KademliaRouter struct {
	dht *kademlia.IpfsDHT
}

/* Instantiate a new KademliaRouter */
func (k KademliaRouter) New(ctx context.Context, h libp2p_host.Host) (KademliaRouter, error) {
	dht, err := kademlia.New(ctx, h)

	return KademliaRouter{dht}, err
}

/* Responsible for converting a Basin URL to a  */
func (k KademliaRouter) ResolveIp(url string) (uint32, error) {
	k.dht.
}
