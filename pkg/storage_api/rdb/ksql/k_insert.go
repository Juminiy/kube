package ksql

import "database/sql"

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

func (i Insert) Build() string {
	return ""
}
