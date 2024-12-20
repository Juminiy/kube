package korm

import (
	"database/sql"
	"github.com/pkg/errors"
)

func Open() (I, error) {
	db, err := sql.Open("sqlite3", "kdb.db")
	if err != nil {
		return nil, errors.Wrap(err, "open database/sql driver sqlite3")
	}
	return &session{DB: db}, nil
}

func (tx *session) Close() error {
	return tx.DB.Close()
}
