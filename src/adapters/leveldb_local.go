package adapters

import (
	"log"

	leveldb "github.com/syndtr/goleveldb/leveldb"
)

var (
	LocalAdapter LevelDbLocalAdapter
)

type LevelDbLocalAdapter struct {
	db *leveldb.DB
}

func (l LevelDbLocalAdapter) Read(url string) []byte {

	val, err := l.db.Get([]byte(url), nil)
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func (l LevelDbLocalAdapter) Write(url string, val []byte) bool {
	err := l.db.Put([]byte(url), val, nil)
	if err != nil {
		log.Fatal(err)
	}

	return true
}

func StartDB() {
	db, err := leveldb.OpenFile("/tmp/db", nil)
	if err != nil {
		log.Fatal(err)
	}

	LocalAdapter = LevelDbLocalAdapter{db}
}
