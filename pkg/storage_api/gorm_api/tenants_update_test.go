package gorm_api

import (
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm/clause"
	"testing"
)

/*var Update = func(t *testing.T, i any, conds ...any) {
	Err(t, _txTenant.Updates(i).Error)
}*/

func TestTenantsUpdateOne(t *testing.T) {
	Err(t, _txTenant().Model(&Product{}).
		Where(clause.Eq{Column: "id", Value: 2}).
		Update("code", "114514").Error)
}

func TestTenantsUpdateList(t *testing.T) {
	Err(t, _txTenant().Model(&Product{}).
		Update("name", gofakeit.ProductName()).Error)
}

func TestUpdate(t *testing.T) {
	Err(t, _txTenant().Model(&Product{}).
		Where(clause.Eq{Column: "id", Value: 1}).
		Update("code", 114514).Error)
}
