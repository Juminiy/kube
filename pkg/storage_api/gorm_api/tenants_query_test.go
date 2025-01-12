package gorm_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"gorm.io/gorm"
	"testing"
)

func TestTenantsQueryOne(t *testing.T) {
	var one Product
	err := _txTenant.First(&one, 1).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			t.Log(gorm.ErrRecordNotFound)
		default:
			util.Must(err)
		}
		return
	}
	t.Log(safe_json.Pretty(one))
}

func TestTenantsQueryList(t *testing.T) {
	var list []Product
	util.Must(_txTenant.Find(&list).Error)
	t.Log(safe_json.Pretty(list))
}

func TestTenantsQueryCount(t *testing.T) {
	var count int64
	util.Must(_txTenant.Model(&Product{}).Count(&count).Error)
	t.Log(count)
}
