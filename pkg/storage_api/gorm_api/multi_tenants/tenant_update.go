package multi_tenants

import (
	"errors"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"gorm.io/gorm"
)

var ErrUpdateTenantAllNotAllowed = errors.New("update tenant all rows or global update is not allowed")

func (cfg *Config) BeforeUpdate(tx *gorm.DB) {
	if (!cfg.UpdateAllowTenantAll || !tx.AllowGlobalUpdate) &&
		clause_checker.NoWhereClause(tx) {
		_ = tx.AddError(ErrUpdateTenantAllNotAllowed)
		return
	}
	cfg.tenantWhereClause(tx)
}

func (cfg *Config) AfterUpdate(tx *gorm.DB) {

}
