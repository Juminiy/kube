package gorm_api

import (
	"gorm.io/gorm"
	"testing"
)

type Consumer struct {
	gorm.Model
	AppID    uint `gorm:"index"`
	TenantID uint `gorm:"index" mt:"tenant" json:"-"`
	UserId   uint `gorm:"index"`
}

func (c *Consumer) BeforeCreate(tx *gorm.DB) error {
	return nil
}

func (c *Consumer) AfterCreate(tx *gorm.DB) error {
	return nil
}

func (c *Consumer) BeforeUpdate(tx *gorm.DB) error {
	return nil
}

func (c *Consumer) AfterUpdate(tx *gorm.DB) error {
	return nil
}

func (c *Consumer) BeforeDelete(tx *gorm.DB) error {
	return nil
}

func (c *Consumer) AfterDelete(tx *gorm.DB) error {
	return nil
}

func (c *Consumer) AfterFind(tx *gorm.DB) error {
	return nil
}

func TestBuiltinHooks(t *testing.T) {

}
