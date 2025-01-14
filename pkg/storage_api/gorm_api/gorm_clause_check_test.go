package gorm_api

import (
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"gorm.io/gorm/clause"
	"testing"
)

func TestClauseCheckCommonCase(t *testing.T) {
	var productList []Product
	err := _txTenant.
		Select([]string{"id", "name", "desc", "code", "price"}).
		Omit("desc").                          // omit no effect
		Omit("`desc`").                        // omit no effect
		Where("id = 1").                       // has expr, no value: but valid
		Where("", "").                         // no expr no value, invalid
		Where("", 1).                          // no expr, has value, invalid
		Where("id <= ").                       // has expr, no value, invalid, but no judge
		Where("name LIKE ?", "").              // has expr, no value, notLike
		Where("tenant_id BETWEEN ? AND ?", 1). // no enough value, invalid
		Where("id >= ?", 1, 2).                // value overflow, invalid
		Or("id = ?", 1).                       // valid
		Or("id = ?", 2).                       // valid
		Not("id = ?", 3).                      // valid
		Not("id = ?", 4).                      // valid
		Order("id desc").
		Order("id asc").
		Order("id DESC").
		Order("id ASC").
		Order(clause.OrderBy{
			Columns: []clause.OrderByColumn{
				{},
				{},
				{},
				{},
			},
			Expression: nil,
		}).              // no order
		Limit(0).        // limit zero
		Limit(-1).       // limit negative
		Offset(-114514). // offset not -1
		Limit(10).
		Offset(0).
		Find(&productList).Error
	/*
		SELECT `id`,`name`,`desc`,`code`,`price`
		FROM `tbl_product`
		WHERE
		(id = 1
		AND id <=
		AND name LIKE ""
		OR id = 1
		OR id = 2
		AND NOT id = 3
		AND NOT id = 4
		AND `tbl_product`.`tenant_id` = 114514)
		AND `tbl_product`.`deleted_at` IS NULL
	*/
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(safe_json.Pretty(productList))
}

func TestClauseCheckRegularCase(t *testing.T) {}
