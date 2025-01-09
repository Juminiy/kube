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

func FromSchema(field *gormschema.Field) Field {
	return Field{
		Name:    field.Name,
		DBTable: field.Schema.Table,
		DBName:  field.DBName,
	}
}

func DeletedAt(schema *gormschema.Schema) *Field {
	deletedAt := schema.LookUpField("DeletedAt")
	if deletedAt == nil {
		deletedAt = schema.LookUpField("deleted_at")
		if deletedAt == nil {
			return nil
		}
	}
	return util.New(FromSchema(deletedAt))
}

func (f Field) WithValue(v ...any) Field {
	if len(v) > 2 {
		f.Values = v
	} else if len(v) == 1 {
		f.Value = v[0]
	}
	return f
}

type FieldDup struct {
	*Tenant
	DeletedAt *Field
	DBTable   string
	Keys      []string
	Groups    map[string][]Field // key = Keys[i], Groups[key] -> FieldGroup
}

// FieldDupInfo
// Create(&Struct), Create(&[]Struct), Create(&[N]Struct)
// Create(&map[string]any{}), Create(&[]map[string]any{})
// Create map[string]any ~ Map[K]V, K(string) is FieldName, V(any) is FieldValue
// Updates(&Struct)
// Updates(&map[string]any{})
// Updates map[string]any ~ Map[K]V, K(string) is ColumnName, V(any) is FieldValue
func (cfg *Config) FieldDupInfo(tx *gorm.DB) *FieldDup {
	schema := tx.Statement.Schema
	if schema == nil {
		return nil
	}

	keys := make([]string, 0, len(schema.Fields)/4)
	groups := make(map[string][]Field, cap(keys))
	slices.All(schema.Fields)(
		func(_ int, field *gormschema.Field) bool {
			if mt, ok := field.Tag.Lookup(cfg.TagKey); ok {
				if key, ok := util.MapElemOk(_Tag(mt), cfg.TagUniqueKey); ok {
					if lo.IndexOf(keys, key) == -1 {
						keys = append(keys, key)
					}
					groups[key] = append(groups[key], FromSchema(field))
				}
			}
			return true
		})

	return &FieldDup{
		Tenant:    cfg.TenantInfo(tx),
		DeletedAt: DeletedAt(schema),
		DBTable:   schema.Table,
		Keys:      keys,
		Groups:    groups,
	}
}

func (cfg *Config) FieldDupCheck(tx *gorm.DB, forUpdate bool) {
	dupInfo := cfg.FieldDupInfo(tx)
	if dupInfo == nil {
		return
	}
	if forUpdate {
		dupInfo.Update(tx) // update map, struct
		return
	}
	dupInfo.Create(tx) // create
}

func (d *FieldDup) Create(tx *gorm.DB) {

}

func (d *FieldDup) Update(tx *gorm.DB) {

}

func (d *FieldDup) simple(tx *gorm.DB) {

}

func (d *FieldDup) complex(tx *gorm.DB) {

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
