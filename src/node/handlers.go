package node

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/sestinj/basin-node/log"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/sestinj/basin-node/pb"
	. "github.com/sestinj/basin-node/structs"
	"github.com/sestinj/basin-node/util"
	"google.golang.org/protobuf/proto"
)

func (b *BasinNode) readReqHandler(s network.Stream) {
	defer s.Close()

	log.Info.Println("Recieved new read stream")

	data := &pb.ReadRequest{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Warning.Println("Failed to read stream: ", err.Error())
		return
	}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Warning.Println("Failed to unmarshal stream: ", err.Error())
		return
	}

	log.Info.Println("Stream has requested the following URL: " + string(data.Url))

	// TODO[ucan][2]: Implement permissions!!! UCANS! Should this be done in ReadResource instead, or is that too much overhead for local calls?
	ctx := context.Background()
	// permissionsUrl := util.GetMetadataUrl(string(data.Url), util.Permissions)
	// log.Println("Looking for perms at: ", permissionsUrl)

	// permissionsRaw, err := b.ReadResource(ctx, permissionsUrl)
	// if err != nil {
	// 	log.Println("Failed to read permissions for the resource")
	// 	return
	// }
	// var permissions []client.PermissionJson
	// err = json.Unmarshal(permissionsRaw, &permissions)
	// if err != nil {
	// 	log.Println("Error unmarshaling permissions for the resource")
	// }
	// log.Println("Permissions: ", permissions)
	// if len(permissions) == 0 {
	// 	log.Println("No access to resource")
	// 	return
	// }

	// Get the actual resource
	resource, err := b.ReadResource(ctx, string(data.Url))
	if err != nil {
		log.Warning.Println("Error reading the requested resource in readReqHandler")
		return
	}
	if !allowed {
		return
	}

	// Sends back the same MessageData.Id so the response can be identified
	resp := &pb.ReadResponse{MessageData: &pb.MessageData{NodeId: b.Host.ID().String(), Id: data.MessageData.Id}, Data: resource}
	sig, err := b.signProtoMsg(resp)
	if err != nil {
		log.Warning.Println("Error signing message: ", err.Error())
	}
	resp.MessageData.Sig = sig

	err = b.sendProtoMsg(s.Conn().RemotePeer(), ProtocolReadRes, resp)
	if err != nil {
		log.Warning.Println("Error sending message: ", err)
	}
}

func (b *BasinNode) readResHandler(s network.Stream) {
	log.Info.Println("New read response stream")
	defer s.Close()

	data := &pb.ReadResponse{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Warning.Println("Failed to read stream: ", err.Error())
		return
	}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Warning.Println("Failed to unmarshal stream: ", err.Error())
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

	log.Info.Println("New subscription response stream")

	data := &pb.SubscriptionResponse{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Warning.Println("Failed to read stream: ", err.Error())
		return
	}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Warning.Println("Failed to unmarshal stream: ", err.Error())
		return
	}

	log.Error.Println("NOT YET IMPLEMENTED")
}

var allowed bool = false

func (b *BasinNode) subReqHandler(s network.Stream) {
	defer s.Close()

	log.Info.Println("New subscription request stream")

	data := &pb.SubscriptionRequest{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Warning.Println("Failed to read stream: ", err.Error())
		return
	}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Warning.Println("Failed to unmarshal stream: ", err.Error())
		return
	}

	// Here is where you call the subscription decider plugin
	allowed = true

	// Add this request to the list of requests
	ctx := context.Background()
	requests, err := b.GetRequests(ctx, "producer")
	if err != nil {
		log.Warning.Println(err)
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
		log.Warning.Println(err)
		return
	}
	err = b.WriteResource(ctx, url, val)
	if err != nil {
		log.Warning.Println(err)
	}

	// TODO[ucan][2]: Automatically accepting for the demo. Also don't want to just search the file like the below. Use UCANs
	// newPs := []client.PermissionJson{{}}
	// rawP, err := json.Marshal(newPs)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// err = b.WriteResource(ctx, util.GetMetadataUrl(url, util.Permissions), rawP)
	// log.Println("wrote newPermissions to: ", util.GetMetadataUrl(url, util.Permissions))
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	log.Info.Println("Successfully recorded new request")
}

func (b *BasinNode) writeResHandler(s network.Stream) {
	log.Error.Fatal("Not yet implemented")
}

func (b *BasinNode) writeReqHandler(s network.Stream) {
	log.Error.Fatal("Not yet implemented")
}
