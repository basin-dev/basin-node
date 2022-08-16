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

	if data.Accepted {
		topic, sub := JoinSchemaTopic(b.Pubsub, data.Url)
		ch := make(chan *Update)
		ctx := context.Background()
		go runNotificationHandlers(ctx, ch)
		log.Warning.Printf("Streaming of topic '%s' ended unexpectedly: %v\n", topic.String(), streamSubscription(ctx, sub, ch))
		// TODO[FEATURE][1]: Keep track of subscriptions, both in a file and a map to the actual sub object. Should also facilitate resubscription upon node restart if persistent mode is turned on.
	} else {
		log.Info.Printf("Request for subscription to %s was rejected.", data.Url)
	}
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

	// FIXME[DEMO][2] Here is where you call the subscription decider plugin
	allowed = true

	// Add this request to the list of requests
	ctx := context.Background()
	requests, err := b.GetRequests(ctx, "producer")
	if err != nil {
		log.Warning.Println(err)
		return
	}

	// TODO[ucan, typegen][2] Need to update this whole function after dealing with UCANs. Fix the thing where you're only taking the first resource in the array
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

	// Sends back the same MessageData.Id so the response can be identified
	resp := &pb.SubscriptionResponse{MessageData: &pb.MessageData{NodeId: b.Host.ID().String(), Id: data.MessageData.Id}, Url: permissions[0].Data[0], Accepted: allowed}
	sig, err := b.signProtoMsg(resp)
	if err != nil {
		log.Warning.Println("Error signing message: ", err.Error())
	}
	resp.MessageData.Sig = sig

	err = b.sendProtoMsg(s.Conn().RemotePeer(), ProtocolSubRes, resp)
	if err != nil {
		log.Warning.Println("Error sending message: ", err)
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

	// TODO[ucan][3]: In progress: Create and respond with a UCAN
	// b.UcanStore.PutToken(ctx)
	// src, err := ucan.NewPrivKeySource(b.PrivKey)
	// ucan.NewTokenParser(ucan.AttenuationConstructorFunc())
	// if err != nil {
	// 	log.Warning.Println("Failed to create new UCAN private key source: ", err.Error())
	// }
}

func (b *BasinNode) writeResHandler(s network.Stream) {
	log.Error.Fatal("Not yet implemented")
}

func (b *BasinNode) writeReqHandler(s network.Stream) {
	log.Error.Fatal("Not yet implemented")
}
