package node

import (
	"context"
	"crypto/ed25519"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/sestinj/basin-node/adapters"
	"github.com/sestinj/basin-node/client"
	didutil "github.com/sestinj/basin-node/did"
	"github.com/sestinj/basin-node/pb"
	. "github.com/sestinj/basin-node/structs"
	. "github.com/sestinj/basin-node/util"
	"golang.org/x/sync/errgroup"
)

type ReadReqAnchor struct {
	Req *pb.ReadRequest
	Ch  chan *pb.ReadResponse
}

/* TODO: All plugins should be here. There should be a Router(s?) and Adapter (for the metadata whose configuration includes the other adapters.) field */
type BasinNode struct {
	Host         host.Host
	ReadRequests map[string]*ReadReqAnchor
	Did          string
	PrivKey      ed25519.PrivateKey
	Http         string
}

const ProtocolReadReq = "/basin/readreq/1.0.0"
const ProtocolReadRes = "/basin/readres/1.0.0"
const ProtocolSubReq = "/basin/subreq/1.0.0"
const ProtocolSubRes = "/basin/subres/1.0.0"

// TODO: ProtocolWriteReq/Res and associated handlers

var (
	TheBasinNode *BasinNode // Is this sus?
)

type BasinNodeConfig struct {
	Http string
	Did  string
	Pw   string
}

func (c *BasinNodeConfig) SetDefaults() {
	if c.Http == "" {
		c.Http = "http://127.0.0.1:8555"
	}
}

func StartBasinNode(config BasinNodeConfig) (BasinNode, error) {
	// Create listener on port
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))

	basin := BasinNode{Host: h, ReadRequests: map[string]*ReadReqAnchor{}, Http: config.Http}
	if err != nil {
		return basin, err
	}

	h.SetStreamHandler(ProtocolReadReq, basin.readReqHandler)
	h.SetStreamHandler(ProtocolReadRes, basin.readResHandler)
	h.SetStreamHandler(ProtocolSubRes, basin.subResHandler)
	h.SetStreamHandler(ProtocolSubReq, basin.subReqHandler)

	basin.LoadPrivateKey(config.Did, config.Pw)

	// Initialize necessary files TODO: Some more though here
	sourcesUrl := GetUserDataUrl(basin.Did, "producer.sources")
	sources, err := json.Marshal([]string{"basin://test"})
	if err != nil {
		log.Fatal("Couldn't initialize files: " + err.Error())
	}
	err = LocalOnlyDb.Write(sourcesUrl, sources)
	if err != nil {
		log.Fatal("Failed to write sources file: " + err.Error())
	}

	TheBasinNode = &basin
	return basin, nil
}

/* Handle a subscription request. */
func (b *BasinNode) HandleSubscriptionRequest(ctx context.Context, did string, permissions *[]PermissionJson) error {
	// TODO: Custom rules for accepting subscription requests so it can be automated
	url := GetUserDataUrl(b.Did, "producer.requests")
	requests, err := b.GetRequests(ctx, "producer")
	if err != nil {
		log.Println("Failed to read producer.requests: ", err.Error())
		return err
	}
	*requests = append(*requests, (*permissions)...)

	data, err := json.Marshal(requests)
	if err != nil {
		log.Println("Failed to marshal requests: ", err.Error())
		return err
	}
	err = b.WriteResource(ctx, url, data)
	if err != nil {
		return err
	}
	return nil
}

/* Sets the given DID and private key to be the current signer for the node */
// TODO: We eventually want the node to be multi-tenant
func (b *BasinNode) LoadPrivateKey(did string, pw string) error {
	priv, err := didutil.ReadKeystore(did, pw)
	if err != nil {
		return err
	}
	b.PrivKey = priv
	b.Did = did
	return nil
}

