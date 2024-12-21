package korm

import (
	"database/sql"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/pkg/errors"
	"log"
)

func Open() (I, error) {
	db, err := sql.Open(_DefaultDriver, _DefaultDB)
	if err != nil {
		return nil, driverErrWrap(err, "open")
	}

	ctx := util.TODOContext()
	cfg := _DefaultConfig

	db.SetMaxIdleConns(cfg.ConnC.IdleConns)
	db.SetMaxOpenConns(cfg.ConnC.OpenConns)
	db.SetConnMaxLifetime(util.TimeSecond(cfg.ConnC.LifeSec))
	db.SetConnMaxIdleTime(util.TimeSecond(cfg.ConnC.IdleSec))

	err = db.PingContext(ctx)
	if err != nil {
		return nil, driverErrWrap(err, "ping")
	}

	return &session{
		DB:  db,
		C:   cfg,
		L:   log.Default(),
		ctx: ctx,
	}, nil
}

func (tx *session) Close() error {
	return tx.DB.Close()
}

func driverErrWrap(err error, do string) error {
	return errors.Wrapf(err, "scope[database/sql], do[%s], driver[%s]", do, _DefaultDriver)
}
