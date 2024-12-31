package multi_tenants

import (
	safe_reflectv3 "github.com/Juminiy/kube/pkg/util/safe_reflect/v3"
	"gorm.io/gorm"
)

type Config struct {
	PluginName string
	Tag        string
}

func (cfg *Config) Name() string {
	return cfg.PluginName
}

func (cfg *Config) Initialize(tx *gorm.DB) error {
	if len(cfg.PluginName) == 0 {
		cfg.PluginName = "multi_tenants"
	}
	if len(cfg.Tag) == 0 {
		cfg.Tag = "mt"
	}

	return errChecker(
		tx.Callback().Create().Before("gorm:create").
			Register(cfg.callbackName(true, 'C'), BeforeCreate),
		tx.Callback().Create().After("gorm:create").
			Register(cfg.callbackName(false, 'C'), AfterCreate),

		tx.Callback().Query().Before("gorm:query").
			Register(cfg.callbackName(true, 'Q'), BeforeQuery),
		tx.Callback().Query().After("gorm:query").
			Register(cfg.callbackName(false, 'Q'), AfterQuery),

		tx.Callback().Update().Before("gorm:update").
			Register(cfg.callbackName(true, 'U'), BeforeUpdate),
		tx.Callback().Update().After("gorm:update").
			Register(cfg.callbackName(false, 'U'), AfterUpdate),

		tx.Callback().Delete().Before("gorm:delete").
			Register(cfg.callbackName(true, 'D'), BeforeDelete),
		tx.Callback().Delete().After("gorm:delete").
			Register(cfg.callbackName(false, 'D'), AfterDelete),

		tx.Callback().Raw().Before("gorm:raw").
			Register(cfg.callbackName(true, 'E'), BeforeRaw),
		tx.Callback().Raw().After("gorm:raw").
			Register(cfg.callbackName(false, 'E'), AfterRaw),

		tx.Callback().Row().Before("gorm:row").
			Register(cfg.callbackName(true, 'R'), BeforeRow),
		tx.Callback().Row().After("gorm:row").
			Register(cfg.callbackName(false, 'R'), AfterRow),
	)
}

func (cfg *Config) callbackName(before bool, do byte) (name string) {
	name += cfg.PluginName + ":"
	if before {
		name += "before_"
	} else {
		name += "after_"
	}
	switch do {
	case 'C': // create
		name += "create"
	case 'Q': // query
		name += "query"
	case 'U': // update
		name += "update"
	case 'D': // delete
		name += "delete"
	case 'E': // raw
		name += "raw"
	case 'R': // row
		name += "row"
	default:
		panic(do)
	}
	return
}

func errChecker(err ...error) error {
	for _, errElem := range err {
		if errElem != nil {
			return errElem
		}
	}
	return nil
}

var Wrap = safe_reflectv3.Wrap

/*
 * reflect.Kind -> T
 * Struct -> --(indirect)--> T
 * SliceStruct -> --(indirect)--> []T, []*...*T
 * ArrayStruct -> --(indirect)--> [N]T, [N]*...*T
 * Map -> --(indirect)--> map[string]any
 * SliceMap -> --(indirect)--> []map[string]any
 */
