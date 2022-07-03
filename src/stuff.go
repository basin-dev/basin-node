package main

import (
	"encoding/json"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

// This is what is run when notification is recieved of fresh data from the source.
// Consequences are:
// - Notification of subscribers
// - Update to freshness record
func logFreshData(urls []string) {

}

type ErrorHandleAction int64

const (
	Panic ErrorHandleAction = iota
	LogFatal
)

func handleErr(err error, action ErrorHandleAction) {
	if err != nil {
		switch action {
		case Panic:
			panic(err)
		case LogFatal:
			log.Fatal(err)
		}
	}
}

func parseUrl(url string) UrlJson {

	urlJson := new(UrlJson)

	return urlJson
}

func printUrl(url UrlJson) string {
	return url.scheme + "://" + url.user + "." + url.domain
}

func getMetadataUrl(dataUrl string, prefix string) string {
	parsed := parseUrl(dataUrl)
	url.domain = "meta." + prefix + "." + url.domain
}

func contains(slice []interface{}, val interface{}) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}
	return false
}

func expired(date string) bool {
	// TODO (date types in go?)
	return false
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

	val, err := db.Get([]byte(url), nil)
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

	err := db.Put([]byte(url), value, nil)

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

func registerSchema(url string, schema []byte, did string, db *leveldb.DB) { // TODO: Should both db and did be made a part of the ctx, or of some other shared object passed everywhere?
	schemaUrl := getMetadataUrl(url, "schema")

	// 1. TODO: Make sure did owns the domain

	// 2. Check whether a schema already exists at this URL. If so, version it.
	// For now we'll assume that the URL by itself returns newest version, but later this might have to be
	// done more explicity. Consider how one might request an older version. Is this a header, part of the path or query?
	err := db.Put([]byte(schemaUrl), schema, nil)

	if err != nil {
		handleErr(err, LogFatal)
	}
}

func validateDataToSchema(data []byte, schema []byte) bool {
	return false
}

type Notification struct {
	Url      string
	Reason   string
	SenderID string
}

// Do we want some append only record of when updates happen?
// When new data, send a message to all subscribers of the URL of the form described above
func registerDataUpdate(url string, db *leveldb.DB) {

	ntf := Notification{
		Url: url,
		Reason: "update",
		date: Date(),
		SenderID: did??Node ID?,
	}

	sendMessage(ctx, url)
}

// Every node maintains the permissions for that user.
// Don't want to store in complete key/value format, because that takes up too much space.
// What happens when people have so many different types of data that even their identifiers and metadata take up too much space for one node to handle?

// Better to organize around access lists?
// e.g. resource -> action -> did
// Add, modify, check, validate

// Does it make sense for basin urls to match entirely with https urls?
// So...
