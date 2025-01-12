package gorm_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestTenantsDeleteOne(t *testing.T) {
	util.Must(_txTenant.Delete(&Product{}, 1).Error)
}

func TestTenantsDeleteList(t *testing.T) {
	util.Must(_txTenant.Delete(&Product{}).Error)
}
