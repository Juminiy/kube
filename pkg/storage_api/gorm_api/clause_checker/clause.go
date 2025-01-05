package clause_checker

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	Where      = "WHERE"
	Returning  = "RETURNING"
	OnConflict = "ON CONFLICT"
	From       = "FROM"
	Set        = "SET"
	Select     = "SELECT"
	Limit      = "LIMIT"
	OrderBy    = "ORDER BY"
	GroupBy    = "GROUP BY"
)

func WhereClause(tx *gorm.DB) (whereClause clause.Where, ok bool) {
	where, whereKey := util.MapElemOk(tx.Statement.Clauses, Where)
	if !whereKey {
		return
	}

	if where.Expression == nil {
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

func OmitNullWhereClause(tx *gorm.DB) bool {
	return !NoWhereClause(tx)
}
