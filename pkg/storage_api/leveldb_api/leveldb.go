package leveldb_api

import (
	"github.com/syndtr/goleveldb/leveldb"
)

func New(fd string) (*leveldb.DB, error) {
	return leveldb.OpenFile(fd, nil)
}
