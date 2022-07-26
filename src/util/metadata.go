package util

import (
	"errors"
	"log"

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

func RequestSubscription(url string) error {
	return nil
}
