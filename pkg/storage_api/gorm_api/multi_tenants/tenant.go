package multi_tenants

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/plugin_register"
	"github.com/Juminiy/kube/pkg/util"
	safe_reflectv2 "github.com/Juminiy/kube/pkg/util/safe_reflect/v2"
	safe_reflectv3 "github.com/Juminiy/kube/pkg/util/safe_reflect/v3"
	"github.com/samber/lo"
	"gorm.io/gorm"
	gormschema "gorm.io/gorm/schema"
	"reflect"
	"sync"
)

type Config struct {
	PluginName   string // no default value, "" will be error, plugin will not be effect
	TagKey       string // default: mt
	TagTenantKey string // default: tenant
	TagUniqueKey string // default: unique
	TxTenantKey  string // default: tenant_id
	TxTenantsKey string // default: tenant_ids
	TxSkipKey    string // default: skip_tenant

	GlobalCfg *SessionConfig // can be overSensed by SessionCfg

	UseTableParseSchema bool
	cacheStore          *sync.Map
}

func (cfg *Config) Name() string {
	return cfg.PluginName
}

func (cfg *Config) Initialize(tx *gorm.DB) error {
	if len(cfg.PluginName) == 0 {
		return plugin_register.NoPluginName
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
	if len(cfg.TxTenantsKey) == 0 {
		cfg.TxTenantsKey = "tenant_ids"
	}
	if len(cfg.TxSkipKey) == 0 {
		cfg.TxSkipKey = "skip_tenant"
	}

	if cfg.GlobalCfg == nil {
		cfg.GlobalCfg = _GlobalCfg
	}

	cfg.cacheStore = new(sync.Map)

	return plugin_register.OneError(
		tx.Callback().Create().Before("gorm:create").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'C'), cfg.BeforeCreate),
		tx.Callback().Create().After("gorm:after_create").
			Register(plugin_register.CallbackName(cfg.PluginName, false, 'C'), cfg.AfterCreate),

		tx.Callback().Query().Before("gorm:query").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'Q'), cfg.BeforeQuery),
		tx.Callback().Query().After("gorm:after_query").
			Register(plugin_register.CallbackName(cfg.PluginName, false, 'Q'), cfg.AfterQuery),

		tx.Callback().Update().Before("gorm:before_update").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'U'), cfg.BeforeUpdate),
		tx.Callback().Update().After("gorm:after_update").
			Register(plugin_register.CallbackName(cfg.PluginName, false, 'U'), cfg.AfterUpdate),

		tx.Callback().Delete().Before("gorm:delete").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'D'), cfg.BeforeDelete),
		tx.Callback().Delete().After("gorm:after_delete").
			Register(plugin_register.CallbackName(cfg.PluginName, false, 'D'), cfg.AfterDelete),
	)
}

func _Ind(rv reflect.Value) safe_reflectv3.Tv {
	return safe_reflectv3.WrapI(rv)
}

func _IndI(i any) safe_reflectv3.Tv {
	return safe_reflectv3.Indirect(i)
}

func _IndISet(i any) safe_reflectv2.Value {
	return safe_reflectv2.Indirect(i)
}

func _DirI(i any) safe_reflectv3.Tv {
	return safe_reflectv3.Direct(i)
}

func _Tag(s string) safe_reflectv3.Tag {
	return safe_reflectv3.ParseTagValue(s)
}

func _AS(i any) []any {
	return safe_reflectv3.ToAnySlice(i)
}

/*
 * reflect.Kind -> T
 * Struct -> --(indirect)--> T
 * SliceStruct -> --(indirect)--> []T, []*...*T
 * ArrayStruct -> --(indirect)--> [N]T, [N]*...*T
 * Map -> --(indirect)--> map[string]any
 * SliceMap -> --(indirect)--> []map[string]any
 */

type Tenant struct {
	Field Field
}

func (cfg *Config) TenantInfo(tx *gorm.DB) *Tenant {
	tenantInfoKey := util.StringJoin(":", cfg.PluginName, cfg.TagKey, cfg.TagTenantKey)
	if tInfo, ok := tx.Get(tenantInfoKey); ok {
		return tInfo.(*Tenant)
	}
	tInfo := cfg.tenantInfo(tx)
	if tInfo != nil {
		tx.Set(tenantInfoKey, tInfo)
	}
	return tInfo
}

func (cfg *Config) tenantInfo(tx *gorm.DB) *Tenant {
	tid, hastid := tx.Get(cfg.TxTenantKey)
	tids, hastids := tx.Get(cfg.TxTenantsKey)
	_, skiptid := tx.Get(cfg.TxSkipKey)
	if (!hastid && !hastids) || // tx no tenant_id or no tenant_ids set
		skiptid { // tx skip tenant_id and tenant_ids
		return nil
	}

	schema := tx.Statement.Schema
	if schema == nil { // no schema
		return nil
	}
	tidField, ok := lo.Find(schema.Fields, func(item *gormschema.Field) bool {
		if mt, ok := item.Tag.Lookup(cfg.TagKey); ok && mt == cfg.TagTenantKey {
			return true
		}
		return false
	})
	if !ok {
		return nil
	}

	field := FieldFromSchema(tidField)
	if hastid {
		field.Value = tid
	}
	if hastids {
		field.Values = _AS(tids)
	}
	return &Tenant{Field: field}
}
