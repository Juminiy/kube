package multi_tenants

import (
	"gorm.io/gorm"
)

func (cfg *Config) tenantWhereClause(tx *gorm.DB) {
	tInfo := cfg.TenantInfo(tx)
	if tInfo == nil {
		return
	}

	tx.Omit(tInfo.DBName).Where(tInfo.ClauseEq())
}
