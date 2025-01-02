package multi_tenants

import "gorm.io/gorm"

func (cfg *Config) BeforeQuery(tx *gorm.DB) {

}

func (cfg *Config) AfterQuery(tx *gorm.DB) {

}
