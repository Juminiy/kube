package clause_checker

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
)

func (cfg *Config) WhereClause(tx *gorm.DB) {
	txClause, ok := WhereClause(tx)
	if !ok {
		return
	}
	slices.All(txClause.Exprs)(func(i int, exprI clause.Expression) bool {
		tx.Logger.Info(tx.Statement.Context, "where_clause[%d]:[%v]", i, exprI)
		return true
	})
}
