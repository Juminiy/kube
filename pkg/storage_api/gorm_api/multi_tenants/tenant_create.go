package multi_tenants

import (
	"gorm.io/gorm"
)

func (cfg *Config) BeforeCreate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := GetSessionConfig(cfg, tx)

	if !sCfg.DisableFieldDup {
		cfg.FieldDupCheck(tx, false)
		if tx.Error != nil {
			return
		}
	}

	tInfo := cfg.TenantInfo(tx)
	if tInfo == nil {
		return
	}
	_Ind(tx.Statement.ReflectValue).SetField(map[string]any{
		tInfo.Name: tInfo.Value, // FieldName
		/*tInfo.DBName: tInfo.Value, // DBName*/
	})
}

func (cfg *Config) AfterCreate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}

	if !GetSessionConfig(cfg, tx).AfterCreateShowTenant {
		tInfo := cfg.TenantInfo(tx)
		if tInfo == nil {
			return
		}
		_Ind(tx.Statement.ReflectValue).SetField(map[string]any{
			tInfo.Name: nil,
		})
	}
}
