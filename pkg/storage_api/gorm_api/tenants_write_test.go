package gorm_api

import "gorm.io/gorm"

type PersonalInfo struct {
	gorm.Model
	Name     string `mt:"unique:name"`
	NumberID string `mt:"unique:name,birth"`
	Birth    int    `mt:"unique:birth"`
}
