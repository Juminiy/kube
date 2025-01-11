package gorm_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm/clause"
	"testing"
)

func TestTenantsUpdateOne(t *testing.T) {
	util.Must(_txTenant.Model(&Product{}).
		Where(clause.Eq{Column: "id", Value: 16}).
		Update("name", "Coca-Cola").Error)
}

func TestTenantsUpdateList(t *testing.T) {
	util.Must(_txTenant.Model(&Product{}).
		Update("name", gofakeit.ProductName()).Error)
}

func TestUpdate(t *testing.T) {
	err := _tx.
		Model(&Product{}).
		Where(clause.Eq{Column: "id", Value: 24}).
		Update("code", 114514).Error
	util.Must(err)
}
