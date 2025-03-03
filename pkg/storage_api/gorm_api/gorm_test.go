package gorm_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	gormsqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	util.Must(tx.Use(&multi_tenants.Config{
		PluginName:          "multi_tenants",
		UseTableParseSchema: true,
	}))
	util.Must(tx.Use(&clause_checker.Config{
		PluginName:                 "clause_checker",
		AllowWriteClauseToRawOrRow: true,
		BeforePlugins:              []string{"multi_tenants"},
	}))
	tx.Plugins["multi_tenants"].(*multi_tenants.Config).
		GraspSchema(tx.DB, &Product{}, &AppUser{}, &Consumer{}, &CalicoWeave{})
	tx.DB = tx.Debug()
	_tx = tx
}

var Enc = safe_json.Pretty
var Dec = safe_json.From
var Err = func(t *testing.T, err error) {
	if err != nil {
		if multi_tenants.IsFieldDupError(err) ||
			errors.Is(err, multi_tenants.ErrDeleteTenantAllNotAllowed) ||
			errors.Is(err, multi_tenants.ErrUpdateTenantAllNotAllowed) ||
			errors.Is(err, gorm.ErrRecordNotFound) {
			t.Log(err)
		} else {
			util.Must(err)
		}
	}
}

func TestInit(t *testing.T) {}
