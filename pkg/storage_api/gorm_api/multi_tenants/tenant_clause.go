package multi_tenants

import (
	"gorm.io/gorm"
)

func (cfg *Config) tenantWhereClause(tx *gorm.DB, forQuery bool) {
	tInfo := cfg.TenantInfo(tx)
	if tInfo == nil {
		return
	}

	if !forQuery {
		tx.Where(tInfo.ClauseEq())
	} else {
		tx.Omit(tInfo.DBName).Where(tInfo.ClauseEq())
	}
}
