package multi_tenants

import (
	"errors"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"gorm.io/gorm"
)

var ErrUpdateTenantAllNotAllowed = errors.New("update tenant all rows or global update is not allowed")

func (cfg *Config) BeforeUpdate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := GetSessionConfig(cfg, tx)

	if (!sCfg.UpdateAllowTenantAll || !tx.AllowGlobalUpdate) &&
		clause_checker.NoWhereClause(tx) {
		_ = tx.AddError(ErrUpdateTenantAllNotAllowed)
		return
	}

	if !sCfg.DisableFieldDup {
		cfg.FieldDupCheck(tx, false)
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

}
