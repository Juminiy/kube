package korm

import (
	"database/sql"
	"github.com/Juminiy/kube/pkg/storage_api/rdb/ksql"
	"github.com/Juminiy/kube/pkg/util"
	"golang.org/x/net/context"
	"log"
)

type session struct {
	DB  *sql.DB
	C   *Config     // C is short for Config, config
	L   *log.Logger // L is short for log.Logger, logger
	B   ksql.B      // B is short for ksql.B, builder
	ctx context.Context
}

func (tx *session) Create(v any) (err error) {
	return tx.trace(func() error {
		tx.B = ksql.Insert{
			CT:     nil,
			Values: nil,
		}
		r, err := tx.DB.ExecContext(tx.ctx, tx.B.Build())
		if err != nil {
			return err
		}

		if _, err := r.LastInsertId(); err != nil {
			return err
		}

		return nil
	})
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

func (tx *session) reset() *session {
	tx.B = nil
	return tx
}

func (tx *session) trace(f util.Func) error {
	return nil
}

func (tx *session) retry(f util.Func) error {
	return nil
}
