package multi_tenants

import "gorm.io/gorm"

func (cfg *Config) BeforeRow(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
}

func (cfg *Config) AfterRow(tx *gorm.DB) {

}
