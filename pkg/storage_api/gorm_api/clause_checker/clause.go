package clause_checker

import (
	"github.com/Juminiy/kube/pkg/util"
	safe_reflectv3 "github.com/Juminiy/kube/pkg/util/safe_reflect/v3"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
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

func checkExprI(exprI clause.Expression) bool {
	switch exprV := exprI.(type) {
	case clause.Eq, clause.Neq, clause.Gt, clause.Gte, clause.Lt, clause.Lte:
		rv := _Dir(exprV).Value
		column := rv.FieldByName("Column")
		_ = rv.FieldByName("Value")
		if column.IsZero() {
			return true
		}
	case clause.Like:

	case clause.Expr:
		if len(exprV.SQL) == 0 ||
			strings.Count(exprV.SQL, "?") != len(exprV.Vars) {
			return true
		}
	}
	return true
}

func checkExprIList(exprIList []clause.Expression) bool {
	return lo.CountBy(exprIList, func(item clause.Expression) bool {
		return checkExprI(item)
	}) != len(exprIList)
}

func checkExprICombination(exprI clause.Expression) bool {
	switch exprI.(type) {
	case clause.OrConditions:

	case clause.NotConditions:

	default:

	}
	return true
}

var _Ind = safe_reflectv3.Indirect
var _Dir = safe_reflectv3.Direct

func AnySliceCountZero(v []any) int {
	return lo.CountBy(v, func(item any) bool {
		return util.AssertZero(item)
	})
}
