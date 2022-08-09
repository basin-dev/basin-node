package node

import (
	"context"
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"io/ioutil"
	"log"
	"strings"

	ggio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sestinj/basin-node/pb"
)

func readProtoMsg[T proto.Message](s network.Stream) (T, error) {
	data := new(T)
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		return *data, err
	}
	err = proto.Unmarshal(buf, *data)
	if err != nil {
		return *data, err
	}

	return *data, nil
}

/* Use the node's current private key to sign the message data and return the signature */
func (b *BasinNode) signProtoMsg(data proto.Message) ([]byte, error) {
	digest, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}
	sig, err := b.PrivKey.Sign(rand.Reader, digest, crypto.Hash(0))
	if err != nil {
		log.Println("Failed to sign message with private key: ", err.Error())
		return nil, err
	}
	return sig, nil
}

/* Verify that the signed message was signed by the DID specified in the MessageData */
func verifyMessage(msg proto.Message, msgData *pb.MessageData) bool {
	// Must set this to nil because it is nil when the message is signed
	sig := msgData.Sig
	msgData.Sig = nil

	raw, err := proto.Marshal(msg)
	if err != nil {
		log.Println("Failed to unmarshal proto message.")
		return false
	}
	msgData.Sig = sig
	pub := []byte(strings.Replace(msgData.Did, "did:key:", "", 1))
	return ed25519.Verify(pub, raw, sig)
}

/* Create a new stream with specified peer and write the protobuf message to the stream under given protocol */
func (b *BasinNode) sendProtoMsg(id peer.ID, p protocol.ID, data proto.Message) error {
	log.Println("PEER INFO: ", id, b.Host.Peerstore().PeerInfo(id))
	s, err := b.Host.NewStream(context.Background(), id, p)
	if err != nil {
		log.Println(err)
		return err
	}
	defer s.Close()

	writer := ggio.NewFullWriter(s)
	err = writer.WriteMsg(data)
	if err != nil {
		log.Println(err)
		s.Reset()
		return err
	}

	return nil
}

/* Create new metadata struct for protobuf message */
func (b *BasinNode) newMessageData(id string) *pb.MessageData {
	return &pb.MessageData{Id: id, NodeId: peer.Encode(b.Host.ID())}
}
