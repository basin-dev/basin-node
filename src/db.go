package main

import (
	"log"

	leveldb "github.com/syndtr/goleveldb/leveldb"
)

func startDB() *leveldb.DB {
	db, err := leveldb.OpenFile("/tmp/db", nil)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
