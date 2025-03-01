package multi_tenants

import (
	"errors"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
)

var ErrUpdateTenantAllNotAllowed = errors.New("update tenant all rows or global update is not allowed")

func (cfg *Config) BeforeUpdate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := GetSessionConfig(cfg, tx)

	if !sCfg.UpdateAllowTenantAll && !tx.AllowGlobalUpdate {
		if clause_checker.NoWhereClause(tx) {
			_ = tx.AddError(ErrUpdateTenantAllNotAllowed)
			return
		}
	}

	if !sCfg.DisableFieldDup {
		cfg.FieldDupCheck(tx, true)
		if tx.Error != nil {
			return
		}
	}

	cfg.tenantClause(tx, false)

	if sCfg.UpdateOmitMapZeroElemKey {
		// TODO: update omit key of map elem is zero
	}
}

func (cfg *Config) AfterUpdate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
}

// TODO: fix
// referred from: callbacks.BeforeUpdate
func beforeUpdateMapCallHook(db *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(db); ok &&
		!db.Statement.SkipHooks && sch.BeforeUpdate {
		setUpDestMapStmtModel(db, sch)
		CallHooks(db, func(v any, tx *gorm.DB) bool {
			if beforeUpdateI, ok := v.(callbacks.BeforeUpdateInterface); ok {
				_ = db.AddError(beforeUpdateI.BeforeUpdate(tx))
				return true
			}
			return false
		})
	}
}

// referred from: callbacks.AfterUpdate
func afterUpdateMapCallHook(db *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(db); ok &&
		!db.Statement.SkipHooks && sch.AfterUpdate {
		CallHooks(db, func(v any, tx *gorm.DB) bool {
			if afterUpdateI, ok := v.(callbacks.AfterUpdateInterface); ok {
				_ = db.AddError(afterUpdateI.AfterUpdate(tx))
				return true
			}
			return false
		})
	}
}
