package gorm_api

import (
	"testing"
)

func TestTenantsQueryOne(t *testing.T) {
	var one Product
	Err(t, _txTenant().First(&one, 1).Error)
	t.Log(Enc(one))
}

func TestTenantsQueryList(t *testing.T) {
	var list []Product
	Err(t, _txTenant().Find(&list).Error)
	t.Log(Enc(list))
}

func TestTenantsQueryCount(t *testing.T) {
	var count int64
	Err(t, _txTenant().Model(&Product{}).Count(&count).Error)
	t.Log(count)
}

func TestQuery(t *testing.T) {
	var product Product
	Err(t, _txTenant().First(&product).Error)
	t.Log(Enc(product))
}

func TestQueryTenantClauseEqOrBugFix(t *testing.T) {
	var list []Product
	Err(t, _txTenant().
		Or("code BETWEEN ? AND ?", 100007, 100010).
		Or("code = ?", 300179).
		Find(&list).Error) // Found Tenant BUG
	t.Log(Enc(list))
}

func TestQueryTenantClauseInOrBugFix(t *testing.T) {
	var list []Product
	Err(t, mtGrp().
		Or("code BETWEEN ? AND ?", 100007, 100010).
		Or("code = ?", 300179).
		Find(&list).Error)
	t.Log(Enc(list))
}
