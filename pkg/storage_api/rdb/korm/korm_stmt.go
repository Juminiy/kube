package korm

import (
	"database/sql"
	"errors"
	"github.com/Juminiy/kube/pkg/storage_api/rdb/kinternal"
	"github.com/Juminiy/kube/pkg/storage_api/rdb/ksql"
	"github.com/Juminiy/kube/pkg/util"
	"golang.org/x/exp/maps"
	"reflect"
	"slices"
	"time"
)

type Stmt struct {
	*session
	Tv _Tv // Tv is short for safe_reflectv3.Tv
	//ETv _ETv // ETv is short for safe_reflectv3.ETv

	FieldColumn map[string]string
	ColumnField map[string]string
	PK          PrimaryKey
	Table       string
	Column      []string
	Values      [][]any

	B kinternal.StringBuilder
	*util.ErrHandle

	RowAffect int64
}

type PrimaryKey struct {
	Field    string
	Column   string
	scanType reflect.Type
}

var ErrNotSchema = errors.New("v is not a schema")

func (tx *session) NewInsert(v any) *Stmt {
	return (&Stmt{session: tx}).Parse(v)
}

func (stmt *Stmt) Parse(v any) *Stmt {
	stmt.ErrHandle = util.NewErrHandle()
	stmt.Tv = _Ind(v)
	//stmt.ETv = stmt.Tv.ETv().
	//	OmitReadByTag2(stmt.C.Tag.Key, stmt.C.Tag.Column, _ROWID).
	//	OmitReadByTag2(stmt.C.Tag.Key, stmt.C.Tag.Column, _ROWID_).
	//	OmitReadByTag2(stmt.C.Tag.Key, stmt.C.Tag.Column, _OID)

	return stmt.parseTable().parsePrimaryKeyField().parseValues()
}

func (stmt *Stmt) parseTable() *Stmt {
	rets, called := stmt.Tv.CallMethod(_SchemaTable, nil)
	if called && len(rets) > 0 {
		if table, ok := rets[0].(string); ok {
			stmt.Table = table
		}
	}
	if len(stmt.Table) == 0 {
		stmt.Has(ErrNotSchema)
	}
	stmt.FieldColumn = stmt.Tv.Tag2(stmt.C.Tag.Key, stmt.C.Tag.Column)
	stmt.ColumnField = util.MapVK(stmt.FieldColumn)
	return stmt
}

func (stmt *Stmt) parsePrimaryKeyField() *Stmt {
	for _, col := range _ROWID_LIST {
		field, ok := util.MapElemOk(stmt.ColumnField, col)
		if ok {
			stmt.PK = PrimaryKey{
				Field:    field,
				Column:   col,
				scanType: util.MapElem(stmt.Tv.FieldType(), field),
			}
			return stmt
		}
	}
	return stmt
}

func (stmt *Stmt) parseValues() *Stmt {
	// map is unordered
	/*stmt.Column = lo.MapToSlice(rv.Tag2(stmt.C.Tag.Key, stmt.C.Tag.Column),
		func(name string, column string) string {
			return column
		})
	stmt.Values = lo.Map(rv.Values(), func(rowMap map[string]any, index int) []any {
		return lo.MapToSlice(rowMap, func(name string, val any) any {
			return val
		})
	})*/

	fieldCol := util.MapDeleteR(maps.Clone(stmt.FieldColumn), stmt.PK.Field)
	fieldVals := stmt.Tv.Values()
	fields := maps.Keys(fieldCol)

	column := make([]string, 0, len(fieldCol))
	slices.All(fields)(func(_ int, field string) bool {
		column = append(column, fieldCol[field])
		return true
	})
	stmt.Column = column

	values := make([][]any, 0, len(fieldVals))
	for _, val := range fieldVals {
		value := make([]any, 0, len(fieldCol))
		slices.All(fields)(func(_ int, field string) bool {
			value = append(value, val[field])
			return true
		})
		values = append(values, value)
	}
	stmt.Values = values

	return stmt
}

func (stmt *Stmt) Execute() *Stmt {
	stmt.B = &ksql.Insert{
		Table:  stmt.Table,
		Column: stmt.Column,
		Values: stmt.Values,
	}

	stmt.Has(stmt.trace(func() error {
		result, err := stmt.DB.ExecContext(stmt.ctx, stmt.B.Build())
		if err != nil {
			return err
		}
		return stmt.AfterCreate(result).All()
	}))
	return stmt
}

func (stmt *Stmt) AfterCreate(result sql.Result) *Stmt {
	lastRowID, idErr := result.LastInsertId()
	if stmt.Has(idErr) {
		return stmt
	}
	vsz := int64(len(stmt.Values))
	for idx := int64(0); idx < vsz; idx++ {
		stmt.Tv.SetField(map[string]any{stmt.PK.Field: lastRowID - vsz + idx + 1}, int(idx))
	}

	rowAffect, rAErr := result.RowsAffected()
	if stmt.Has(rAErr) {
		return stmt
	}
	stmt.RowAffect = rowAffect
	return stmt
}

func (stmt *Stmt) trace(f util.Func) error {
	timeBegin := time.Now()
	err := f()
	stmt.L.Printf("time: [%s] row_affect: [%d] sql: [%s]\n",
		util.HumanTimeDesc(time.Now().Sub(timeBegin)),
		stmt.RowAffect,
		stmt.B.Build(),
	)

	return err
}

func (stmt *Stmt) retry(f util.Func) error {
	return nil
}