/* The uniform interface for retrieving any Basin resource, local or remote */
func (b *BasinNode) ReadResource(ctx context.Context, url string) ([]byte, error) {
	// Get list of sources on this node (can't call GetSources for infinite loop)
	// TODO: Should be using something more efficient so we don't have to search over whole array
	srcsUrl := GetUserDataUrl(b.Did, "producer.sources")
	data, err := LocalOnlyDb.Read(srcsUrl)
	if err != nil {
		log.Println("Failed to read sources from LocalOnlyDb: ", err.Error())
		return nil, err
	}
	srcs := Unmarshal[[]string](data)

	if Contains(*srcs, url) {
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
		req := &pb.ReadRequest{Url: url, MessageData: b.newMessageData(uuid.New().String())}
		sig, err := b.signProtoMsg(req)
		if err != nil {
			log.Println(err)
		}
		req.MessageData.Sig = sig

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

/* The uniform interface for writing to any Basin resource, local or remote */
func (b *BasinNode) WriteResource(ctx context.Context, url string, value []byte) error {
	return adapters.MainAdapter.Write(url, value)
}

// Working on making the metadata appear...
func (b *BasinNode) Register(ctx context.Context, url string, adapter client.AdapterJson, permissions []client.PermissionJson, schema map[string]interface{}) error {
	// A couple of todos for later...
	// 1. TODO: Make sure did owns the domain
	// 2. TODO: Check whether a schema already exists at this domain. If so, version it.
	// For now we'll assume that the URL by itself returns newest version, but later this might have to be
	// done more explicity. Consider how one might request an older version. Is this a header, part of the path or query?

	// TODO: Note that right here we just loaded a file from the filesystem and threw it into LevelDB
	// This is when we want to start storing things as actual files? Just start thinking about it.

	// Run all the file writes in parallel

	g, ctx := errgroup.WithContext(ctx)

	// SCHEMA
	schemaUrl := GetMetadataUrl(url, Schema)
	g.Go(func() error {
		schemaRaw, err := json.Marshal(schema)
		if err != nil {
			log.Println("Error marshalling schema file: " + err.Error())
			return err
		}
		return b.WriteResource(ctx, schemaUrl, schemaRaw)
	})

	// PERMISSIONS
	permUrl := GetMetadataUrl(url, Permissions)
	g.Go(func() error {
		permRaw, err := json.Marshal(permissions)
		if err != nil {
			log.Println("Error marshalling permissions file: " + err.Error())
			return err
		}
		return b.WriteResource(ctx, permUrl, permRaw)
	})

	// ADAPTER CONFIG
	adpUrl := GetMetadataUrl(url, Adapter)
	g.Go(func() error {
		// TODO: Again you're causing problems not having a source of truth for type generation :((((
		adpRaw, err := json.Marshal(adapter)
		if err != nil {
			log.Println("Error marshalling adapter file: " + err.Error())
			return err
		}
		return b.WriteResource(ctx, adpUrl, adpRaw)
	})

	// SOURCES
	sourcesUrl := GetUserDataUrl(b.Did, "producer.sources")
	g.Go(func() error {
		currSrcs, err := LocalOnlyDb.Read(sourcesUrl)
		var srcs []string
		err = json.Unmarshal(currSrcs, &srcs)
		if err != nil {
			log.Println("Error parsing sources file: " + err.Error())
			return err
		}
		srcs = append(srcs, url)
		finalSrcs, err := json.Marshal(srcs)
		if err != nil {
			log.Println("Error marshalling sources: " + err.Error())
			return err
		}
		// TODO: Should this be written to the LocalAdapter instead of LocalOnlyDb??
		// return b.WriteResource(ctx, sourcesUrl, finalSrcs)
		return LocalOnlyDb.Write(sourcesUrl, finalSrcs)
	})

	if err := g.Wait(); err != nil {
		log.Println("Error writing to one of the files: " + err.Error())
		return err
	}

	// Register with the routing table
	err := HostRouter.RegisterUrl(ctx, url)
	if err != nil {
		log.Println("Error regstering URL to Kademlia DHT: " + err.Error())
		return err
	}

	// Just like any other update - should tell subscribers (want a function for this)

	log.Println("Successfully registered resource at " + url)
	return nil
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
	url := GetUserDataUrl(b.Did, mode+".sources")
	val, err := b.ReadResource(ctx, url)
	if err != nil {
		return nil, err
	}

	return Unmarshal[[]string](val), nil
}

func (b *BasinNode) GetRequests(ctx context.Context, mode string) (*[]PermissionJson, error) {
	url := GetUserDataUrl(b.Did, mode+".requests")
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
