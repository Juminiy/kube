package clause_checker

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"slices"
)

// WhereClause
// Expr or ExprList
func (cfg *Config) WhereClause(tx *gorm.DB) {
	txClause, ok := WhereClause(tx)
	if !ok {
		return
	}

	exprIList := make([]clause.Expression, 0, len(txClause.Exprs))
	slices.All(txClause.Exprs)(func(_ int, exprI clause.Expression) bool {
		if checkExprI(exprI) {
			exprIList = append(exprIList, exprI)
		}
		return true
	})
	whereClause := tx.Statement.Clauses[Where]
	txClause.Exprs = exprIList
	whereClause.Expression = txClause
	tx.Statement.Clauses[Where] = whereClause
}

func WhereClause(tx *gorm.DB) (whereClause clause.Where, ok bool) {
	where, wok := util.MapElemOk(tx.Statement.Clauses, Where)
	if !wok {
		return
	}
	if whereClause, ok = where.Expression.(clause.Where); ok {
		ok = len(whereClause.Exprs) > 0
	}
	return
}

func NoWhereClause(tx *gorm.DB) bool {
	_, ok := WhereClause(tx)
	return !ok && !destKindIsStructAndHasPrimaryKeyNotZero(tx.Statement)
}

// referred from: callbacks.BuildQuerySQL
// has at least one primaryKey value is not zero
func destKindIsStructAndHasPrimaryKeyNotZero(stmt *gorm.Statement) bool {
	if stmt.SQL.Len() == 0 {
		if stmt.ReflectValue.Kind() == reflect.Struct &&
			stmt.ReflectValue.Type() == stmt.Schema.ModelType {
			for _, primaryField := range stmt.Schema.PrimaryFields {
				if _, isZero := primaryField.ValueOf(stmt.Context, stmt.ReflectValue); !isZero {
					return true
				}
			}
		}
	}
	return false
}
