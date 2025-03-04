package gorm_api

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"gorm.io/gorm"
	"testing"
)

func TestCallbacksBeforeUpdate(t *testing.T) {
	// update Struct Hook is success
	Err(t, txMixed().Updates(&Consumer{
		Model: gorm.Model{ID: 2},
		AppID: 20,
	}).Error)

	// update Map
	Err(t, txMixed().Session(&gorm.Session{AllowGlobalUpdate: true}).
		Model(&Consumer{}).Updates(map[string]any{
		"id":     2,
		"app_id": 20,
	}).Error)
}

// gorm.ErrInvalidValue
func TestPlainUpdate(t *testing.T) { // panic
	/*Err(t, _txTenant().Model(Consumer{}).
	Where(clause_checker.ClauseColumnEq("id", 2)).
	Updates(map[string]any{
		"app_id": 20,
	}).Error)*/

	Err(t, txHooksUpdate().Model(Consumer{}).
		Where(clause_checker.ClauseColumnEq("id", 2)).
		Updates(map[string]any{
			"app_id": 20,
		}).Error)
}

func TestCallbacksBeforeUpdateHooks(t *testing.T) {
	Err(t, txHooksUpdate().Model(&Consumer{}).
		Updates(map[string]any{
			"app_id":    20,
			"id":        2,
			"no_column": "cacal", // no column named no_column // also safe
		}).Error)

	Err(t, txHooksUpdate().Table(`tbl_consumer`).
		Updates(map[string]any{
			"app_id":    20,
			"id":        2,
			"no_column": "cacal", // no column named no_column // also safe
		}).Error)

}

func txHooksUpdate() *gorm.DB {
	return txMixed().Set(multi_tenants.SessionCfg, multi_tenants.SessionConfig{
		UpdateOmitMapUnknownKey:  true,
		UpdateMapSetPkToClause:   true,
		BeforeCreateMapCallHooks: true,
		UpdateMapCallHooks:       true,
		AfterFindMapCallHooks:    true,
	})
}

func TestOmitMapZeroValue(t *testing.T) {
	Err(t, _txTenant().Set(multi_tenants.SessionCfg, multi_tenants.SessionConfig{
		UpdateOmitMapZeroElemKey: true,
	}).Table(`tbl_consumer`).
		Where("id = ?", 2).
		Updates(map[string]any{
			"app_id": 0,
		}).
		Error)
}

func TestUpdateNothing(t *testing.T) {
	Err(t, _txTenant().Updates(&Consumer{Model: gorm.Model{ID: 2}}).Error)
}
