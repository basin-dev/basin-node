package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	ggio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sestinj/basin-node/adapters"
	"github.com/sestinj/basin-node/pb"
	. "github.com/sestinj/basin-node/structs"
	. "github.com/sestinj/basin-node/util"
	"golang.org/x/sync/errgroup"
)

type ReadReqAnchor struct {
	Req *pb.ReadRequest
	Ch  chan *pb.ReadResponse
}

type BasinNode struct {
	Host         host.Host
	ReadRequests map[string]*ReadReqAnchor
}

const ProtocolReadReq = "/basin/readreq/1.0.0"
const ProtocolReadRes = "/basin/readres/1.0.0"

// TODO: ProtocolWriteReq/Res and associated handlers

func StartBasinNode() (BasinNode, error) {
	// Create listener on port
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))

	basin := BasinNode{Host: h, ReadRequests: map[string]*ReadReqAnchor{}}
	if err != nil {
		return basin, err
	}

	h.SetStreamHandler(ProtocolReadReq, basin.readReqHandler)
	h.SetStreamHandler(ProtocolReadRes, basin.readResHandler)

	return basin, nil
}

func (b *BasinNode) readReqHandler(s network.Stream) {
	defer s.Close()

	log.Println("Recieved new read stream")

	data := new(pb.ReadRequest)
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Println(err)
		return
	}
	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Stream has requested the following URL: " + string(data.Url))

	// Sends back the same MessageData.Id so the response can be identified
	resp := &pb.ReadResponse{MessageData: &pb.MessageData{NodeId: b.Host.ID().String(), Id: data.MessageData.Id}, Data: nil}

	err = b.sendProtoMsg(s.Conn().RemotePeer(), s.Protocol(), resp)
	if err != nil {
		log.Println(err)
	}
}

func (b *BasinNode) readResHandler(s network.Stream) {
	defer s.Close()

	log.Println("New read response stream")

	data := new(pb.ReadResponse)
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		log.Println(err)
		return
	}
	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println(err)
		return
	}

	// The reason we don't pass any errors through the channel is that the only errors occuring happen before we can get the channel.
	anchor, exists := b.ReadRequests[data.MessageData.Id]
	if !exists {
		return
	}

	anchor.Ch <- data
	close(anchor.Ch)
}

func (b *BasinNode) writeResHandler(s network.Stream) {
	log.Fatal("Not yet implemented")
}

func (b *BasinNode) writeReqHandler(s network.Stream) {
	log.Fatal("Not yet implemented")
}

