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
			BeforeCreateMapCallHooks: true,
			AfterCreateMapCallHooks:  true,
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
	AppSecret  string  `gorm:"->:false;<-"`      // writeOnly, notWriteable
}

// create(created_at, updated_at, tenant_id, user_id, name, pumping, elephant)

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
	tx.Statement.Omit("user_id")
	return nil
}

func (w *CalicoWeave) BeforeDelete(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		tx.Statement.AddClause(ClauseUserID(tx.Statement, userID))
	}
	tx.Statement.Omit("user_id")
	return nil
}

func (w *CalicoWeave) AfterFind(tx *gorm.DB) error {
	w.UserID = util.MaxUint
	return nil
}

func TestInitFullFeatureTable(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(&CalicoWeave{}))
}

func TestCreateWithTenantUserDefaultValue(t *testing.T) {
	cw0 := CalicoWeave{
		Name:      "sandbox-1",
		LocID:     rand.Uint(),
		AppID:     rand.UintN(128),
		AppMe:     gofakeit.Name(),
		AppYr:     gofakeit.Hobby(),
		AppSecret: encrypt.Md5("my-secret"),
	}
	Err(t, txFull().Create(&cw0).Error)
	t.Log(Enc(cw0))

	cw1 := map[string]any{
		"Name":      "sandbox-2",
		"LocID":     rand.Uint(),
		"AppID":     rand.UintN(128),
		"AppMe":     gofakeit.Name(),
		"AppYr":     gofakeit.Hobby(),
		"AppSecret": encrypt.Md5("no-secret"),
	}
	Err(t, txFull().Model(&CalicoWeave{}).Create(&cw1).Error)
	t.Log(Enc(cw1))

	cw2 := map[string]any{
		"Name":      "sandbox-3",
		"LocID":     rand.Uint(),
		"AppID":     rand.UintN(128),
		"AppMe":     gofakeit.Name(),
		"AppYr":     gofakeit.Hobby(),
		"AppSecret": encrypt.Md5("show-secret"),
	}
	Err(t, txFull().Table(`tbl_calico_weave`).Create(&cw2).Error)
	t.Log(Enc(cw2))
}

func TestCreateList(t *testing.T) {
	cwList0 := []CalicoWeave{
		{Name: "handsome-5"},
		{Name: "handsome-6"},
	}
	Err(t, txFull().Create(&cwList0).Error)
	t.Log(Enc(cwList0))

	cwList1 := []map[string]any{
		{"Name": "handsome-3"},
		{"Name": "handsome-4"},
	}
	Err(t, txFull().Table(`tbl_calico_weave`).Create(&cwList1).Error)
	t.Log(Enc(cwList1))
}

func TestUpdateTenantUserSetPkToClause(t *testing.T) {
	Err(t, txFull().Updates(&CalicoWeave{
		Model:    gorm.Model{ID: 2},
		TenantID: util.MaxUint,
		UserID:   util.MaxUint,
		Name:     "MyName",
	}).Error)

	Err(t, txFull().Table(`tbl_calico_weave`).Updates(&map[string]any{
		"id":        2,
		"tenant_id": util.MaxUint,
		"user_id":   util.MaxUint,
		"name":      "MyName",
	}).Error)
}

func TestDeleteWithTenantUser(t *testing.T) {
	Err(t, txFull().Delete(&CalicoWeave{}, "id = ?", 2).Error)

	Err(t, txFull().Delete(&[]CalicoWeave{}, "id = ?", 2).Error)

	Err(t, txFull().Table(`tbl_calico_weave`).
		Where("id = ?", 2).
		Delete(map[string]any{}).Error)
}

func TestQueryWithTenant(t *testing.T) {
	var cw CalicoWeave
	Err(t, txFull().First(&cw, 10).Error)
	t.Log(Enc(cw))

	var cwList []CalicoWeave
	Err(t, txFull().Find(&cwList, 10, 11).Error)
	t.Log(Enc(cwList))
}

func TestUpdateFieldDup(t *testing.T) {
	Err(t, txFull().Table(`tbl_calico_weave`).
		Updates(map[string]any{
			"id":        3,                      // primaryKey
			"tenant_id": 114514,                 // ignored by callbacks
			"user_id":   1919810,                // ignored by hooks
			"username":  "My-Name-is-LiHua",     // ignored by callbacks unknownSetting
			"name":      "Li-Hua",               // dupGroup["name"] and Set to
			"loc_id":    33,                     // dupGroup["loc_app"] and Set to
			"app_id":    156,                    // dupGroup["loc_app"],["app"] and Set to
			"app_me":    "which-is-my-handsome", // dupGroup["app"] and Set to
			"app_yr":    "my-bingo-done",        // dupGroup["app"] and Set to
		}).Error)
}
