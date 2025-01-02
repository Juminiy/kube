package multi_tenants

import "gorm.io/gorm"

func (cfg *Config) BeforeRow(tx *gorm.DB) {

}

func (cfg *Config) AfterRow(tx *gorm.DB) {

}
