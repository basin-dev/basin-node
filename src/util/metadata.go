package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	. "github.com/sestinj/basin-node/adapters"
	. "github.com/sestinj/basin-node/structs"
)

type MetadataPrefix int64

const (
	Permissions MetadataPrefix = iota
	Royalties
	Schema
	Manifest
)

func (m MetadataPrefix) String() string {
	switch m {
	case Permissions:
		return "permissions"
	case Royalties:
		return "royalties"
	case Schema:
		return "schema"
	case Manifest:
		return "manifest"
	default:
		return "unknown"
	}
}

// Here we read the metadata...but how does it appear? Work begins in the section below

func ReadResource(url string) ([]byte, error) {
	// if Contains(*GetSources("producer"), url) {
	// 	// Determine which adapter to use

	// 	// This would also probably be the place to implement different guarantees?

	// } else {
	// 	// Use DNS/DHT to route to the node that produces this basin url

	// 	// Call requestResource to the next hop
	// }
	return nil, errors.New("Not yet implemented")
}

func WriteResource(url string, value []byte) error {
	// Do the same thing as ReadResource, if it's a local resource, just use the local adapter. And for now mostly everything should be.
	return errors.New("Not yet implemented")
}

func GetWalletInfo() *WalletInfoJson {
	data, err := LocalOnlyDb.Read("wallet")
	if err != nil {
		log.Fatal(err)
	}

	return Unmarshal[WalletInfoJson](data)
}

func GetPermissions(dataUrl string) *[]PermissionJson {
	url := GetMetadataUrl(dataUrl, Permissions)
	mdata, err := ReadResource(url)
	if err != nil {
		log.Fatal(err)
	}

	return Unmarshal[[]PermissionJson](mdata)
}

func GetSchema(dataUrl string) *SchemaJson {
	url := GetMetadataUrl(dataUrl, Schema)

	mdata, err := ReadResource(url)
	if err != nil {
		log.Fatal(err)
	}

	return Unmarshal[SchemaJson](mdata)
}

func GetSources(mode string) *[]string {
	walletInfo := GetWalletInfo()

	url := GetUserDataUrl(walletInfo.Did, mode+".sources")
	mdata, err := ReadResource(url)
	if err != nil {
		log.Fatal(err)
	}

	return Unmarshal[[]string](mdata)
}

func GetRequests(mode string) *[]PermissionJson {
	walletInfo := GetWalletInfo()

	url := GetUserDataUrl(walletInfo.Did, mode+".requests")
	mdata, err := ReadResource(url)
	if err != nil {
		log.Fatal(err)
	}

	return Unmarshal[[]PermissionJson](mdata)
}

func GetSchemas(mode string) *[]SchemaJson {
	sources := GetSources(mode)

	var schemas []SchemaJson

	for _, source := range *sources {
		schema := GetSchema(source)
		schemas = append(schemas, *schema)
	}

	return &schemas
}

// Working on making the metadata appear...
func Register(manifestPath string) error {
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

	err = WriteResource(permUrl, permsRaw)
	if err != nil {
		return err
	}

	// SCHEMA
	schemaUrl := GetMetadataUrl(manifest.Url, Schema)
	schemaRaw, err := json.Marshal(manifest.Schema) // TODO: What is the shape of the schema?
	err = WriteResource(schemaUrl, schemaRaw)
	if err != nil {
		return err
	}

	// MANIFEST
	manifestUrl := GetMetadataUrl(manifest.Url, Manifest)
	// TODO: Note that right here we just loaded a file from the filesystem and threw it into LevelDB
	// This is when we want to start storing things as actual files? Just start thinking about it.
	err = WriteResource(manifestUrl, manifestRaw)
	if err != nil {
		return err
	}

	// SOURCES
	walletInfo := GetWalletInfo()
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
	err = WriteResource(sourcesUrl, finalSrcs)
	if err != nil {
		return err
	}

	// Just like any other update - should tell subscribers (want a function for this)

	return nil
}

func RequestSubscription(url string) error {
	return nil
}
