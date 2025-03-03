package gorm_api

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/encrypt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"math/rand/v2"
	"testing"
)

func txFull() *gorm.DB {
	return txPure().
		Set("tenant_id", 1919810).
		Set("tenant_ids", []uint{1919810, 1, 2, 3}).
		Set("user_id", 114514).
		Set(multi_tenants.SessionCfg, multi_tenants.SessionConfig{
			DisableFieldDup:          false,
			ComplexFieldDup:          true,
			DeleteAllowTenantAll:     false,
			BeforeDeleteDoQuery:      true,
			UpdateAllowTenantAll:     false,
			UpdateOmitMapZeroElemKey: false,
			UpdateOmitMapUnknownKey:  true,
			UpdateMapSetPkToClause:   true,
			AfterCreateShowTenant:    true,
			BeforeQueryOmitField:     true,
			AfterQueryShowTenant:     true,
			CreateMapCallHooks:       true,
			UpdateMapCallHooks:       true,
			AfterFindMapCallHooks:    true,
		})
}

type CalicoWeave struct {
	gorm.Model         // gorm.Model create:set, gorm.DeletedAt: update,delete,query:clause
	TenantID   uint    `gorm:"index" mt:"tenant"`    // callbacksDo create:set, update,delete,query:clause
	UserID     uint    `gorm:"index"`                // hooksDo create:set, update,delete:clause
	Name       string  `gorm:"not null" mt:"unique"` // callbacksDo not_null dup_check
	Desc       string  // do nothing
	Pumping    float64 `gorm:"default:5.5"` // callbacksDo create:setDefault
	Elephant   float64 // hooksDo create:setDefault
	LocID      uint    `mt:"unique:loc_app"`     // callbacksDo not_null dup_check
	AppID      uint    `mt:"unique:loc_app,app"` // callbacksDo not_null dup_check
	AppMe      string  `mt:"unique:app"`         // callbacksDo not_null dup_check
	AppYr      string  `mt:"unique:app"`         // callbacksDo not_null dup_check
	AppSecret  string  `gorm:"->:false"`         // readOnly
}

func (w *CalicoWeave) BeforeCreate(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		w.UserID = cast.ToUint(userID)
	}
	w.Elephant = 3.3
	return nil
}

func (w *CalicoWeave) AfterCreate(tx *gorm.DB) error {
	w.UserID = util.MaxUint
	return nil
}

func (w *CalicoWeave) BeforeUpdate(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		tx.Statement.AddClause(ClauseUserID(tx.Statement, userID))
	}
	return nil
}

func (w *CalicoWeave) BeforeDelete(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		tx.Statement.AddClause(ClauseUserID(tx.Statement, userID))
	}
	return nil
}

func TestInitFullFeatureTable(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(&CalicoWeave{}))
}

func TestCreateWithTenantUserDefaultValue(t *testing.T) {
	cw0 := CalicoWeave{
		Name:      "sandbox-9",
		Desc:      "Isolation Sandbox Environment",
		LocID:     rand.Uint(),
		AppID:     rand.UintN(128),
		AppMe:     gofakeit.Name(),
		AppYr:     gofakeit.Hobby(),
		AppSecret: encrypt.Md5("my-secret"),
	}
	Err(t, txFull().Create(&cw0).Error)
	t.Log(Enc(cw0))

	cw1 := map[string]any{
		"Name":      "sandbox-8",
		"LocID":     rand.Uint(),
		"AppID":     rand.UintN(128),
		"AppSecret": encrypt.Md5("no-secret"),
	}
	Err(t, txFull().Model(&CalicoWeave{}).Create(&cw1).Error)
	t.Log(Enc(cw1))

	cw2 := map[string]any{
		"Name":      "sandbox-7",
		"LocID":     rand.Uint(),
		"AppID":     rand.UintN(128),
		"AppSecret": encrypt.Md5("show-secret"),
	}
	Err(t, txFull().Table(`tbl_calico_weave`).Create(&cw2).Error)
	t.Log(Enc(cw2))
}
