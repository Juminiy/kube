package gorm_api

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"golang.org/x/exp/maps"
	gormsqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/soft_delete"
	"testing"
)

var _tx *DB

func init() {
	tx, err := New(gorm.Config{
		Dialector: gormsqlite.Open("kdb.db"),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:         "tbl_",
			SingularTable:       true,
			NameReplacer:        nil,
			NoLowerCase:         false,
			IdentifierMaxLength: 255,
		},
	})
	util.Must(err)
	//util.SigNotify(func() {
	//	defer util.SilentCloseIO("gorm io", tx)
	//})
	util.Must(tx.Default().AutoMigrate(
		&Product{},
		&WrapType1{}, &WrapType2{}, &WrapType3{}))
	util.Must(tx.Use(&multi_tenants.Config{}))
	tx.DB = tx.Debug()
	_tx = tx
	_txTenant = _tx.Set("tenant_id", uint(114514))
}

type Product struct {
	gorm.Model
	Name       string `mt:"unique:name"`
	Desc       string `mt:"unique:name"`
	NetContent string
	Code       uint `mt:"unique:code"`
	Price      int64
	TenantID   uint `gorm:"index;" mt:"tenant" json:"-"`
}

var Enc = safe_json.Pretty
var Dec = safe_json.From

func TestCreate(t *testing.T) {
	var product = Product{
		Name:  "Beef1Ton",
		Desc:  "one ton of beef", // group["name"] is valid
		Code:  114514,            // group["code"] is valid
		Price: 177013,
	}
	/*
		SELECT COUNT(0)
		FROM tbl_product
		WHERE 1=1
		AND (1!=1 OR `name`='Beef1Ton' OR `desc`='one ton of beef')
		AND (1!=1 OR `code`=114514)
	*/
	err := _tx.Create(&product).Error
	util.Must(err)
}

func TestCreate2(t *testing.T) {
	var product = Product{
		Name:       "Beef1Ton",
		Desc:       "", // Zero group["name"] is invalid, ignore
		NetContent: "1000kg",
		Code:       114514, // group["code"] is valid
		Price:      177013,
	}
	/*
		SELECT COUNT(0)
		FROM tbl_product
		WHERE 1=1
		AND (1!=1 OR `code`=114514)
	*/
	err := _tx.Create(&product).Error
	util.Must(err)
}

func TestCreate3(t *testing.T) {
	var product = Product{
		Name:       "Beef1Ton",
		Desc:       "", // Zero group["name"] is invalid, ignore
		NetContent: "1000kg",
		Code:       0, // Zero group["code"] is invalid, ignore
		Price:      177013,
	}
	/*
		No COUNT SQL
	*/
	err := _tx.Create(&product).Error
	util.Must(err)
}

func TestUpdate(t *testing.T) {
	err := _tx.
		Model(&Product{}).
		Where(clause.Eq{Column: "id", Value: 1}).
		Update("name", "Beef1Ton").Error
	util.Must(err)
}

func TestDelete(t *testing.T) {
	err := _tx.Delete(&Product{}, 1).Error
	util.Must(err)
}

func TestQuery(t *testing.T) {
	var product Product
	err := _tx.First(&product).Error
	util.Must(err)
	t.Log(Enc(product))
}

type InnerType struct {
	Name string
	Desc string
}

type WrapType1 struct {
	gorm.Model
	InnerType
	WrapString string
	//WrapStruct        InnerType // Error
	//WrapPointerStruct *InnerType // Error
}

type WrapType2 struct {
	DeletedAt gorm.DeletedAt
	*InnerType
	WrapString string
}

type InnerType2 struct {
	InnerType
	InnerString string
}

type WrapType3 struct {
	DeletedAt soft_delete.DeletedAt
	InnerType2
}

func showSchema(schema *schema.Schema) string {
	return safe_json.Pretty(map[string]any{
		"Name":             schema.Name,
		"Table":            schema.Table,
		"DBNames":          schema.DBNames,
		"FieldsByName":     maps.Keys(schema.FieldsByName),
		"FieldsByBindName": maps.Keys(schema.FieldsByBindName),
		"FieldsByDBName":   maps.Keys(schema.FieldsByDBName),
	})
}

func TestSchema(t *testing.T) {
	for _, ttx := range []*gorm.DB{
		_tx.Find(&WrapType1{}),
		_tx.Find(&WrapType2{}),
		_tx.Find(&WrapType3{})} {
		t.Log(showSchema(ttx.Statement.Schema))
	}
}
