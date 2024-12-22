package ksql

import (
	"github.com/Juminiy/kube/pkg/util/zerobuf/bufwriter"
	"github.com/spf13/cast"
	"sync"
)

type Insert struct {
	BraceTable  Brace
	BraceColumn Brace
	BraceText   Brace
	BraceValue  []func(string) string
	Table       string
	Column      []string
	Values      [][]any
	sync.Once
	sqlBytes []byte
}

func (i *Insert) set() *Insert {
	if i.BraceTable.NotValid() {
		i.BraceTable = accent
	}
	if i.BraceColumn.NotValid() {
		i.BraceColumn = accent
	}
	if i.BraceText.NotValid() {
		i.BraceText = sQuote
	}
	i.BraceValue = make([]func(string) string, len(i.Column))
	for idx, val := range i.Values[0] {
		switch _SP(cast.ToString(val)).Category() {
		case Num:
			i.BraceValue[idx] = none.Do
		default:
			i.BraceValue[idx] = sQuote.Do
		}
	}
	for idx, col := range i.Column {
		i.Column[idx] = i.BraceColumn.Do(col)
	}
	return i
}

func (i *Insert) NotValid() bool {
	return len(i.Table) == 0 ||
		len(i.Column) == 0 ||
		len(i.Values) == 0 ||
		len(i.Column) != len(i.Values[0])
}

func (i *Insert) Build() string {
	if i.NotValid() {
		return ""
	}
	i.Do(func() {
		i.set()
		sqlBuf := bufwriter.New().
			Words("INSERT", "INTO", i.BraceTable.Do(i.Table)).
			Byte('(').WordsSep(",", i.Column...).Byte(')').
			Word("VALUES").
			Byte('(').WordsaSepBrace(",", i.BraceValue, i.Values[0]).Byte(')')

		for idx := 1; idx < len(i.Values); idx++ {
			sqlBuf.Byte(',').
				Byte('(').WordsaSepBrace(",", i.BraceValue, i.Values[idx]).Byte(')')
		}

		i.sqlBytes = _S2B(sqlBuf.String())
	})
	return _B2S(i.sqlBytes)
}
