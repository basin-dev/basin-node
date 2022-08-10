package node

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/sestinj/basin-node/pb"
	. "github.com/sestinj/basin-node/structs"
	"github.com/sestinj/basin-node/util"
	"google.golang.org/protobuf/proto"
)

func (b *BasinNode) readReqHandler(s network.Stream) {
	defer s.Close()

	log.Println("Recieved new read stream")

	data := &pb.ReadRequest{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Println("Failed to read stream: ", err.Error())
		return
	}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println("Failed to unmarshal stream: ", err.Error())
		return
	}

	log.Println("Stream has requested the following URL: " + string(data.Url))

	// Get the actual resource
	resource, err := b.ReadResource(context.Background(), string(data.Url))
	if err != nil {
		log.Println("Error reading the requested resource in readReqHandler")
		return
	}

	// Sends back the same MessageData.Id so the response can be identified
	resp := &pb.ReadResponse{MessageData: &pb.MessageData{NodeId: b.Host.ID().String(), Id: data.MessageData.Id}, Data: resource}
	sig, err := b.signProtoMsg(resp)
	if err != nil {
		log.Println(err)
	}
	resp.MessageData.Sig = sig

	err = b.sendProtoMsg(s.Conn().RemotePeer(), ProtocolReadRes, resp)
	if err != nil {
		log.Println(err)
	}
}

func (b *BasinNode) readResHandler(s network.Stream) {
	log.Println("New read response stream")
	defer s.Close()

	data := &pb.ReadResponse{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Println("Failed to read stream: ", err.Error())
		return
	}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println("Failed to unmarshal stream: ", err.Error())
		return
	}

	// The reason we don't pass any errors through the channel is that the only errors occuring happen before we can get the channel.
	anchor, exists := b.ReadRequests[data.MessageData.Id]
	if !exists {
		return
	}

	anchor.Ch <- data
}

func (b *BasinNode) subResHandler(s network.Stream) {
	defer s.Close()

	log.Println("New subscription response stream")

	data := &pb.SubscriptionResponse{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Println("Failed to read stream: ", err.Error())
		return
	}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println("Failed to unmarshal stream: ", err.Error())
		return
	}

	log.Println("NOT YET IMPLEMENTED")
}

func (b *BasinNode) subReqHandler(s network.Stream) {
	defer s.Close()

	log.Println("New subscription request stream")

	data := &pb.SubscriptionRequest{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Println("Failed to read stream: ", err.Error())
		return
	}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println("Failed to unmarshal stream: ", err.Error())
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

	var permissions []PermissionJson
	for _, permission := range data.Permissions {
		var capabilities []CapabilityJson
		for _, capability := range permission.Capabilities {
			capabilities = append(capabilities, CapabilityJson{Expiration: capability.Expiration, Action: capability.Action})
		}
		permissions = append(permissions, PermissionJson{Data: permission.Data, Capabilities: capabilities, Entities: permission.Entities})
	}

	*requests = append(*requests, permissions...)
	url := util.GetUserDataUrl(b.Did, "producer.requests")
	val, err := json.Marshal(*requests)
	if err != nil {
		log.Println(err)
		return
	}
	err = b.WriteResource(ctx, url, val)
	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully recorded new request")
}

func (b *BasinNode) writeResHandler(s network.Stream) {
	log.Fatal("Not yet implemented")
}

func (b *BasinNode) writeReqHandler(s network.Stream) {
	log.Fatal("Not yet implemented")
}
