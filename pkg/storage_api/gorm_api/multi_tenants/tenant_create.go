package multi_tenants

import "gorm.io/gorm"

func (cfg *Config) BeforeCreate(tx *gorm.DB) {
	tid, ok := cfg.tenantValid(tx)
	if !ok {
		return
	}

	stmt := tx.Statement
	refv := _Ind(stmt.ReflectValue)
	refv.SetField(map[string]any{
		"": tid,
	})
}

func (cfg *Config) AfterCreate(tx *gorm.DB) {

}
