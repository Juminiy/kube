package gorm_api

import (
	"gorm.io/gorm"
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

func TestCreate5(t *testing.T) {
	Err(t, _txTenant().Create(&Product{
		Name:       "Juice",
		Desc:       "Juicy Orange Blue",
		NetContent: "500ml/pack",
		Code:       110110,
		Price:      330,
	}).Error)
	Err(t, _txTenant().Model(&Product{}).Create(&map[string]any{
		"Name":       "Butter",
		"Desc":       "Anda",
		"NetContent": "220g/ounce",
		"Code":       130122,
		"Price":      577,
	}).Error)
}

func TestDoDelete(t *testing.T) {
	Err(t, _txTenant().
		Or("code BETWEEN ? AND ?", 100007, 100014).
		Or("code = ?", 300179). // Found Tenant BUG
		Delete(&Product{}).Error)
}

func TestDoCreate(t *testing.T) {
	Err(t, _txTenant().Create(&Product{
		Name:       "Shoes",
		Desc:       "Walk and Run",
		NetContent: "2pack/pair",
		Code:       300179,
		Price:      29900,
	}).Error)
}

type Product0 struct {
	gorm.Model
	Name       string
	Desc       string `gorm:"default:no_body"`
	NetContent string
	Code       uint
	Price      int64
	TenantID   uint `gorm:"index;" mt:"tenant" json:"-"`
}

func TestCreateProduct0(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(&Product0{}))
	Err(t, _txTenant().Create(&Product0{
		Name:       "Choco Milk",
		Desc:       "",
		NetContent: "335ml",
		Code:       1178921,
		Price:      3330,
	}).Error)
}

func TestCreateMapOneWriteBackAutoIncPk(t *testing.T) {
	mapOne := map[string]any{"Name": "Beer", "Desc": "Local lager beer", "NetContent": "500ml", "Code": 100007, "Price": 500}
	Err(t, _txTenant().Model(&Product0{}).Create(mapOne).Error)
	t.Log(Enc(mapOne))
}

func TestCreateMapPtrOneWriteBackAutoIncPk(t *testing.T) {
	mapOne := map[string]any{"Name": "Bee", "Desc": "Local lager bee", "NetContent": "500ml", "Code": 100011, "Price": 500}
	Err(t, _txTenant().Model(&Product0{}).Create(&mapOne).Error)
	t.Log(Enc(mapOne))
}

func TestCreateMapListWriteBackAutoIncPk(t *testing.T) {
	/*mapList := []map[string]any{
		{"Name": "Noodles", "Desc": "Instant noodles", "NetContent": "5 packs", "Code": 100008, "Price": 1000},
		{"Name": "Shampoo", "Desc": "Herbal shampoo", "NetContent": "400ml", "Code": 100009, "Price": 2500},
		{"Name": "Toothpaste", "Desc": "Mint toothpaste", "NetContent": "120g", "Code": 100010, "Price": 800},
	}
	// unsupported Create MapList
	Err(t, _txListDup().Model(&Product{}).Create(mapList).Error)
	t.Log(Enc(mapList))*/
}

func TestCreateMapListPtrWriteBackAutoIncPk(t *testing.T) {
	mapList := []map[string]any{
		{"Name": "Noodle", "Desc": "Instant noodle", "NetContent": "5 packs", "Code": 100012, "Price": 1000},
		{"Name": "Shampo", "Desc": "Herbal shampo", "NetContent": "400ml", "Code": 100013, "Price": 2500},
		{"Name": "Toothpast", "Desc": "Mint toothpast", "NetContent": "120g", "Code": 100014, "Price": 800},
	}
	Err(t, _txListDup().Model(&Product0{}).Create(&mapList).Error)
	t.Log(Enc(mapList))
}
