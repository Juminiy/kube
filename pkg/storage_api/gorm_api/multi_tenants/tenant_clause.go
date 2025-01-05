package multi_tenants

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (cfg *Config) tenantWhereClause(tx *gorm.DB) {
	tInfo := cfg.TenantInfo(tx)
	if tInfo == nil {
		return
	}

	tx.Omit(tInfo.DBName).Where(clause.Eq{
		Column: clause.Column{
			Table: tInfo.DBTable,
			Name:  tInfo.DBName,
		},
		Value: tInfo.Value,
	})
}
