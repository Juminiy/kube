package clause_checker

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
)

// OrderByClause
// ORDER BY column or ORDER BY columnList
func (cfg *Config) OrderByClause(tx *gorm.DB) {
	txClause, ok := OrderByClause(tx)
	if !ok {
		return
	}

	columns := make([]clause.OrderByColumn, 0, len(txClause.Columns))
	slices.All(txClause.Columns)(func(_ int, column clause.OrderByColumn) bool {
		if len(column.Column.Name) > 0 {
			columns = append(columns, column)
		}
		return true
	})
	orderClause := tx.Statement.Clauses[OrderBy]
	txClause.Columns = columns
	orderClause.Expression = txClause
	tx.Statement.Clauses[OrderBy] = orderClause
}

func OrderByClause(tx *gorm.DB) (orderByClause clause.OrderBy, ok bool) {
	orderBy, ook := util.MapElemOk(tx.Statement.Clauses, OrderBy)
	if !ook {
		return
	}
	if orderByClause, ok = orderBy.Expression.(clause.OrderBy); ok {
		ok = len(orderByClause.Columns) > 0
	}
	return
}

// tmp not to do so
func omitOrderByNotKnownColumn(tx *gorm.DB) {}
