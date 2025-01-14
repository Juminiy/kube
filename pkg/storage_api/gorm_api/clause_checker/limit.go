package clause_checker

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// LimitClause
// LIMIT and OFFSET
func (cfg *Config) LimitClause(tx *gorm.DB) {

}

func LimitClause(tx *gorm.DB) (limitClause clause.Limit, ok bool) {
	limit, ook := util.MapElemOk(tx.Statement.Clauses, Limit)
	if !ook {
		return
	}
	if limitClause, ok = limit.Expression.(clause.Limit); ok {
		ok = limitClause.Limit != nil && *limitClause.Limit > 0
	}
	return

}
