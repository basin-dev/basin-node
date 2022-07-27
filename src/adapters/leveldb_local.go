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

func (l LevelDbLocalAdapter) Read(url string) ([]byte, error) {

	val, err := l.db.Get([]byte(url), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return val, nil
}

func (l LevelDbLocalAdapter) Write(url string, val []byte) error {
	err := l.db.Put([]byte(url), val, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func StartDB() (*leveldb.DB, error) {
	db, err := leveldb.OpenFile("/tmp/db", nil)
	if err != nil {
		log.Fatal(err)
	}

	LocalAdapter = LevelDbLocalAdapter{db}

	return db, err
}
