package adapters

import (
	"fmt"

	"github.com/sestinj/basin-node/log"
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
		return nil, fmt.Errorf("Error getting key '%s' from leveldb: %w\n", url, err)
	}

	return val, nil
}

func (l LevelDbLocalAdapter) Write(url string, value []byte) error {
	err := l.db.Put([]byte(url), value, nil)
	if err != nil {
		log.Error.Println(err)
	}

	return err
}

func StartDB(path string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile("/tmp/"+path, nil)
	if err != nil {
		log.Error.Fatal("Failed to open LevelDB: " + err.Error() + ".\n Might solve this by deleting the /tmp/db folder if you are fine clearing the database.")
	}

	LocalAdapter = LevelDbLocalAdapter{db}

	return db, err
}
