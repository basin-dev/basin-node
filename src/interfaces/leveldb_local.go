package interfaces

import (
	. "github.com/sestinj/basin-node/util"
	leveldb "github.com/syndtr/goleveldb/leveldb"
)

func read(url string, db *leveldb.DB) []byte {
	val, err := db.Get([]byte(url), nil)
	HandleErr(err, LogFatal)

	return val
}

func write(url string, val []byte, db *leveldb.DB) bool {
	err := db.Put([]byte(url), val, nil)
	HandleErr(err, LogFatal)

	return true
}
