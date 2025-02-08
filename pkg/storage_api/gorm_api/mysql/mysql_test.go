package mysql

import (
	"errors"
	"fmt"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"gopkg.in/yaml.v3"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io"
	"os"
	"testing"
)

var _tx *gorm_api.DB

var _cfg = struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Protocol string `yaml:"Protocol"`
	Addr     string `yaml:"Addr"`
	DBName   string `yaml:"DBName"`
}{}

func init() {
	cfgPath, err := os.Open("mysql8.yaml")
	util.Must(err)
	cfgBytes, err := io.ReadAll(cfgPath)
	util.Must(err)
	util.Must(yaml.Unmarshal(cfgBytes, &_cfg))
	tx, err := gorm_api.New(gorm.Config{
		Dialector: gormmysql.Open(
			fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local",
				_cfg.Username, _cfg.Password, _cfg.Protocol, _cfg.Addr, _cfg.DBName)),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:         "tbl_",
			SingularTable:       true,
			NameReplacer:        nil,
			NoLowerCase:         false,
			IdentifierMaxLength: 255,
		},
		PrepareStmt: true,
	})
	util.Must(err)
	util.Must(tx.Use(&multi_tenants.Config{
		PluginName: "multi_tenants",
	}))
	util.Must(tx.Use(&clause_checker.Config{
		PluginName:                 "clause_checker",
		AllowWriteClauseToRawOrRow: true,
		BeforePlugins:              []string{"multi_tenants"},
	}))
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
			errors.Is(err, gorm.ErrRecordNotFound) ||
			FullGroupBy(err) {
			t.Log(err)
		} else {
			util.Must(err)
		}
	}
}
