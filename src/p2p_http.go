/* The nodes of our network communicate with each other over HTTP, through a different interface than the externally facing one defined in http_interface.go. This file contains both the client and server code. */

package main

import (
	"context"
	"errors"

	"net/http"

	libp2p_host "github.com/libp2p/go-libp2p-core/host"
	. "github.com/libp2p/go-libp2p-gostream"
	p2phttp "github.com/libp2p/go-libp2p-http"
)

const protocolId = "libp2p"

func StartClient(ctx context.Context, h libp2p_host.Host) *http.Client {
	tr := &http.Transport{}

	// TODO: Take another look at this RegisterProtocol function...looks like it can do some really useful stuff
	tr.RegisterProtocol("basin", p2phttp.NewTransport(h))
	client := &http.Client{Transport: tr}

	return client
}

func StartP2pHttp(ctx context.Context, h libp2p_host.Host) error {
	listener, err := Listen(h, p2phttp.DefaultP2PProtocol)
	if err != nil {
		return err
	}
	defer listener.Close()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi!"))
	})

	server := &http.Server{}
	server.Serve(listener)

	return errors.New("P2P HTTP Server has been closed.")
}
