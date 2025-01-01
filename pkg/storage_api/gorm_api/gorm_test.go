package gorm_api

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	gormsqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
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
	util.Must(tx.Default().AutoMigrate(&Product{}))
	util.Must(tx.Use(&multi_tenants.Config{}))
	tx.DB = tx.Debug()
	_tx = tx
}

type Product struct {
	gorm.Model
	Name     string `mt:"unique:name"`
	Desc     string
	Code     uint
	Price    int64
	TenantID uint `gorm:"index;" mt:"tenant" json:"-"`
}

var Enc = safe_json.Pretty
var Dec = safe_json.From

func TestCreate(t *testing.T) {
	var product = Product{
		Name:  "Beff1Ton",
		Desc:  "one ton of beef",
		Code:  114514,
		Price: 177013,
	}
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
