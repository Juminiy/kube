package multi_tenants

import (
	"errors"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"github.com/Juminiy/kube/pkg/util"
	gormdrivermysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
)

var ErrDeleteTenantAllNotAllowed = errors.New("delete tenant all rows or global update is not allowed")

func (cfg *Config) BeforeDelete(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := GetSessionConfig(cfg, tx)

	if !sCfg.DeleteAllowTenantAll && !tx.AllowGlobalUpdate {
		if clause_checker.NoWhereClause(tx) {
			_ = tx.AddError(ErrDeleteTenantAllNotAllowed)
			return
		}
	}

	if sCfg.BeforeDeleteDoQuery && util.MapOk(tx.Statement.Clauses, "RETURNING") &&
		dialectorNotSupportReturningClause(tx.Dialector) {
		cfg.doQueryBeforeDelete(tx)
	}

	cfg.tenantClause(tx, false)
}

func (cfg *Config) AfterDelete(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
}

func (cfg *Config) doQueryBeforeDelete(tx *gorm.DB) {
	ntx := tx.Session(&gorm.Session{NewDB: true})

	ntx = _SkipQueryCallback.Set(ntx)

	if tInfo := cfg.TenantInfo(tx); tInfo != nil {
		tInfo.AddClause(ntx)
	}

	if schema := tx.Statement.Schema; schema != nil {
		slices.All(schema.QueryClauses)(func(_ int, c clause.Interface) bool {
			ntx.Statement.AddClause(c)
			return true
		})
	}

	if txClause, ok := clause_checker.WhereClause(tx); ok {
		ntx.Where(txClause)
	}

	if returning, ok := util.MapElemOk(tx.Statement.Clauses, "RETURNING"); ok {
		if returningClause, ok := returning.Expression.(clause.Returning); ok {
			slices.All(returningClause.Columns)(func(_ int, column clause.Column) bool {
				ntx.Statement.Selects = append(ntx.Statement.Selects, column.Name)
				return true
			})
		}
	} else if len(tx.Statement.Selects) != 0 {
		ntx.Statement.Selects = tx.Statement.Selects
	}

	err := ntx.Find(tx.Statement.Dest).Error
	if err != nil {
		tx.Logger.Error(tx.Statement.Context, "before delete, do query, error: %s", err.Error())
	}
}

func dialectorNotSupportReturningClause(dialector gorm.Dialector) bool {
	return util.ElemIn(dialector.Name(), gormdrivermysql.DefaultDriverName)
}
