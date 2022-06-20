package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	libp2p "github.com/libp2p/go-libp2p"
	libp2p_host "github.com/libp2p/go-libp2p-core/host"

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

func main() {
	// Parse command line flags
	username := *flag.String("name", "neura", "Username for authentication of user with node")
	pass := *flag.String("pass", "neura", "Password for authentication of user with node")
	flag.Parse()

	ctx := context.Background()

	// Create listener on port
	host, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		println("host err")
		panic(err)
	}

	// Create pubsub
	ps, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		println("ps err")
		panic(err)
	}

	// Setup LAN Discovery
	if err := setupLANDiscovery(host); err != nil {
		println("lan err")
		panic(err)
	}

	topic, sub := joinSchemaTopic(ps)

	updateChan := make(chan *Update)

	println(topic.String())

	go streamSubscription(ctx, sub, updateChan, host.ID())

	go periodicMsgs(ctx, topic, host.ID())

	// I'd like this to be another goroutine, so what's the best way to have a main forever loop?
	// Right now just leaving the last as non-goroutine so the others get started
	go printUpdateStream(updateChan)

	db := startDB()

	db.Put([]byte("node/username"), []byte(username), nil)
	db.Put([]byte("node/pass"), []byte(pass), nil)

	RunHTTPServer(db)

	// The need for consensus is an open problem right now. If people are just controlling their own data, then they should be the owners of it. And then a Byzantine Fault Tolerant algo isn't needed. But why must you be decentralized at that point? Perhaps for pubsub? Or is it instead for data replication for speed?
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

// func setupKademlia(ctx, context.Context, libp2p_host.Host) error {
// 	dht, err := kademlia.New(ctx, h)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return dht
// }

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
			println(err)
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
		println("Error serializing message.")
	}

	err = topic.Publish(ctx, data)
	if err != nil {
		fmt.Println("Error: ")
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

// package main

// func main() {
// 	println("a")

// 	loop("2")
// }

// func loop(msg string) {
// 	for {
// 		println(msg)
// 	}
// }
