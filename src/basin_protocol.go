package main

import (
	"context"
	"errors"
	"io/ioutil"
	"log"

	ggio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sestinj/basin-node/adapters"
	"github.com/sestinj/basin-node/pb"
	. "github.com/sestinj/basin-node/util"
)

type ReadReqAnchor struct {
	Req *pb.ReadRequest
	Ch  chan *pb.ReadResponse
}

type BasinNode struct {
	Host         host.Host
	ReadRequests map[string]*ReadReqAnchor
}

const ProtocolReadReq = "/basin/readreq/1.0.0"
const ProtocolReadRes = "/basin/readres/1.0.0"

// TODO: ProtocolWriteReq/Res and associated handlers

func StartBasinNode() (BasinNode, error) {
	// Create listener on port
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))

	basin := BasinNode{Host: h, ReadRequests: map[string]*ReadReqAnchor{}}
	if err != nil {
		return basin, err
	}

	h.SetStreamHandler(ProtocolReadReq, basin.readReqHandler)
	h.SetStreamHandler(ProtocolReadRes, basin.readResHandler)

	return basin, nil
}

func (b *BasinNode) readReqHandler(s network.Stream) {
	defer s.Close()

	log.Println("Recieved new read stream")

	data := new(pb.ReadRequest)
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Println(err)
		return
	}
	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Stream has requested the following URL: " + string(data.Url))

	// Sends back the same MessageData.Id so the response can be identified
	resp := &pb.ReadResponse{MessageData: &pb.MessageData{NodeId: b.Host.ID().String(), Id: data.MessageData.Id}, Data: nil}

	err = b.sendProtoMsg(s.Conn().RemotePeer(), s.Protocol(), resp)
	if err != nil {
		log.Println(err)
	}
}

func (b *BasinNode) readResHandler(s network.Stream) {
	defer s.Close()

	log.Println("New read response stream")

	data := new(pb.ReadResponse)
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Println(err)
		return
	}
	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println(err)
		return
	}

	anchor, exists := b.ReadRequests[data.MessageData.Id]
	if !exists {
		return
	}

	anchor.Ch <- data
	close(anchor.Ch)
}

func (b *BasinNode) writeResHandler(s network.Stream) {
	log.Fatal("Not yet implemented")
}

func (b *BasinNode) writeReqHandler(s network.Stream) {
	log.Fatal("Not yet implemented")
}

func (b *BasinNode) sendProtoMsg(id peer.ID, p protocol.ID, data proto.Message) error {
	s, err := b.Host.NewStream(context.Background(), id, p)
	if err != nil {
		log.Println(err)
		return err
	}
	defer s.Close()

	writer := ggio.NewFullWriter(s)
	err = writer.WriteMsg(data)
	if err != nil {
		log.Println(err)
		s.Reset()
		return err
	}

	return nil
}

func (b *BasinNode) NewMessageData(id string) *pb.MessageData {
	return &pb.MessageData{Id: id, NodeId: peer.Encode(b.Host.ID())}
}

func (b *BasinNode) ReadResource(ctx context.Context, url string) (chan<- []byte, error) {
	valCh := make(chan<- []byte)
	if Contains(*GetSources("producer"), url) {
		// Determine which adapter to use
		// Should the file with info on how to call the adapter be stored in the adapter itself, or a local key/value, or a normal key/value?
		resCh, err := adapters.MainAdapter.Read(url) // TODO: Implement the MetaAdapter, which includes figuring out hooking up adapters
		if err != nil {
			log.Println(err)
			return nil, err
		}

		go func() {
			valCh <- resCh
			close(valCh)
		}()

		return valCh, nil
	} else {
		// Use DHT to route to the node that produces this basin url
		pi, err := HostRouter.ResolvePeer(ctx, url)
		if err != nil {
			return nil, err
		}

		// TODO: Protobufs are more efficient, but not sure the best way to wait for response. Using HTTP rn.
		// Big problem with using HTTP is you have to have the ip4 multiaddr protocol for the node. Missing out on a lot of opportunities.
		// This is actually a huge problem, because not all nodes should have to have a domain, and their IP addresses will change, so the entry in the DHT will be outdated.
		req := &pb.ReadRequest{Url: url, MessageData: b.NewMessageData(uuid.New().String())}

		err = b.sendProtoMsg(pi.ID, ProtocolReadReq, req)
		if err != nil {
			return nil, err
		}

		resCh := make(chan *pb.ReadResponse)

		anchor := &ReadReqAnchor{Req: req, Ch: resCh}

		b.ReadRequests[req.MessageData.Id] = anchor

		log.Println("Waiting for response to id " + req.MessageData.Id)

		go func() {
			// Wait for the response to come back through the channel. TODO: Maximum wait time (this should be solved at a different layer probably, which will take care of retries and everything)
			res := <-resCh
			log.Println("Recieved response for request id " + res.MessageData.Id)
			valCh <- res.Data
			close(valCh)
		}()

		return valCh, nil
	}
}

func (b *BasinNode) WriteResource(ctx context.Context, url string, value []byte) error {
	// Do the same thing as ReadResource, if it's a local resource, just use the local adapter. And for now mostly everything should be.
	return errors.New("Not yet implemented")
}
