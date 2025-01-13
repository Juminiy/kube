package clause_checker

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/plugin_register"
	"gorm.io/gorm"
)

type Config struct {
	PluginName string
	// effect on where clause
	LikeNoPrefixMatch    bool // ignore or warn of (column LIKE '%which')
	IndexColumnNoExpr    bool // ignore or warn of indexed_column use expr or function
	InExprMaxValuesLen   *int
	OrderByNoIndexColumn bool
	BinaryExprStrongType bool
	NoRegexp             bool

	// effect on where clause on raw and row
	AllowWrapRawOrRowByClause bool
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
			Register(cfg.PluginName+":before_delete", cfg.WhereClause),
		tx.Callback().Update().Before("gorm:before_update").
			Register(cfg.PluginName+":before_delete", cfg.WhereClause),
		tx.Callback().Query().Before("gorm:query").
			Register(cfg.PluginName+":before_delete", cfg.WhereClause),
	)
}
