package gorm_api

import (
	"gorm.io/gorm/clause"
	"testing"
)

func TestClauseCheckCommonCase(t *testing.T) {
	var productList []Product
	err := _txTenant().
		Select([]string{"id", "name", "desc", "code", "price"}).
		Omit("desc").    // omit no effect
		Omit("`desc`").  // omit no effect
		Where("id = 1"). // has expr, no value: but valid
		Where("", "").   // no expr no value, invalid
		Where("", 1).    // no expr, has value, invalid
		//Where("id <= ").                       // has expr, no value, invalid, but no judge
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
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(Enc(productList))
}

func TestClauseCheckRegularCase(t *testing.T) {}

func TestClauseWriteToRowOrRaw(t *testing.T) {
	prods := []Product{}
	err := _txTenant().Raw("SELECT * FROM `tbl_product`").
		Where("").
		Where("", "").
		Where("r = ?").
		Where("a >= ? AND b <= ?", 1, 2, 3).
		Where("id BETWEEN ? AND ?", 1, 1000).
		Where("name LIKE ?", "%co%").
		Limit(10).
		Offset(3).
		Scan(&prods).Error
	Err(t, err)
	t.Log(Enc(prods))
}
