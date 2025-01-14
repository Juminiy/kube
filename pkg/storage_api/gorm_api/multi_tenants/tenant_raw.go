package multi_tenants

import (
	"gorm.io/gorm"
)

func (cfg *Config) BeforeRaw(tx *gorm.DB) { // sql.DB.Exec
	if tx.Error != nil {
		return
	}
	// TODO: clause Builder and clause Checker
}

func (cfg *Config) AfterRaw(tx *gorm.DB) {

}
