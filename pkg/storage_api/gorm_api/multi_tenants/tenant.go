package multi_tenants

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/plugin_register"
	"github.com/Juminiy/kube/pkg/util"
	safe_reflectv3 "github.com/Juminiy/kube/pkg/util/safe_reflect/v3"
	"github.com/samber/lo"
	"gorm.io/gorm"
	gormschema "gorm.io/gorm/schema"
	"reflect"
)

type Config struct {
	PluginName               string
	TagKey                   string
	TagTenantKey             string
	TagUniqueKey             string
	TxTenantKey              string
	TxSkipKey                string
	DisableFieldDup          bool // effect on create and update
	ComplexFieldDup          bool // effect on create
	DeleteAllowTenantAll     bool // effect on delete, only tenant, no other where clause
	QueryBeforeDelete        bool // effect on delete
	UpdateAllowTenantAll     bool // effect on update, only tenant, no other where clause
	UpdateOmitMapZeroElemKey bool // effect on update
	AfterCreateNoHideTenant  bool // effect on create
	AfterQueryShowTenant     bool // effect on query
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
	if len(cfg.TxSkipKey) == 0 {
		cfg.TxSkipKey = "skip_tenant"
	}

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

		tx.Callback().Raw().Before("gorm:raw").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'E'), cfg.BeforeRaw),
		tx.Callback().Raw().After("gorm:raw").
			Register(plugin_register.CallbackName(cfg.PluginName, false, 'E'), cfg.AfterRaw),

		tx.Callback().Row().Before("gorm:row").
			Register(plugin_register.CallbackName(cfg.PluginName, true, 'R'), cfg.BeforeRow),
		tx.Callback().Row().After("gorm:row").
			Register(plugin_register.CallbackName(cfg.PluginName, false, 'R'), cfg.AfterRow),
	)
}

func _Ind(rv reflect.Value) safe_reflectv3.Tv {
	return safe_reflectv3.WrapI(rv)
}

func _IndI(i any) safe_reflectv3.Tv {
	return safe_reflectv3.Indirect(i)
}

func _Tag(s string) safe_reflectv3.Tag {
	return safe_reflectv3.ParseTagValue(s)
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
	Field
}

func (cfg *Config) TenantInfo(tx *gorm.DB) *Tenant {
	tenantInfoKey := util.StringJoin(":", cfg.PluginName, cfg.TagKey, cfg.TagTenantKey, cfg.TxTenantKey)
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
	_, skiptid := tx.Get(cfg.TxSkipKey)
	if !hastid || // tx no tenant_id set
		skiptid { // tx skip tenant_id
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

	return &Tenant{
		Field: FieldFromSchema(tidField).WithValue(tid),
	}
}
