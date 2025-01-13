package gorm_api

import (
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

func TestClauseCheckCommonCase(t *testing.T) {
	var productList []Product
	err := _txTenant.
		Select([]string{"id", "name", "desc", "code", "price"}).
		Omit("desc").
		Where("id = 1").                       // has expr, no value: valid
		Where("", "").                         // no expr no value
		Where("", 1).                          // no expr, has value
		Where("id <= ").                       // has expr, no value
		Where("name LIKE ?", "").              // has expr, no value
		Where("tenant_id BETWEEN ? AND ?", 1). // no enough value
		Where("id >= ?", 1, 2).                // value overflow
		Or("id = ?", 1).
		Or("id = ?", 2).
		Not("id = ?", 3).
		Not("id = ?", 4).
		Order("").       // no order
		Limit(0).        // limit zero
		Limit(-1).       // limit negative
		Offset(-114514). // offset not -1
		Find(&productList).Error
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(safe_json.Pretty(productList))
}

func TestClauseCheckRegularCase(t *testing.T) {}
