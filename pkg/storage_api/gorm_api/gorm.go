package gorm_api

import (
	"database/sql"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
	db *sql.DB
}

func New(cfg gorm.Config) (*DB, error) {
	tx, err := gorm.Open(cfg.Dialector, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "gorm open <- open by config and dialector")
	}
	sqldb, err := tx.DB()
	if err != nil {
		return nil, errors.Wrap(err, "gorm open <- get *sql.DB")
	}
	return &DB{
		DB: tx,
		db: sqldb,
	}, nil
}

func (db *DB) Close() error {
	return errors.Wrap(db.db.Close(), "gorm close <- *sql.DB.Close")
}

func (db *DB) Default() *DB {
	db.db.SetConnMaxLifetime(util.TimeSecond(8))
	db.db.SetConnMaxIdleTime(util.TimeSecond(8))
	db.db.SetMaxIdleConns(8)
	db.db.SetMaxOpenConns(8)
	return db
}
