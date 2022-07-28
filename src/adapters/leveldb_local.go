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

func (l LevelDbLocalAdapter) Read(url string) chan ReadPromise {
	ch := make(chan ReadPromise)

	go func() {
		defer close(ch)

		val, err := l.db.Get([]byte(url), nil)
		if err != nil {
			log.Println(err)
			ch <- ReadPromise{Data: val, Err: err}
			return
		}

		ch <- ReadPromise{Data: val, Err: nil}
	}()

	return ch
}

func (l LevelDbLocalAdapter) Write(url string, value []byte) chan error {
	ch := make(chan error)

	go func() {
		defer close(ch)

		err := l.db.Put([]byte(url), value, nil)
		if err != nil {
			log.Println(err)
		}
		ch <- err
	}()

	return ch
}

func StartDB() (*leveldb.DB, error) {
	db, err := leveldb.OpenFile("/tmp/db", nil)
	if err != nil {
		log.Fatal(err)
	}

	LocalAdapter = LevelDbLocalAdapter{db}

	return db, err
}
