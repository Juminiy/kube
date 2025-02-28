package gorm_api

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	gormschema "gorm.io/gorm/schema"
	"testing"
)

type Consumer struct {
	gorm.Model
	AppID    uint `gorm:"index"`
	TenantID uint `gorm:"index" mt:"tenant"`
	UserID   uint `gorm:"index"`
	VisitAt  gorm.DeletedAt
	LookupAt gorm.DeletedAt
}

func (c *Consumer) BeforeCreate(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		c.UserID = cast.ToUint(userID)
	}
	return nil
}

func (c *Consumer) AfterCreate(tx *gorm.DB) error {
	c.UserID = 0
	return nil
}

func (c *Consumer) BeforeUpdate(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		tx.Statement.AddClause(ClauseUserID(tx.Statement.Schema.FieldsByName["UserID"], userID))
	}
	return nil
}

func (c *Consumer) AfterUpdate(tx *gorm.DB) error {
	return nil
}

func (c *Consumer) BeforeDelete(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		tx.Statement.AddClause(ClauseUserID(tx.Statement.Schema.FieldsByName["UserID"], userID))
	}
	return nil
}

func (c *Consumer) AfterDelete(tx *gorm.DB) error {
	return nil
}

func (c *Consumer) AfterFind(tx *gorm.DB) error {
	return nil
}

func ClauseUserID(field *gormschema.Field, userID any) clause.Interface {
	f := multi_tenants.FieldFromSchema(field)
	f.Value = userID
	return &multi_tenants.Tenant{Field: f}
}

func TestBuiltinHooks(t *testing.T) {
	/*Err(t, txMigrate().AutoMigrate(&Consumer{}))*/
}

func TestCallbacksBeforeCreate(t *testing.T) {
	// create Struct Hook is success
	var consumerStruct = Consumer{
		AppID: 10,
	}
	Err(t, txMixed().Create(&consumerStruct).Error)
	t.Log(Enc(consumerStruct))

	// create Map
	var consumerMap = map[string]any{
		"app_id": 11,
	}
	Err(t, txMixed().Model(&Consumer{}).Create(&consumerMap).Error)
	// no user_id
	t.Log(Enc(consumerMap))

	var consumerMap2 = map[string]any{
		"app_id": 22,
	}
	Err(t, txMixed().Table(`tbl_consumer`).Create(&consumerMap2).Error)
	// no default value
	// no user_id
	// no write back pk
	t.Log(Enc(consumerMap2))
}

func TestCallbacksBeforeUpdate(t *testing.T) {
	// update Struct Hook is success
	Err(t, txMixed().Updates(&Consumer{
		Model: gorm.Model{ID: 2},
		AppID: 20,
	}).Error)

	// update Map
}

func TestCallbacksBeforeDelete(t *testing.T) {
	Err(t, txMixed().Delete(&Consumer{
		Model: gorm.Model{ID: 2},
		AppID: 20,
	}).Error)
}
