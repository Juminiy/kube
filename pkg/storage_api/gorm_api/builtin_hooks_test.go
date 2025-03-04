package gorm_api

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
)

type Consumer struct {
	gorm.Model
	AppID    uint `gorm:"index"`
	TenantID uint `gorm:"index" mt:"tenant"`
	UserID   uint `gorm:"index"`
	VisitAt  gorm.DeletedAt
	LookupAt gorm.DeletedAt
	Region   string `gorm:"default:CN"`
}

func (c *Consumer) BeforeCreate(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		c.UserID = cast.ToUint(userID)
	}
	c.VisitAt = multi_tenants.ValidTime(c.VisitAt, tx.NowFunc())
	c.LookupAt = multi_tenants.ValidTime(c.LookupAt, tx.NowFunc().AddDate(1, 0, 0))
	return nil
}

func (c *Consumer) AfterCreate(tx *gorm.DB) error {
	c.UserID = 0
	return nil
}

func (c *Consumer) BeforeUpdate(tx *gorm.DB) error {
	tx.Logger.Info(tx.Statement.Context, "you are in PointerTo Consumer before update hooks")
	if userID, ok := tx.Get("user_id"); ok {
		tx.Statement.AddClause(ClauseUserID(tx.Statement, userID))
	}
	return nil
}

func (c *Consumer) AfterUpdate(tx *gorm.DB) error {
	tx.Logger.Info(tx.Statement.Context, "you are in PointerTo Consumer after update hooks")
	return nil
}

func (c *Consumer) BeforeDelete(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		tx.Statement.AddClause(ClauseUserID(tx.Statement, userID))
	}
	return nil
}

func (c *Consumer) AfterDelete(tx *gorm.DB) error {
	return nil
}

func (c *Consumer) AfterFind(tx *gorm.DB) error {
	c.UserID = 0
	return nil
}

func ClauseUserID(stmt *gorm.Statement, userID any) clause.Interface {
	f := multi_tenants.FieldFromSchema(stmt.Schema.FieldsByName["UserID"])
	f.Value = userID
	return &multi_tenants.Tenant{Field: f}
}

func TestBuiltinHooks(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(&Consumer{}))
}

func txHooks() *gorm.DB {
	return txMixed().Set(multi_tenants.SessionCfg, multi_tenants.SessionConfig{
		BeforeCreateMapCallHooks: true,
		UpdateMapCallHooks:       true,
		AfterFindMapCallHooks:    true,
	})
}
