package util

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/xeipuuv/gojsonschema"
)

// This is what is run when notification is recieved of fresh data from the source.
// Consequences are:
// - Notification of subscribers
// - Update to freshness record
func logFreshData(urls []string) {
	return true
}

func contains(slice []interface{}, val interface{}) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}
	return false
}

func expired(date time.Time) bool {
	return date.After(date.Now())
}

func authorizeAction(url string, action string, did string, db *leveldb.DB) bool {
	// Would be useful to instead start a "session" that preloads all permissions of the user for quicker checking. Or just index better.

	// You could index this all much more efficiently, but for the sake of making everything a key value in the same way, what do you do?
	// Maybe this is where you build in custom indexing? What does that mean though? A separate domain for the different indexings?
	// Or a part of the scheme that specifies which indexing to use?
	permissions := getPermissions(url, db)
	for _, permission := range permissions {
		for _, capability := range permission.capabilities {
			if capability.action == action && !expired(capability.expiration) {
				for _, entity := range permission.entities {
					if entity == did {
						return true
					}
				}
			}
		}
	}
	return false
}

func getPermissions(url string, db *leveldb.DB) *[]PermissionJson {
	data, err := db.Get([]byte(url), nil)
	if err != nil {
		handleErr(err, LogFatal)
	}

	permissions := new([]PermissionJson)
	json.Unmarshal(data, permissions)

	return permissions
}

func readData(url string, did string, db *leveldb.DB) []byte {
	if !authorizeAction(url, "read", did, db) {
		return false
	}

	val, err := db.Get([]byte(url), nil) // TODO - Call the HTTP URL that implements the Basin Interface
	if err != nil {
		return false // TODO: Need some way to propogate errors
	}

	return val
}

func writeData(url string, value []byte, did string, db *leveldb.DB) bool {
	if !authorizeAction(url, "write", did, db) {
		return false
	}

	// TODO: Validate with schema

	err := db.Put([]byte(url), value, nil) // TODO: Again, call HTTP URL, Basin Interface, blah, blah

	if err != nil {
		return false
	}

	return true
}

func addPermission(permission PermissionJson, db *leveldb.DB) {
	batch := new(leveldb.Batch)
	for _, url := range permission.Data {
		key := []byte(url + "::permissions")

		has, err := db.Has(key, nil)
		handleErr(err, Panic)

		if has {
			curr, err := db.Get(key, nil)
			handleErr(err, Panic)

			currJson := new([]PermissionJson)
			err = json.Unmarshal(curr, currJson)
			handleErr(err, Panic)

			*currJson = append(*currJson, permission)

			data, err := json.Marshal(currJson)
			handleErr(err, Panic)

			batch.Put(key, data)
		} else {
			permissions := [1]PermissionJson{permission}
			data, err := json.Marshal(permissions)
			handleErr(err, Panic)

			batch.Put(key, data)
		}
	}
	err := db.Write(batch, nil)
	handleErr(err, Panic)
}

// TODO
func registerSchema(url string, schema []byte, did string, db *leveldb.DB) { // TODO: Should both db and did be made a part of the ctx, or of some other shared object passed everywhere?
	schemaUrl := getMetadataUrl(url, "schema")

	// 1. TODO: Make sure did owns the domain

	// 2. Check whether a schema already exists at this domain. If so, version it.
	// For now we'll assume that the URL by itself returns newest version, but later this might have to be
	// done more explicity. Consider how one might request an older version. Is this a header, part of the path or query?
	err := db.Put([]byte(schemaUrl), schema, nil)

	if err != nil {
		handleErr(err, LogFatal)
	}
}

func validateDataToSchema(data []byte, dataUrl string, db *leveldb.DB) (bool, []gojsonschema.ResultError) {
	schemaUrl := getMetadataUrl(dataUrl, "schema")
	schema, err := db.Get([]byte(schemaUrl), nil)
	handleErr(err, LogFatal)

	schemaLoader := gojsonschema.NewBytesLoader(schema)
	dataLoader := gojsonschema.NewBytesLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	handleErr(err, LogFatal)

	return result.Valid(), result.Errors()
}

type Notification struct {
	Url      string
	Reason   string
	SenderID string
}

// TODO
// Do we want some append only record of when updates happen?
// When new data, send a message to all subscribers of the URL of the form described above
func registerDataUpdate(url string, db *leveldb.DB) {

	urlJson := parseUrl(url)

	ntf := Notification{
		Url: url,
		Reason: "update",
		date: time.Now(),
		SenderID: urlJson.user
	}

	sendMessage(ctx, url)
}

