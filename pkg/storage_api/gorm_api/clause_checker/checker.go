package clause_checker

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/plugin_register"
	safe_reflectv3 "github.com/Juminiy/kube/pkg/util/safe_reflect/v3"
	"gorm.io/gorm"
)

type Config struct {
	PluginName string

	// effect on where clause
	LikeNoPrefixMatch    bool // ignore or warn of (column LIKE '%which')
	IndexColumnNoExpr    bool // ignore or warn of indexed_column use expr or function
	InExprMaxValuesLen   *int
	BinaryExprStrongType bool
	NoRegexp             bool

	// effect on where clause on raw and row
	AllowWrapRawOrRowByClause bool

	// effect on orderBy clause
	OrderByNoIndexColumn bool
}

func (cfg *Config) Name() string {
	return cfg.PluginName
}

func (cfg *Config) Initialize(tx *gorm.DB) error {
	if len(cfg.PluginName) == 0 {
		return plugin_register.NoPluginName
	}
	return plugin_register.OneError(
		tx.Callback().Delete().Before("gorm:delete").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'D'), cfg.Clause),

		tx.Callback().Update().Before("gorm:before_update").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'U'), cfg.Clause),

		tx.Callback().Query().Before("gorm:query").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'Q'), cfg.Clause),

		tx.Callback().Raw().Before("gorm:raw").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'E'), cfg.RowRawClause),

		tx.Callback().Row().Before("gorm:row").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'R'), cfg.RowRawClause),
	)
}

var _Ind = safe_reflectv3.Indirect
var _Dir = safe_reflectv3.Direct
