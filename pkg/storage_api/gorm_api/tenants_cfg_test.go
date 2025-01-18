package gorm_api

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"gorm.io/gorm"
	"testing"
)

func TestQueryBeforeDeleteOne(t *testing.T) {
	prod := Product{}
	err := _txTenant().Set(multi_tenants.SessionCfg, multi_tenants.SessionConfig{
		QueryBeforeDelete: true,
	}).Delete(&prod, 1).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(Enc(prod))
}

func TestQueryBeforeDeleteList(t *testing.T) {
	prod := []Product{}
	err := _txTenant().Set(multi_tenants.SessionCfg, multi_tenants.SessionConfig{
		QueryBeforeDelete: true,
	}).Delete(&prod, []int{2, 3, 4}).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(Enc(prod))
}

func TestDeleteOne(t *testing.T) {
	prod := Product{}
	err := _txTenant().Delete(&prod, 5).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(Enc(prod))
}

func TestDeleteList(t *testing.T) {
	prod := Product{}
	err := _txTenant().Delete(&prod, []int{6, 7, 8}).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(Enc(prod))
}

var mtGrp = func() *gorm.DB {
	return _tx.Set("tenant_ids", []int{1, 2, 3, 4, 5, 114514})
}

func TestQueryList(t *testing.T) {
	var list []Product
	Err(t, mtGrp().Find(&list).Error)
	t.Log(Enc(list))
}
