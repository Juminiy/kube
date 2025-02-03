package clause_checker

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/plugin_register"
	"github.com/Juminiy/kube/pkg/util"
	safe_reflectv3 "github.com/Juminiy/kube/pkg/util/safe_reflect/v3"
	"gorm.io/gorm"
	"slices"
)

type Config struct {
	PluginName string

	// TODO: not implement
	// effect on where clause
	LikeNoPrefixMatch    bool // ignore or warn of (column LIKE '%which')
	IndexColumnNoExpr    bool // ignore or warn of indexed_column use expr or function
	InExprMaxValuesLen   *int // restrict of IN expression max values count
	BinaryExprStrongType bool // binary expression must strong type compare
	NoRegexp             bool // forbidden of regexp expression

	// effect on where clause, on raw and row, suggest to Wrap sql by: SELECT * FROM (your sql string)
	AllowWriteClauseToRawOrRow bool

	// TODO: not implement
	// effect on orderBy clause
	OrderByNoIndexColumn bool

	BeforePlugins []string
}

func (cfg *Config) Name() string {
	return cfg.PluginName
}

func (cfg *Config) Initialize(tx *gorm.DB) error {
	if len(cfg.PluginName) == 0 {
		return plugin_register.NoPluginName
	}

	if registerRawRowErr := plugin_register.OneError(
		tx.Callback().Raw().Before("gorm:raw").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'E'), cfg.RowRawClause),

		tx.Callback().Row().Before("gorm:row").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'R'), cfg.RowRawClause),
	); registerRawRowErr != nil {
		return registerRawRowErr
	}

	var registerBeforeErr error
	var hasBefore bool
	slices.All(cfg.BeforePlugins)(func(_ int, s string) bool {
		if util.MapOk(tx.Plugins, s) {
			hasBefore = true
			registerBeforeErr = plugin_register.OneError(
				tx.Callback().Delete().
					Before(plugin_register.CallbackName(s, true, 'D')).
					Register(plugin_register.CallbackName(cfg.PluginName, true, 'D'), cfg.Clause),

				tx.Callback().Update().
					Before(plugin_register.CallbackName(s, true, 'U')).
					Register(plugin_register.CallbackName(cfg.PluginName, true, 'U'), cfg.Clause),

				tx.Callback().Query().
					Before(plugin_register.CallbackName(s, true, 'Q')).
					Register(plugin_register.CallbackName(cfg.PluginName, true, 'Q'), cfg.Clause),
			)
			if registerBeforeErr != nil {
				return false
			}
		}
		return true
	})
	if registerBeforeErr != nil || hasBefore {
		return registerBeforeErr
	}

	return plugin_register.OneError(
		tx.Callback().Delete().Before("gorm:delete").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'D'), cfg.Clause),

		tx.Callback().Update().Before("gorm:before_update").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'U'), cfg.Clause),

		tx.Callback().Query().Before("gorm:query").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'Q'), cfg.Clause),
	)
}

var _Ind = safe_reflectv3.Indirect
var _Dir = safe_reflectv3.Direct

const SkipRawOrRow = "skip_raw_row"
