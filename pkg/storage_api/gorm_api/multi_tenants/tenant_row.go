package multi_tenants

import (
	"gorm.io/gorm"
)

func (cfg *Config) BeforeRow(tx *gorm.DB) { // sql.DB.Query
	if tx.Error != nil {
		return
	}
	// TODO: clause Builder and clause Checker
}

func (cfg *Config) AfterRow(tx *gorm.DB) {

}
