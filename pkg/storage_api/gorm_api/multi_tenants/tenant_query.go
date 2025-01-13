package multi_tenants

import "gorm.io/gorm"

func (cfg *Config) BeforeQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}

	if _SkipQueryCallbackBeforeWriteCountUnique.Get(tx) {
		return
	}

	cfg.tenantWhereClause(tx)
}

func (cfg *Config) AfterQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
}

var _SkipQueryCallbackBeforeWriteCountUnique = Skip{Key: "skip_query_callback_before_write_count_unique"}
