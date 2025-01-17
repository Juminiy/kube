package gorm_api

import (
	"testing"
)

func TestTenantsCreateOne(t *testing.T) {
	err := _txTenant().Create(&Product{
		Name:       "Coca-Cola",
		Desc:       "Most Popular Drink in the World",
		NetContent: "500ml",
		Code:       8,
		Price:      299,
	}).Error
	Err(t, err)
}

func TestTenantsCreateList(t *testing.T) {
	Err(t,
		_txTenant().Create(&[]Product{
			{Name: "Milk", Desc: "Fresh milk", NetContent: "1L", Code: 100001, Price: 800},
			{Name: "Bread", Desc: "Whole wheat bread", NetContent: "500g", Code: 100002, Price: 1200},
			{Name: "Rice", Desc: "Long grain rice", NetContent: "5kg", Code: 100003, Price: 4500},
			{Name: "Eggs", Desc: "Free-range eggs", NetContent: "12 pcs", Code: 100004, Price: 1500},
			{Name: "Chicken", Desc: "Fresh chicken breast", NetContent: "1kg", Code: 100006, Price: 3500},
		}).Error)
}

func TestTenantsCreateMap(t *testing.T) {
	Err(t,
		_txTenant().Model(&Product{}).Create(&map[string]any{
			"Name":       "Apple",
			"Desc":       "Fresh red apples",
			"NetContent": "1kg",
			"Code":       100005,
			"Price":      1000,
		}).Error)
}

func TestTenantsCreateMapList(t *testing.T) {
	Err(t,
		_txTenant().Model(&Product{}).Create(&[]map[string]any{
			{"Name": "Beer", "Desc": "Local lager beer", "NetContent": "500ml", "Code": 100007, "Price": 500},
			{"Name": "Noodles", "Desc": "Instant noodles", "NetContent": "5 packs", "Code": 100008, "Price": 1000},
			{"Name": "Shampoo", "Desc": "Herbal shampoo", "NetContent": "400ml", "Code": 100009, "Price": 2500},
			{"Name": "Toothpaste", "Desc": "Mint toothpaste", "NetContent": "120g", "Code": 100010, "Price": 800},
		}).Error)
}

func TestCreate(t *testing.T) {
	var product = Product{
		Name:  "Beef1Ton",
		Desc:  "one ton of beef", // group["name"] is valid
		Code:  114514,            // group["code"] is valid
		Price: 177013,
	}
	err := _txTenant().Create(&product).Error
	Err(t, err)
}

func TestCreate2(t *testing.T) {
	var product = Product{
		Name:       "Beef1Ton",
		Desc:       "", // Zero group["name"] is invalid, ignore
		NetContent: "1000kg",
		Code:       114514, // group["code"] is valid
		Price:      177013,
	}
	err := _txTenant().Create(&product).Error
	Err(t, err)
}

func TestCreate3(t *testing.T) {
	var product = Product{
		Name:       "Beef2Ton",
		Desc:       "", // Zero group["name"] is invalid, ignore
		NetContent: "2000kg",
		Code:       0, // Zero group["code"] is invalid, ignore
		Price:      177013,
	}
	err := _txTenant().Create(&product).Error
	Err(t, err)
}
