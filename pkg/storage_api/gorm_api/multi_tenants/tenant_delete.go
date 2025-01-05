package multi_tenants

import (
	"errors"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"gorm.io/gorm"
)

var ErrDeleteTenantAllNotAllowed = errors.New("delete tenant all rows or global update is not allowed")

func (cfg *Config) BeforeDelete(tx *gorm.DB) {
	if (!cfg.DeleteAllowTenantAll || !tx.AllowGlobalUpdate) &&
		clause_checker.NoWhereClause(tx) {
		_ = tx.AddError(ErrDeleteTenantAllNotAllowed)
		return
	}
	cfg.tenantWhereClause(tx)
}

func (cfg *Config) AfterDelete(tx *gorm.DB) {

}
