package main

import "github.com/syndtr/goleveldb/leveldb"

// This is what is run when notification is recieved of fresh data from the source.
// Consequences are:
// - Notification of subscribers
// - Update to freshness record
func logFreshData(urls []string) {

}

func addPermission(permission Permission, db *leveldb.DB) {
	for _, url := range permission.Data {
		batch := new(leveldb.Batch)
		batch.Put([]byte(url))
	}
}

// Every node maintains the permissions for that user.
// Don't want to store in complete key/value format, because that takes up too much space.
// What happens when people have so many different types of data that even their identifiers and metadata take up too much space for one node to handle?
