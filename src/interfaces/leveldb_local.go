package interfaces

import (
	leveldb "github.com/syndtr/goleveldb/leveldb"
)

func read(url string, db *leveldb.DB) []byte {
	val, err := db.Get([]byte(url), nil)
	handleErr(err, LogFatal)

	return val
}

func write(url string, val []byte, db *leveldb.DB) bool {
	err := db.Put([]byte(url), val, nil)
	handleErr(err, LogFatal)

	return true
}