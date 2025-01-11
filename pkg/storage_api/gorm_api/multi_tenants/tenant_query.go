package multi_tenants

import "gorm.io/gorm"

func (cfg *Config) BeforeQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	if _SkipWriteBeforeCount.Get(tx) {
		return
	}
	cfg.tenantWhereClause(tx)
}

func (cfg *Config) AfterQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
}

var _SkipWriteBeforeCount = Skip{Key: "skip_write_before_count"}
