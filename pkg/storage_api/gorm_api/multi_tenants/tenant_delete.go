package multi_tenants

import (
	"errors"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"gorm.io/gorm"
)

var ErrDeleteTenantAllNotAllowed = errors.New("delete tenant all rows or global update is not allowed")

func (cfg *Config) BeforeDelete(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := GetSessionConfig(cfg, tx)

	if (!sCfg.DeleteAllowTenantAll || !tx.AllowGlobalUpdate) &&
		clause_checker.NoWhereClause(tx) {
		_ = tx.AddError(ErrDeleteTenantAllNotAllowed)
		return
	}

	cfg.tenantClause(tx, false)

	if sCfg.QueryBeforeDelete {
		// TODO: Query by tx where scan to tx.Dest
	}
}

func (cfg *Config) AfterDelete(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
}
