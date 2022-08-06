package node

import (
	"context"
	"encoding/json"
	"log"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/sestinj/basin-node/pb"
	. "github.com/sestinj/basin-node/structs"
	"github.com/sestinj/basin-node/util"
)

func (b *BasinNode) readReqHandler(s network.Stream) {
	defer s.Close()

	log.Println("Recieved new read stream")

	data, err := readProtoMsg[*pb.ReadRequest](s)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Stream has requested the following URL: " + string(data.Url))

	// Sends back the same MessageData.Id so the response can be identified
	resp := &pb.ReadResponse{MessageData: &pb.MessageData{NodeId: b.Host.ID().String(), Id: data.MessageData.Id}, Data: nil}
	sig, err := b.signProtoMsg(resp)
	if err != nil {
		log.Println(err)
	}
	resp.MessageData.Sig = sig

	err = b.sendProtoMsg(s.Conn().RemotePeer(), s.Protocol(), resp)
	if err != nil {
		log.Println(err)
	}
}

func (b *BasinNode) readResHandler(s network.Stream) {
	defer s.Close()

	log.Println("New read response stream")

	data, err := readProtoMsg[*pb.ReadResponse](s)
	if err != nil {
		log.Println(err)
		return
	}

	// The reason we don't pass any errors through the channel is that the only errors occuring happen before we can get the channel.
	anchor, exists := b.ReadRequests[data.MessageData.Id]
	if !exists {
		return
	}

	anchor.Ch <- data
	close(anchor.Ch)
}

func (b *BasinNode) subResHandler(s network.Stream) {
	defer s.Close()

	log.Println("New subscription response stream")

	_, err := readProtoMsg[*pb.SubscriptionResponse](s)
	if err != nil {
		log.Println(err)
		return
	}

	log.Fatal("NOT YET IMPLEMENTED")
}

func (b *BasinNode) subReqHandler(s network.Stream) {
	defer s.Close()

	log.Println("New subscription request stream")

	data, err := readProtoMsg[*pb.SubscriptionRequest](s)
	if err != nil {
		log.Println(err)
		return
	}

	// Here is where you call the subscription decider plugin

	// Add this request to the list of requests
	ctx := context.Background()
	requests, err := b.GetRequests(ctx, "producer")
	if err != nil {
		log.Println(err)
		return
	}
	var capabilities []CapabilityJson
	for _, capability := range data.Capabilities {
		capabilities = append(capabilities, CapabilityJson{Expiration: capability.Expiration, Action: capability.Action})
	}
	*requests = append(*requests, PermissionJson{Entities: []string{data.MessageData.Did}, Data: []string{data.Url}, Capabilities: capabilities})
	url := util.GetUserDataUrl(b.Did, "producer.requests")
	val, err := json.Marshal(*requests)
	if err != nil {
		log.Println(err)
		return
	}
	err = b.WriteResource(ctx, url, val)
}

func (b *BasinNode) writeResHandler(s network.Stream) {
	log.Fatal("Not yet implemented")
}

func (b *BasinNode) writeReqHandler(s network.Stream) {
	log.Fatal("Not yet implemented")
}
