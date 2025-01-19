package clause_checker

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	return !ok
}

func (cfg *Config) RowRawClause(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}

	if cfg.AllowWrapRawOrRowByClause {
		// TODO: clause Builder and clause Checker
	}
}