type Request struct {
	SenderID string
	Url string
	Action string // Read, write, or request capabilities
	Capabilities? []CapabilityJson
	// I'm pretty sure HTTP deals with the whole concept of request IDs, so why not use it?
}

// TODO - is this supposed to be subscribe function?
func requestResource(url string) []byte {
	// Since it will be stored on the source node, need to find that node and request
	// Find via IP/DNS? But can URL Schemes other than http be resolved?
	// Does libp2p offer some obvious solution to this?
	// Would make sense though for entities to run their nodes as servers at a certain URL...they shouldn't have to own one though.
	// nvm...you want to route to the username part of the domain

	parsedUrl := parseUrl(url)
	user = parsedUrl.user
	// Then does user need to be resolved to DID?
	// Then resolve the DID to an IP address.
	// Then make the request to the IP address.
	// Types of requests to a resource might include:
	// - Get the value
	// - Obtain/purchase a capability (this is the same as just writing to the permissions resource, but this has to be a special case)
	// - More generally, exercise one of these capabilities (read, write)
}

// Every node maintains the permissions for that user.
// Don't want to store in complete key/value format, because that takes up too much space.
// What happens when people have so many different types of data that even their identifiers and metadata take up too much space for one node to handle?

// Better to organize around access lists?
// e.g. resource -> action -> did
// Add, modify, check, validate

// Does it make sense for basin urls to match entirely with https urls?
// So...

func getSchema(dataUrl string, db *leveldb.DB) *[]SchemaJson {
	url := getMetadataUrl(dataUrl, "schema")
	data, err := db.Get([]byte(url), nil)
	handleErr(err, LogFatal)

	schema := new([]SchemaJson)
	json.Unmarshal(data, schema)

	return schema
}

func getSources(mode string, db *leveldb.DB) *[]string {
	walletInfo := getWalletInfo(db)
	url := getUserDataUrl(walletInfo.did, mode + ".urls") // OR could do "urls.<mode>". Should probably match the CLI's order
	data, err := db.Get([]byte(url), nil)
	handleErr(err, LogFatal)

	sources := new([]string)
	json.Unmarshal(data, sources)

	return sources
}

// I need to figure out a way to access the leveldb.DB without passing it around, but for now it's in every function call...

func getReputation(mode string, db *leveldb.DB) *ReputationJson {
	walletInfo := getWalletInfo(db)
	url := getUserDataUrl(walletInfo.did, mode + ".reputation")
	data, err := db.Get([]byte(url), nil)
	handleErr(err, LogFatal)

	reputation := new(ReputationJson)
	json.Unmarshal(data, reputation)

	return reputation
}

func getWalletInfo(db *leveldb.DB) *WalletInfoJson {
	// "Local data" will almost certainly have to be a concept. There's no reason another node should be able to query this info.
	// It makes sense to still store in leveldb though, so separating by using a different scheme: "local://"
	data, err := db.Get([]byte("local://wallet"), nil)
	handleErr(err, LogFatal)

	walletInfo := new([]WalletInfoJson)
	json.Unmarshal(data, walletInfo)

	return walletInfo
}

func getRequests() {
	// I would hope this isn't necessary.
	// Like HTTP, all requests either happen right away or are failed, or is there something you're seeing that makes this necessary?
}

func getRoyalties(db *leveldb.DB) *RoyaltiesJson {
	walletInfo := getWalletInfo(db)
	url := getUserDataUrl(walletInfo.did, mode + ".royalties")
	data, err := db.Get([]byte(url), nil)
	handleErr(err, LogFatal)

	royalties := new(RoyaltiesJson)
	json.Unmarshal(data, royalties)

	return royalties
}

func getSchemas(mode string, db *leveldb.DB) *[]SchemaJson {
	basinUrls := getBasinUrls(mode, db)

	var schemas []SchemaJson

	for _, basinUrl := range basinUrls {
		schema := getSchema(basinUrl, db)
		append(schemas, schema)
	}

	return *schemas
}

func getCacheExpectations(mode string, db *leveldb.DB) *[]CacheExpectationJson {
	walletInfo := getWalletInfo(db)
	url := getUserDataUrl(walletInfo.did, mode + ".royalties")
	data, err := db.Get([]byte(url), nil)
	handleErr(err, LogFatal)

	royalties := new(RoyaltiesJson)
	json.Unmarshal(data, royalties)

	return royalties
}

