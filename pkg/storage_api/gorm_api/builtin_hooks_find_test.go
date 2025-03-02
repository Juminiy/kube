package gorm_api

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"testing"
)

func TestFindMapHooks(t *testing.T) {
	var consumerMap map[string]any
	Err(t, txHooks().Table(`tbl_consumer`).First(&consumerMap, 156).Error)
	t.Log(Enc(consumerMap))

	var consumerMapList []map[string]any
	Err(t, txHooks().Table(`tbl_consumer`).Find(&consumerMapList, 157, 158, 159).Error)
	t.Log(Enc(consumerMapList))
}

func TestBeforeQueryOmitField(t *testing.T) {
	var list []AppUser
	Err(t, _txTenant().Set(multi_tenants.SessionCfg, multi_tenants.SessionConfig{
		BeforeQueryOmitField: true,
	}).First(&list).Error)
	t.Log(Enc(list))
}
