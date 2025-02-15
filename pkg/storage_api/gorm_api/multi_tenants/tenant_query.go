package multi_tenants

import (
	"gorm.io/gorm"
)

func (cfg *Config) BeforeQuery(tx *gorm.DB) {
	if tx.Error != nil || _SkipQueryCallback.Ok(tx) {
		return
	}
	if GetSessionConfig(cfg, tx).BeforeQueryOmitField {
		cfg.beforeQueryOmit(tx)
	}

	cfg.tenantClause(tx, true)
}

func (cfg *Config) AfterQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}

	if !GetSessionConfig(cfg, tx).AfterQueryShowTenant {
		if tInfo := cfg.TenantInfo(tx); tInfo != nil {
			_Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tInfo.Field.Name: nil, // FieldName
			})
		}
	}
}

var _SkipQueryCallback = Cfg{
	key: "skip_query_callback",
}

func (cfg *Config) beforeQueryOmit(tx *gorm.DB) {
	// replaced by gorm tag `->:false`
}
