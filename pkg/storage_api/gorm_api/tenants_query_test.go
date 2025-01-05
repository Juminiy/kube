package gorm_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

func TestTenantsQueryOne(t *testing.T) {
	var one Product
	util.Must(_txTenant.First(&one, 16).Error)
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
