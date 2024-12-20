package korm

import (
	"database/sql"
	"github.com/Juminiy/kube/pkg/storage_api/rdb/ksql"
)

type session struct {
	*sql.DB
	ksql.B
}

// insert into `tbl_korm`(`id`,`name`) values (1, 'otel');
func (tx *session) Create(v any) error {
	return nil
}

func (tx *session) Delete(v any) error {
	return nil
}

func (tx *session) Update(v any) error {
	return nil
}

// select * from `tbl_korm`;
func (tx *session) Query(v any) error {
	return nil
}
