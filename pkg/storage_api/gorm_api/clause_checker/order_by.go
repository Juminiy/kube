package clause_checker

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// OrderByClause
// ORDER BY column or ORDER BY columnList
func (cfg *Config) OrderByClause(tx *gorm.DB) {

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
