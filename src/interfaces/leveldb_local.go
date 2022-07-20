package interfaces

import (
	"context"

	. "github.com/sestinj/basin-node/util"
	leveldb "github.com/syndtr/goleveldb/leveldb"
)

var (
	localAdapter LocalAdapter
)

type LocalAdapter struct {
	db *leveldb.DB
}

func (l LocalAdapter) Read(url string, ctx context.Context) []byte {

	val, err := l.db.Get([]byte(url), nil)
	HandleErr(err, LogFatal)

	return val
}

func (l LocalAdapter) Write(url string, val []byte) bool {
	err := l.db.Put([]byte(url), val, nil)
	HandleErr(err, LogFatal)

	return true
}

func StartDB() {
	db, err := leveldb.OpenFile("/tmp/db", nil)
	HandleErr(err, LogFatal)

	localAdapter = LocalAdapter{db}
}
