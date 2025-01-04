package multi_tenants

import (
	safe_reflectv3 "github.com/Juminiy/kube/pkg/util/safe_reflect/v3"
	"gorm.io/gorm"
	"reflect"
)

type Config struct {
	PluginName      string
	TagKey          string
	TagTenantKey    string
	TagUniqueKey    string
	TxTenantKey     string
	TxSkipKey       string
	DisableFieldDup bool
	EmitMapZeroElem bool
}

func (cfg *Config) Name() string {
	return cfg.PluginName
}

func (cfg *Config) Initialize(tx *gorm.DB) error {
	if len(cfg.PluginName) == 0 {
		cfg.PluginName = "multi_tenants"
	}
	if len(cfg.TagKey) == 0 {
		cfg.TagKey = "mt"
	}
	if len(cfg.TagTenantKey) == 0 {
		cfg.TagTenantKey = "tenant"
	}
	if len(cfg.TagUniqueKey) == 0 {
		cfg.TagUniqueKey = "unique"
	}
	if len(cfg.TxTenantKey) == 0 {
		cfg.TxTenantKey = "tenant_id"
	}
	if len(cfg.TxSkipKey) == 0 {
		cfg.TxSkipKey = "skip_tenant"
	}

	return errChecker(
		tx.Callback().Create().Before("gorm:create").
			Register(cfg.callbackName(true, 'C'), cfg.BeforeCreate),
		tx.Callback().Create().After("gorm:create").
			Register(cfg.callbackName(false, 'C'), cfg.AfterCreate),

		tx.Callback().Query().Before("gorm:query").
			Register(cfg.callbackName(true, 'Q'), cfg.BeforeQuery),
		tx.Callback().Query().After("gorm:query").
			Register(cfg.callbackName(false, 'Q'), cfg.AfterQuery),

		tx.Callback().Update().Before("gorm:update").
			Register(cfg.callbackName(true, 'U'), cfg.BeforeUpdate),
		tx.Callback().Update().After("gorm:update").
			Register(cfg.callbackName(false, 'U'), cfg.AfterUpdate),

		tx.Callback().Delete().Before("gorm:delete").
			Register(cfg.callbackName(true, 'D'), cfg.BeforeDelete),
		tx.Callback().Delete().After("gorm:delete").
			Register(cfg.callbackName(false, 'D'), cfg.AfterDelete),

		tx.Callback().Raw().Before("gorm:raw").
			Register(cfg.callbackName(true, 'E'), cfg.BeforeRaw),
		tx.Callback().Raw().After("gorm:raw").
			Register(cfg.callbackName(false, 'E'), cfg.AfterRaw),

		tx.Callback().Row().Before("gorm:row").
			Register(cfg.callbackName(true, 'R'), cfg.BeforeRow),
		tx.Callback().Row().After("gorm:row").
			Register(cfg.callbackName(false, 'R'), cfg.AfterRow),
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

func _Ind(rv reflect.Value) safe_reflectv3.Tv {
	return safe_reflectv3.WrapI(rv)
}

/*
 * reflect.Kind -> T
 * Struct -> --(indirect)--> T
 * SliceStruct -> --(indirect)--> []T, []*...*T
 * ArrayStruct -> --(indirect)--> [N]T, [N]*...*T
 * Map -> --(indirect)--> map[string]any
 * SliceMap -> --(indirect)--> []map[string]any
 */

func (cfg *Config) tenantValid(tx *gorm.DB) (any, bool) {
	tid, hastid := tx.Get(cfg.TxTenantKey)
	_, skiptid := tx.Get(cfg.TxSkipKey)
	if !hastid || skiptid {
		return nil, false
	}
	return tid, true
}