func (b *BasinNode) sendProtoMsg(id peer.ID, p protocol.ID, data proto.Message) error {
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

func (b *BasinNode) NewMessageData(id string) *pb.MessageData {
	return &pb.MessageData{Id: id, NodeId: peer.Encode(b.Host.ID())}
}

func (b *BasinNode) ReadResource(ctx context.Context, url string) ([]byte, error) {
	if Contains(*b.GetSources(ctx, "producer"), url) {
		// Determine which adapter to use
		// Should the file with info on how to call the adapter be stored in the adapter itself, or a local key/value, or a normal key/value?
		return adapters.MainAdapter.Read(url) // TODO: Implement the MetaAdapter, which includes figuring out hooking up adapters
	} else {
		// Use DHT to route to the node that produces this basin url
		pi, err := HostRouter.ResolvePeer(ctx, url)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// TODO: Protobufs are more efficient, but not sure the best way to wait for response. Using HTTP rn.
		// Big problem with using HTTP is you have to have the ip4 multiaddr protocol for the node. Missing out on a lot of opportunities.
		// This is actually a huge problem, because not all nodes should have to have a domain, and their IP addresses will change, so the entry in the DHT will be outdated.
		req := &pb.ReadRequest{Url: url, MessageData: b.NewMessageData(uuid.New().String())}

		err = b.sendProtoMsg(pi.ID, ProtocolReadReq, req)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		resCh := make(chan *pb.ReadResponse)
		defer close(resCh)

		anchor := &ReadReqAnchor{Req: req, Ch: resCh}
		b.ReadRequests[req.MessageData.Id] = anchor
		log.Println("Waiting for response to id " + req.MessageData.Id)

		// Wait for the response to come back through the channel.
		// TODO: Maximum wait time (this should be solved at a different layer probably, which will take care of retries and everything. But note that we're not getting errors here throught the channel)
		res := <-resCh
		log.Println("Recieved response for request id " + res.MessageData.Id)

		return res.Data, nil
	}
}

func (b *BasinNode) WriteResource(ctx context.Context, url string, value []byte) error {
	return adapters.MainAdapter.Write(url, value)
	// Do the same thing as ReadResource, if it's a local resource, just use the local adapter. And for now mostly everything should be.
}

// Working on making the metadata appear...
func (b *BasinNode) Register(ctx context.Context, manifestPath string) error {
	// A couple of todos for later...
	// 1. TODO: Make sure did owns the domain
	// 2. TODO: Check whether a schema already exists at this domain. If so, version it.
	// For now we'll assume that the URL by itself returns newest version, but later this might have to be
	// done more explicity. Consider how one might request an older version. Is this a header, part of the path or query?

	manifestRaw, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return err
	}

	manifest := new(ManifestJson)
	err = json.Unmarshal(manifestRaw, manifest)
	if err != nil {
		return err
	}

	// TODO: First, check whether a manifest already exists (whether we are creating a new version or just registering for the first time)
	// For now always assume that all registers are first time, and overwrite each other.

	// Run all the file writes in parallel
	g, ctx := errgroup.WithContext(ctx)

	// PERMISSIONS
	permUrl := GetMetadataUrl(manifest.Url, Permissions)
	perms := []PermissionJson{}
	if manifest.PublicRead {
		// If public, then create a statement allowing all
		// Otherwise, initial permissions are none
		perm := PermissionJson{
			Data: []string{},
			Capabilities: []CapabilityJson{
				CapabilityJson{
					Action:     "read",
					Expiration: "never",
				},
			},
			Entities: []string{"*"},
		}
		perms = append(perms, perm)
	}

	permsRaw, err := json.Marshal(perms)
	if err != nil {
		return err
	}

	g.Go(func() error { return b.WriteResource(ctx, permUrl, permsRaw) })

	// SCHEMA
	schemaUrl := GetMetadataUrl(manifest.Url, Schema)
	schemaRaw, err := json.Marshal(manifest.Schema) // TODO: What is the shape of the schema?
	g.Go(func() error { return b.WriteResource(ctx, schemaUrl, schemaRaw) })

	// MANIFEST
	manifestUrl := GetMetadataUrl(manifest.Url, Manifest)
	// TODO: Note that right here we just loaded a file from the filesystem and threw it into LevelDB
	// This is when we want to start storing things as actual files? Just start thinking about it.
	g.Go(func() error { return b.WriteResource(ctx, manifestUrl, manifestRaw) })

	// SOURCES
	walletInfo := b.GetWalletInfo()
	sourcesUrl := GetUserDataUrl(walletInfo.Did, "producer.sources")
	currSrcs, err := LocalOnlyDb.Read(sourcesUrl)
	var srcs []string
	err = json.Unmarshal(currSrcs, srcs)
	if err != nil {
		return err
	}
	srcs = append(srcs, manifest.Url)
	finalSrcs, err := json.Marshal(srcs)
	if err != nil {
		return err
	}
	g.Go(func() error { return b.WriteResource(ctx, sourcesUrl, finalSrcs) })

	// Register with the routing table
	err = HostRouter.RegisterUrl(ctx, manifest.Url)
	if err != nil {
		return err
	}

	// Just like any other update - should tell subscribers (want a function for this)

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (b *BasinNode) GetWalletInfo() *WalletInfoJson {
	data, err := LocalOnlyDb.Read("wallet")
	if err != nil {
		log.Fatal(err)
	}

	return Unmarshal[WalletInfoJson](data)
}

func (b *BasinNode) GetPermissions(ctx context.Context, dataUrl string) (*[]PermissionJson, error) {
	url := GetMetadataUrl(dataUrl, Permissions)
	val, err := b.ReadResource(ctx, url)
	if err != nil {
		return nil, err
	}
	return Unmarshal[[]PermissionJson](val), nil
}

func (b *BasinNode) GetSchema(ctx context.Context, dataUrl string) (*SchemaJson, error) {
	url := GetMetadataUrl(dataUrl, Schema)
	val, err := b.ReadResource(ctx, url)
	if err != nil {
		return nil, err
	}
	return Unmarshal[SchemaJson](val), nil
}

func (b *BasinNode) GetSources(ctx context.Context, mode string) (*[]string, error) {
	walletInfo := b.GetWalletInfo()

	url := GetUserDataUrl(walletInfo.Did, mode+".sources")
	val, err := b.ReadResource(ctx, url)
	if err != nil {
		return nil, err
	}
	// TODO: You've made an infinite loop with ReadResource and GetSources :(
	return Unmarshal[[]string](val), nil
}

func (b *BasinNode) GetRequests(ctx context.Context, mode string) (*[]PermissionJson, error) {
	walletInfo := b.GetWalletInfo()

	url := GetUserDataUrl(walletInfo.Did, mode+".requests")
	val, err := b.ReadResource(ctx, url)
	if err != nil {
		return nil, err
	}

	return Unmarshal[[]PermissionJson](val), nil
}

func (b *BasinNode) GetSchemas(ctx context.Context, mode string) (*[]SchemaJson, error) {
	srcs, err := b.GetSources(ctx, mode)
	if err != nil {
		return nil, err
	}

	var schemas []SchemaJson
	g, ctx := errgroup.WithContext(ctx)
	for _, source := range *srcs {
		g.Go(func() error {
			schema, err := b.GetSchema(ctx, source)
			if err != nil {
				return err
			}
			schemas = append(schemas, *schema)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &schemas, nil
}

func (b *BasinNode) RequestSubscription(ctx context.Context, url string) error {
	return nil
}

// TODO: Realized I have the channel pattern inside out: want to allow the function to take time (doesn't need to return channel immediately) and t
// hen turn it into a goroutine in the calling function if needed. This is the more flexible way, and much less verbose. Will clean soon"
