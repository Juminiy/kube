package multi_tenants

import "gorm.io/gorm"

func (cfg *Config) BeforeCreate(tx *gorm.DB) {
	_, ok := cfg.tenantValid(tx)
	if !ok {
		return
	}
	stmt := tx.Statement
	refv := _Ind(stmt.ReflectValue)
	refv.SetField(map[string]any{})
}

func AfterCreate(tx *gorm.DB) {

}
