package multi_tenants

import (
	"gorm.io/gorm"
)

func (cfg *Config) BeforeQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}

	if _SkipQueryCallback.Ok(tx) {
		return
	}

	cfg.tenantClause(tx, true)
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

var _SkipQueryCallback = Cfg{
	key: "skip_query_callback",
}
