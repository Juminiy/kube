package korm

import (
	"database/sql"
	"golang.org/x/net/context"
	"log"
)

type session struct {
	DB  *sql.DB
	C   *Config     // C is short for Config, config
	L   *log.Logger // L is short for log.Logger, logger
	ctx context.Context
}

func (tx *session) Create(v any) (err error) {
	return tx.NewInsert(v).Execute().All()
}

func (tx *session) Delete(v any) error {
	return nil
}

func (tx *session) Update(v any) error {
	return nil
}

func (tx *session) Query(v any) error {
	return nil
}
