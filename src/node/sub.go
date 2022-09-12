package node

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sestinj/basin-node/client"
	didutil "github.com/sestinj/basin-node/did"
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

// This function does not handle the permissions requests user metadata file. It is up to another function to search through that and call this one, and also up to that function to delete the entries in the permissions requests file. This function creates a UCAN token, appends the permissions to the permissions file (there's code for this in stuff???), and sends the message to the requester
func (b *BasinNode) GrantSubscription(ctx context.Context, url string, permissions []client.PermissionJson) error {
	// Create UCAN token
	key, err := crypto.UnmarshalEd25519PrivateKey(b.PrivKey)
	if err != nil {
		return fmt.Errorf("Error granting subscription: %w\n", err)
	}

	_, err = didutil.CreateAttenuatedToken(ctx, url, key, permissions[0].Entities[0], *permissions[0].Capabilities[0].Expiration, []string{*permissions[0].Capabilities[0].Action})
	if err != nil {
		return fmt.Errorf("Error creating token in grant subscription: %w\n", err)
	}

	// Join permissions to the metadata file

	// Send the response message to requester

	return nil
}
