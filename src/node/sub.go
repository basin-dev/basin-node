package node

import (
	"context"

	"github.com/google/uuid"
	"github.com/sestinj/basin-node/client"
	"github.com/sestinj/basin-node/pb"
)

/* This is where the logic happens for a subscription request */
func (b *BasinNode) RequestSubscriptions(ctx context.Context, permissions []client.PermissionJson) error {
	// TODO[ucan][1]: Only going to request first URL for now.
	pi, err := HostRouter.ResolvePeer(ctx, (permissions)[0].Data[0])
	if err != nil {
		return err
	}

	var permissionsProto []*pb.Permission
	for _, permission := range permissions {
		var capabilities []*pb.Capability // FIXME[typegen][2]: : ( Again...need a single source for types...why do I even have to write this function? Should be abstracted much better.
		for _, capability := range permission.Capabilities {
			capabilities = append(capabilities, &pb.Capability{Action: *capability.Action, Expiration: capability.Expiration.String()})
		}
		permissionsProto = append(permissionsProto, &pb.Permission{Data: permission.Data, Capabilities: capabilities, Entities: permission.Entities})
	}

	msg := &pb.SubscriptionRequest{
		Permissions: permissionsProto,
		MessageData: b.newMessageData(uuid.New().String()),
	}

	sig, err := b.signProtoMsg(msg)
	if err != nil {
		return err
	}
	msg.MessageData.Sig = sig

	err = b.sendProtoMsg(pi.ID, ProtocolSubReq, msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *BasinNode) RequestSingleSubscription(ctx context.Context, url string, permissions []client.PermissionJson) error {
	return nil // See the todo above
}

// The subscription request handler should be a plugin—so it can happen manually, or through any automatic process as coded by the producer.
// It should be able to both answer on the spot and poll the list of requests.
