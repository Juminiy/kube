package multi_tenants

import (
	"errors"
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"gorm.io/gorm"
	gormschema "gorm.io/gorm/schema"
	"slices"
)

type Field struct {
	Name    string
	DBTable string
	DBName  string
	Value   any
	Values  []any
}

type FieldDup struct {
	*Tenant
	DBTable string
	Keys    []string
	Groups  map[string][]Field // key = Keys[i], Groups[key] -> FieldGroup
}

func (cfg *Config) FieldDupInfo(tx *gorm.DB) *FieldDup {
	schema := tx.Statement.Schema
	if schema == nil {
		return nil
	}

	_ = _Ind(tx.Statement.ReflectValue).Values()
	keys := make([]string, 0, len(schema.Fields)/4)
	groups := make(map[string][]Field, cap(keys))
	slices.All(schema.Fields)(
		func(_ int, field *gormschema.Field) bool {
			if mt, ok := field.Tag.Lookup(cfg.TagKey); ok {
				if key, ok := util.MapElemOk(_Tag(mt), cfg.TagUniqueKey); ok {
					if lo.IndexOf(keys, key) == -1 {
						keys = append(keys, key)
					}
					groups[key] = append(groups[key], Field{
						Name:    field.Name,
						DBTable: schema.Table,
						DBName:  field.DBName,
						Value:   nil,
						Values:  nil,
					})
				}
			}
			return true
		})

	return &FieldDup{
		Tenant:  cfg.TenantInfo(tx),
		DBTable: schema.Table,
		Keys:    keys,
		Groups:  groups,
	}
}

type FieldDupError interface {
	error
	DBTable() string
	TenantDBName() string
	TenantValue() any
	DupDBName() []string
}

func IsFieldDupError(err error) bool {
	return errors.Is(err, fieldDupErr{})
}

type fieldDupErr struct {
	dbTable      string
	tenantDBName string
	tenantValue  any
	dbName       []string
}

func (e fieldDupErr) Error() string {
	fieldDupDesc := fmt.Sprintf("field dup error, table:[%s] column:%v",
		e.dbTable, e.dbName)
	if len(e.tenantDBName) > 0 && e.tenantValue != nil {
		return fmt.Sprintf("%s, in tenant:([%s]:[%v])",
			fieldDupDesc, e.tenantDBName, e.tenantValue)
	}
	return fieldDupDesc
}

func (e fieldDupErr) DBTable() string {
	return e.dbTable
}

func (e fieldDupErr) TenantDBName() string {
	return e.tenantDBName
}

func (e fieldDupErr) TenantValue() any {
	return e.tenantValue
}

func (e fieldDupErr) DupDBName() []string {
	return e.dbName
}
