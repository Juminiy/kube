package korm

import "database/sql"

type Table struct {
	TName        func() string
	WithoutRowID bool
}

type Column struct {
	Type string
	sql.ColumnType
}

type Schema interface{ TName() string }

const _DefaultDriver = `sqlite3`
const _DefaultDB = `kdb.db`
const _ROWID = `rowid`
const _OID = `oid`
const _ROWID_ = `_rowid_`