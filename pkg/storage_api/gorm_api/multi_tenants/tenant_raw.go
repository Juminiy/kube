package multi_tenants

import "gorm.io/gorm"

func (cfg *Config) BeforeRaw(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
}

func (cfg *Config) AfterRaw(tx *gorm.DB) {

}
