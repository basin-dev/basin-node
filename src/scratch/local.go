package scratch

import (
	"fmt"

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

func (l LocalOnlyDataInterface) Read(key string) ([]byte, error) {
	val, err := l.db.Get([]byte(l.keyPrefix+key), nil)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (l LocalOnlyDataInterface) Write(key string, val []byte) error {
	err := l.db.Put([]byte(l.keyPrefix+key), val, nil)
	if err != nil {
		return fmt.Errorf("Error writing to local only db: %w\n", err)
	}

	return nil
}

func StartLocalOnlyDb(db *leveldb.DB, keyPrefix string) LocalOnlyDataInterface {
	LocalOnlyDb = LocalOnlyDataInterface{db, keyPrefix}
	return LocalOnlyDb
}
