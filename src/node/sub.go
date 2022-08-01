package node

import (
	"context"

	"github.com/google/uuid"
	"github.com/sestinj/basin-node/pb"
	. "github.com/sestinj/basin-node/structs"
)

/* This is where the logic happens for a subscription request */
func (b *BasinNode) RequestSubscription(ctx context.Context, url string, capabilities []CapabilityJson) error {
	pi, err := HostRouter.ResolvePeer(ctx, url)
	if err != nil {
		return err
	}

	var capabilitiesProto []*pb.Capability
	for _, capability := range capabilities {
		capabilitiesProto = append(capabilitiesProto, &pb.Capability{Action: capability.Action, Expiration: capability.Expiration})
	}

	msg := &pb.SubscriptionRequest{
		Url:          url,
		Capabilities: capabilitiesProto,
		MessageData:  b.newMessageData(uuid.New().String()),
	}

	sig, err := b.signProtoMsg(msg)
	msg.MessageData.Sig = sig

	err = b.sendProtoMsg(pi.ID, ProtocolSubReq, msg)
	if err != nil {
		return err
	}

	return nil
}

// The subscription request handler should be a plugin—so it can happen manually, or through any automatic process as coded by the producer.
// It should be able to both answer on the spot and poll the list of requests.
