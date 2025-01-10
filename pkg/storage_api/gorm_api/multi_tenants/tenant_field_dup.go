package multi_tenants

import (
	"errors"
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	gormschema "gorm.io/gorm/schema"
	"maps"
	"reflect"
	"slices"
)

type Field struct {
	Name    string
	DBTable string
	DBName  string
	Value   any
	Values  []any
}

func FieldFromSchema(field *gormschema.Field) Field {
	return Field{
		Name:    field.Name,
		DBTable: field.Schema.Table,
		DBName:  field.DBName,
	}
}

func (f Field) WithValue(v ...any) Field {
	if len(v) > 2 {
		f.Values = v
	} else if len(v) == 1 {
		f.Value = v[0]
	}
	return f
}

func (f Field) ClauseEq() clause.Eq {
	return clause.Eq{
		Column: clause.Column{
			Table: f.DBTable,
			Name:  f.DBName,
		},
		Value: f.Value,
	}
}

type FieldDup struct {
	*Tenant
	DeletedAt   *Field
	DBTable     string
	ColumnField map[string]string
	FieldValue  map[string]any
	Groups      map[string][]string // Groups[key] -> FieldGroup
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

	names := make([]string, 0, len(schema.Fields)/4)
	dbNames := make([]string, 0, cap(names))
	groups := make(map[string][]string, cap(names))
	slices.All(schema.Fields)(
		func(_ int, field *gormschema.Field) bool {
			if mt, ok := field.Tag.Lookup(cfg.TagKey); ok {
				if key, ok := util.MapElemOk(_Tag(mt), cfg.TagUniqueKey); ok {
					names = append(names, field.Name)
					dbNames = append(dbNames, field.DBName)
					groups[key] = append(groups[key], field.Name)
				}
			}
			return true
		})
	if len(groups) == 0 {
		return nil
	}

	return &FieldDup{
		Tenant:    cfg.TenantInfo(tx),
		DeletedAt: DeletedAt(schema),
		DBTable:   schema.Table,
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
	rval := _Ind(tx.Statement.ReflectValue)
	switch rval.Type.Kind() {
	case reflect.Struct:
		rval.StructValues()
		d.simple(tx)

	case reflect.Map:
		rval.MapValues()
		d.simple(tx)

	case reflect.Slice, reflect.Array:
		d.complex(tx)

	default: // ignore case
	}
}

func (d *FieldDup) Update(tx *gorm.DB) {
	rval := _Ind(tx.Statement.ReflectValue)
	rval.MapValues()
	d.simple(tx)
}

func (d *FieldDup) simple(tx *gorm.DB) {
	if len(d.Groups) == 0 {
		return
	}
	ntx := tx.Session(&gorm.Session{NewDB: true, SkipHooks: true}).
		Table(d.DBTable)

	if d.Tenant != nil {
		ntx.Where(d.Tenant)
	}

	if d.DeletedAt != nil {
		// maybe not required,
		// check SkipHooks whether effect on soft_delete
	}

	maps.All(d.Groups)(func(key string, names []string) bool {

		return true
	})
}

func (d *FieldDup) complex(tx *gorm.DB) {
	if len(d.Groups) == 0 {
		return
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
