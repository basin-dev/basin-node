package node

import (
	"context"
	"encoding/json"

	libp2p_host "github.com/libp2p/go-libp2p-core/host"
	"github.com/sestinj/basin-node/log"

	// kademlia "github.com/libp2p/go-libp2p-kad-dht"
	peer "github.com/libp2p/go-libp2p-core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const DISCOVERY_NAME = "basin"

// The Update message sent over pubsub
type Update struct {
	Url      string
	SenderID string
}

type discoveryNotifee struct {
	h libp2p_host.Host
}

func getFullTopicName(topicName string) string {
	return DISCOVERY_NAME + ":" + topicName
}

func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	log.Info.Printf("discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		log.Warning.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}

func JoinSchemaTopic(ps *pubsub.PubSub, topicName string) (*pubsub.Topic, *pubsub.Subscription) {
	// Join/create a pubsub topic
	topic, err := ps.Join(getFullTopicName(topicName))
	if err != nil {
		panic(err)
	}

	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}

	return topic, sub
}

func streamSubscription(ctx context.Context, sub *pubsub.Subscription, ch chan *Update) error {
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			log.Warning.Println("PubSub stream recieved an error: ", err.Error())
			// When a goroutine fails, better to close and return then to panic and destroy the entire application
			close(ch)
			return err
		}

		// Uncomment except for super basic testing purposes
		// if msg.ReceivedFrom == selfID {
		// 	continue
		// }

		update := new(Update)
		err = json.Unmarshal(msg.Data, update)
		if err != nil {
			continue
		}

		log.Info.Println("Unmarshaled Update:", update.Url, "from", update.SenderID)
		ch <- update
	}
}

// TODO[FEATURE][2] Run all of the notification handler plugins each time there is an update. Right now just hardwired in the printUpdate handler. Should have a struct and interface, then loop through a list of them.
// - you want to build a queue on the node though that keeps track of what updates there are, then the handler plugins cross these off of the list.
func runNotificationHandlers(ctx context.Context, ch chan *Update) error {
	for {
		update := <-ch
		printUpdate(update)
	}
}

func printUpdate(update *Update) {
	println(update.SenderID + ": " + update.Url)
}

/* Instantiate a new libp2p PubSub */
func StartPubSub(ctx context.Context, host libp2p_host.Host) (*pubsub.PubSub, error) {
	ps, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		return nil, err
	}

	return ps, nil
}
