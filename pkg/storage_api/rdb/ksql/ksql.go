package ksql

import "database/sql"

type B interface {
	Build() string
}

type builder struct{}

func (b *builder) Build() string {
	return ""
}

type BraceF func(string) string

var accent = func(s string) string {
	return "`" + s + "`"
}

var sQuote = func(s string) string {
	return "'" + s + "'"
}

var quote = func(s string) string {
	return `"` + s + `"`
}

var percent = func(s string) string {
	return "%" + s + "%"
}

var bracket = func(s string) string {
	return "(" + s + ")"
}

// insert into tbl_korm(id, name, extras) values (1, 'frame', '{"cve": "none", "r_list":[1]}')
type Insert struct {
	BraceTable  BraceF
	BraceColumn BraceF
	BraceText   BraceF
	CT          []sql.ColumnType
	Values      [][]any
}

func (i *Insert) set() *Insert {
	if i.BraceTable == nil {
		i.BraceTable = accent
	}
	if i.BraceColumn == nil {
		i.BraceColumn = accent
	}
	if i.BraceText == nil {
		i.BraceText = sQuote
	}
	return i
}

func (i *Insert) Build() string {
	return ""
}
