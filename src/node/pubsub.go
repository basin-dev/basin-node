package node

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	libp2p_host "github.com/libp2p/go-libp2p-core/host"
	"github.com/sestinj/basin-node/log"

	// kademlia "github.com/libp2p/go-libp2p-kad-dht"
	peer "github.com/libp2p/go-libp2p-core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"

	mdns "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

const DISCOVERY_NAME = "schema-topics"

type Update struct {
	Message  string
	SenderID string
}

type discoveryNotifee struct {
	h libp2p_host.Host
}

func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	fmt.Printf("discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}

func setupLANDiscovery(h libp2p_host.Host) error {
	service := mdns.NewMdnsService(h, DISCOVERY_NAME, &discoveryNotifee{h: h})
	return service.Start()
}

func joinSchemaTopic(ps *pubsub.PubSub) (*pubsub.Topic, *pubsub.Subscription) {
	// Join/create a pubsub topic
	topic, err := ps.Join("schema-topic:" + "test-topic")
	if err != nil {
		panic(err)
	}

	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}
	return topic, sub
}

func streamSubscription(ctx context.Context, sub *pubsub.Subscription, ch chan *Update, selfID peer.ID) {
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			log.Warning.Println(err.Error())
			// When a goroutine fails, better to close and return then to panic and destroy the entire application
			close(ch)
			return
		}

		// Uncomment except for super basic testing purposes
		// if msg.ReceivedFrom == selfID {
		// 	continue
		// }

		update := new(Update)
		err = json.Unmarshal(msg.Data, update)
		if err != nil {
			// This is why err handling is so explicit: you don't want to panic all the time
			continue
		}

		fmt.Println("Unmarshaled Update:", update.Message, "from", update.SenderID)

		ch <- update
	}
}

func printUpdateStream(ch chan *Update) {
	for {
		update := <-ch

		println(update.SenderID + ": " + update.Message)
	}
}

func sendMessage(ctx context.Context, topic *pubsub.Topic, msg string, selfID peer.ID) {
	update := Update{
		Message: msg, SenderID: selfID.String(),
	}
	data, err := json.Marshal(update)
	if err != nil {
		log.Warning.Println("Error serializing message: ", err.Error())
	}

	err = topic.Publish(ctx, data)
	if err != nil {
		log.Warning.Println("Error publishing to topic: ", err.Error())
	}
}

func periodicMsgs(ctx context.Context, topic *pubsub.Topic, selfID peer.ID) {
	ticker := time.NewTicker(10 * time.Second)

	for {
		fmt.Println("tick")
		time := <-ticker.C

		sendMessage(ctx, topic, "Test Message at "+time.String(), selfID)
	}
}

/* Instantiate a new libp2p PubSub */
func StartPubSub(ctx context.Context, host libp2p_host.Host) (*pubsub.PubSub, error) {
	ps, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// func main() {

// 	// Setup LAN Discovery
// 	if err := setupLANDiscovery(host); err != nil {
// 		println("lan err")
// 		panic(err)
// 	}

// 	topic, sub := joinSchemaTopic(ps)

// 	updateChan := make(chan *Update)

// 	println(topic.String())

// 	go streamSubscription(ctx, sub, updateChan, host.ID())

// 	go periodicMsgs(ctx, topic, host.ID())

// 	// I'd like this to be another goroutine, so what's the best way to have a main forever loop?
// 	// Right now just leaving the last as non-goroutine so the others get started
// 	printUpdateStream(updateChan)
// }
