package main

import (
	"bufio"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	. "github.com/libp2p/go-libp2p-gostream"
)

type BasinNode struct {
	Host host.Host
}

const Protocol = "/basin/test/1.0.0"

func StartBasinNode() (BasinNode, error) {
	// Create listener on port
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))

	h.SetStreamHandler(Protocol, streamHandler)

	basin := BasinNode{h}
	if err != nil {
		return basin, err
	}

	return basin, nil
}

func streamHandler(stream network.Stream) {
	defer stream.Close()

	log.Println("Recieved new stream")

	buf := bufio.NewReader(stream)
	val, _, err := buf.ReadLine()
	if err != nil {
		log.Println(err)
		stream.Reset()
	}
	log.Println("Stream has sent the following message: " + string(val))

	_, err = stream.Write(val)
}
