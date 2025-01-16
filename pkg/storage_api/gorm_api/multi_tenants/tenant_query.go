package multi_tenants

import "gorm.io/gorm"

func (cfg *Config) BeforeQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}

	if _SkipQueryCallbackBeforeWriteCountUnique.Get(tx) {
		return
	}

	cfg.tenantWhereClause(tx, true)
}

func (cfg *Config) AfterQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}

	if !GetSessionConfig(cfg, tx).AfterQueryShowTenant {
		tInfo := cfg.TenantInfo(tx)
		if tInfo == nil {
			return
		}
		_Ind(tx.Statement.ReflectValue).SetField(map[string]any{
			tInfo.Name: nil,
		})
	}
}

var _SkipQueryCallbackBeforeWriteCountUnique = SingleConfig{
	Key: "skip_query_callback_before_write_count_unique",
}
