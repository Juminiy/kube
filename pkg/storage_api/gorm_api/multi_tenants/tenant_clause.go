package multi_tenants

import (
	"gorm.io/gorm"
)

func (cfg *Config) tenantClause(tx *gorm.DB, forQuery bool) {
	tInfo := cfg.TenantInfo(tx)
	if tInfo == nil {
		return
	}

	tx.Where(tInfo.Clause())
	if !forQuery {
		tx.Omit(tInfo.DBName)
	}
}
