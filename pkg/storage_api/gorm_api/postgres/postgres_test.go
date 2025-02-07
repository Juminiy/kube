package postgres

import (
	"errors"
	"fmt"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"gopkg.in/yaml.v3"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io"
	"os"
	"testing"
)

var _tx *gorm_api.DB

var _cfg = struct {
	Host     string `yaml:"Host"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
	Port     int    `yaml:"Port"`
	SSLMode  string `yaml:"SSLMode"`
	TimeZone string `yaml:"TimeZone"`
}{}

func init() {
	cfgPath, err := os.Open("postgres.yaml")
	util.Must(err)
	cfgBytes, err := io.ReadAll(cfgPath)
	util.Must(err)
	util.Must(yaml.Unmarshal(cfgBytes, &_cfg))
	tx, err := gorm_api.New(gorm.Config{
		Dialector: gormpostgres.Open(
			fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
				_cfg.Host, _cfg.User, _cfg.Password, _cfg.DBName, _cfg.Port, _cfg.SSLMode, _cfg.TimeZone)),
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
			GroupingError(err) ||
			UndefinedColumn(err) ||
			UniqueViolation(err) {
			t.Log(err)
		} else {
			util.Must(err)
		}
	}
}
