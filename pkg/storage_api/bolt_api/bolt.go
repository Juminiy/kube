package bolt_api

import (
	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/boltdb/bolt"
)

func New(fd string) (*bolt.DB, error) {
	return bolt.Open(fd, internal_api.FilePerm, nil)
}
