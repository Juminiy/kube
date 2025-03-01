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
	c.UserID = 0
	return nil
}

func ClauseUserID(field *gormschema.Field, userID any) clause.Interface {
	f := multi_tenants.FieldFromSchema(field)
	f.Value = userID
	return &multi_tenants.Tenant{Field: f}
}

func TestBuiltinHooks(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(&Consumer{}))
}

func txHooks() *gorm.DB {
	return txMixed().Set(multi_tenants.SessionCfg, multi_tenants.SessionConfig{
		CreateMapCallHooks:    true,
		UpdateMapCallHooks:    true,
		AfterFindMapCallHooks: true,
	})
}

/*
	gorm do not support MapType Alias

type ConsumerMap map[string]any

	func (m ConsumerMap) BeforeCreate(tx *gorm.DB) error {
		if userID, ok := tx.Get("user_id"); ok {
			m["UserID"] = userID
		}
		return nil
	}

	func (m ConsumerMap) AfterCreate(tx *gorm.DB) error {
		delete(m, "UserID")
		return nil
	}
*/
func TestCallbacksBeforeCreate(t *testing.T) {
	// create Struct Hook is success
	var consumerStruct = Consumer{
		AppID: 11,
	}
	Err(t, txMixed().Create(&consumerStruct).Error)
	// create with user_id
	t.Log(Enc(consumerStruct))

	// create Map
	var consumerMap = map[string]any{
		"AppID": 22,
	}
	Err(t, txMixed().Model(&Consumer{}).Create(&consumerMap).Error)
	// create with no user_id
	t.Log(Enc(consumerMap))

	var consumerMap2 = map[string]any{
		"AppID": 33,
	}
	Err(t, txMixed().Table(`tbl_consumer`).Create(&consumerMap2).Error)
	// create with no user_id
	t.Log(Enc(consumerMap2))

	// unsupported type: panic
	/*var consumerMapV2 = ConsumerMap{
		"AppID": 44,
	}
	Err(t, txMixed().Table(`tbl_consumer`).Create(&consumerMapV2).Error)
	// create ? user_id ?
	t.Log(Enc(consumerMapV2))*/
}

func TestCreateMapHooks(t *testing.T) {
	// one map with hooks
	var consumerMap2 = map[string]any{
		"AppID": 33,
	}
	Err(t, txHooks().Table(`tbl_consumer`).Create(&consumerMap2).Error)
	// create with user_id
	t.Log(Enc(consumerMap2))

	// map list with hooks
	var consumerMapList = []map[string]any{}
	for i := 0; i < 100; i++ {
		consumerMapList = append(consumerMapList, map[string]any{
			"AppID": (i + 1) * 5,
		})
	}
	Err(t, txHooks().Table(`tbl_consumer`).Create(&consumerMapList).Error)
	t.Log(Enc(consumerMapList))

}

func TestFindMapHooks(t *testing.T) {
	var consumerMap map[string]any
	Err(t, txHooks().Table(`tbl_consumer`).First(&consumerMap, 156).Error)
	t.Log(Enc(consumerMap))

	var consumerMapList []map[string]any
	Err(t, txHooks().Table(`tbl_consumer`).Find(&consumerMapList, 157, 158, 159).Error)
	t.Log(Enc(consumerMapList))
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
