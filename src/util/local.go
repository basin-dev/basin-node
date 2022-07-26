package util

import (
	"log"

	leveldb "github.com/syndtr/goleveldb/leveldb"
)

var (
	LocalOnlyDb LocalOnlyDataInterface
)

/*
An interface to the local LevelDB that automatically prefixes so that the namespace is guaranteed to be separated
*/
type LocalOnlyDataInterface struct {
	db        *leveldb.DB
	keyPrefix string
}

func (l LocalOnlyDataInterface) Read(key string) []byte {
	val, err := l.db.Get([]byte(l.keyPrefix+key), nil)
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func (l LocalOnlyDataInterface) Write(key string, val []byte) bool {
	err := l.db.Put([]byte(l.keyPrefix+key), val, nil)
	if err != nil {
		log.Fatal(err)
	}

	return true
}

func StartLocalOnlyDb(db *leveldb.DB, keyPrefix string) LocalOnlyDataInterface {
	LocalOnlyDb = LocalOnlyDataInterface{db, keyPrefix}
	return LocalOnlyDb
}
