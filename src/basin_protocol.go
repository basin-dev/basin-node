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

type BasinNode struct {
	Host         host.Host
	ReadRequests map[string]*pb.ReadRequest
}

const ProtocolReadReq = "/basin/read/1.0.0"
const ProtocolWriteReq = "/basin/write/1.0.0"

func StartBasinNode() (BasinNode, error) {
	// Create listener on port
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))

	basin := BasinNode{Host: h, ReadRequests: map[string]*pb.ReadRequest{}}
	if err != nil {
		return basin, err
	}

	h.SetStreamHandler(ProtocolReadReq, basin.readHandler)
	h.SetStreamHandler(ProtocolWriteReq, basin.writeHandler)

	return basin, nil
}

func (b *BasinNode) readHandler(s network.Stream) {
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

	resp := &pb.ReadResponse{MessageData: &pb.MessageData{NodeId: b.Host.ID().String()}, Data: nil}

	err = b.sendProtoMsg(s.Conn().RemotePeer(), s.Protocol(), resp)
	if err != nil {
		log.Println(err)
	}
}

func (b *BasinNode) writeHandler(s network.Stream) {}

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

func (b *BasinNode) NewMessageData() *pb.MessageData {
	return &pb.MessageData{Id: uuid.New().String(), NodeId: peer.Encode(b.Host.ID())}
}

func (b *BasinNode) ReadResource(ctx context.Context, url string) ([]byte, error) {
	if Contains(*GetSources("producer"), url) {
		// Determine which adapter to use
		// Should the file with info on how to call the adapter be stored in the adapter itself, or a local key/value, or a normal key/value?
		val, err := adapters.MainAdapter.Read(url) // TODO: Implement the MetaAdapter, which includes figuring out hooking up adapters
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return val, nil

		// This would also probably be the place to implement different guarantees?

	} else {
		// Use DNS/DHT to route to the node that produces this basin url
		pi, err := HostRouter.ResolvePeer(ctx, url)
		if err != nil {
			return nil, err
		}

		// TODO: Protobufs are more efficient, but not sure the best way to wait for response. Using HTTP rn.
		// Big problem with using HTTP is you have to have the ip4 multiaddr protocol for the node. Missing out on a lot of opportunities.
		// This is actually a huge problem, because not all nodes should have to have a domain, and their IP addresses will change, so the entry in the DHT will be outdated.
		// req := &pb.ReadRequest{Url: url, MessageData: b.NewMessageData()}

		// err = b.sendProtoMsg(pi.ID, ProtocolReadReq, req)
		// if err != nil {
		// 	return nil, err
		// }

		// b.ReadRequests[req.MessageData.Id] = req
	}

	return nil, errors.New("Not yet implemented")
}

func (b *BasinNode) WriteResource(ctx context.Context, url string, value []byte) error {
	// Do the same thing as ReadResource, if it's a local resource, just use the local adapter. And for now mostly everything should be.
	return errors.New("Not yet implemented")
}
