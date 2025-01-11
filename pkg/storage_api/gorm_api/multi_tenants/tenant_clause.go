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

type Skip struct {
	Key string
}

func (s Skip) Set(tx *gorm.DB) *gorm.DB {
	return tx.Set(s.Key, struct{}{})
}

func (s Skip) Get(tx *gorm.DB) bool {
	_, ok := tx.Get(s.Key)
	return ok
}
