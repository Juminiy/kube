package korm

import "database/sql"

// create table if not exists `tbl_korm`(`id`, `name`, `desc`, `extras`)
// drop table if exists `tbl_korm`;

type Table struct {
}

type Column struct {
	Type string
	sql.ColumnType
}

type Schema interface{ TName() string }
