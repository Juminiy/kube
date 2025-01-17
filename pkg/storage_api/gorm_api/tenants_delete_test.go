package gorm_api

import (
	"testing"
)

var Delete = func(t *testing.T, i any, conds ...any) {
	Err(t, _txTenant().Delete(i, conds...).Error)
}

func TestTenantsDeleteOne(t *testing.T) {
	Delete(t, &Product{}, 1)
}

func TestTenantsDeleteList(t *testing.T) {
	Delete(t, &Product{})
}
