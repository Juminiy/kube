package multi_tenants

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	gormschema "gorm.io/gorm/schema"
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
	Tenant      *Tenant
	DeletedAt   *Field
	DBTable     string
	FieldColumn map[string]string
	FieldValue  map[string]any
	ColumnField map[string]string
	ColumnValue map[string]any
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

	columnField := make(map[string]string, len(schema.DBNames)/4)
	groups := make(map[string][]string, len(schema.DBNames)/4)
	slices.All(schema.Fields)(
		func(_ int, field *gormschema.Field) bool {
			if mt, ok := field.Tag.Lookup(cfg.TagKey); ok {
				if key, ok := util.MapElemOk(_Tag(mt), cfg.TagUniqueKey); ok {
					columnField[field.DBName] = field.Name
					groups[key] = append(groups[key], field.Name)
				}
			}
			return true
		})
	if len(groups) == 0 {
		return nil
	}

	return &FieldDup{
		Tenant:      cfg.TenantInfo(tx),
		DeletedAt:   DeletedAt(schema),
		DBTable:     schema.Table,
		FieldColumn: util.MapVK(columnField),
		ColumnField: columnField,
		Groups:      groups,
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
		d.FieldValue = rval.StructValues()
		d.simple(tx)

	case reflect.Map:
		d.FieldValue = rval.MapValues()
		d.simple(tx)

	case reflect.Slice, reflect.Array:
		d.complex(tx)

	default: // ignore case
	}
}

func (d *FieldDup) Update(tx *gorm.DB) {
	dest := tx.Statement.Dest
	switch columnValue := dest.(type) {
	case map[string]any:
		d.ColumnValue = columnValue

	case *map[string]any:
		d.ColumnValue = *columnValue

	default:
		d.ColumnValue = _IndI(dest).MapValues()
	}
	d.simple(tx)
}

func (d *FieldDup) simple(tx *gorm.DB) {
	if len(d.Groups) == 0 {
		return
	}
	if len(d.FieldValue) == 0 {
		d.FieldValue = lo.MapKeys(d.ColumnValue, func(_ any, column string) string {
			return d.ColumnField[column]
		})
	} else if len(d.ColumnValue) == 0 {
		d.ColumnValue = lo.MapKeys(d.FieldValue, func(_ any, name string) string {
			return d.FieldColumn[name]
		})
	} else {
		return
	}

	// where clause 1. values
	var orExpr clause.Expression = clause_checker.FalseExpr()
	var noExpr = true
	slices.All(lo.MapToSlice(d.Groups, func(_ string, names []string) clause.Expression {
		var andExpr clause.Expression = clause_checker.TrueExpr()
		slices.All(names)(func(_ int, name string) bool {
			fieldValue, ok := d.FieldValue[name]
			if !ok || _IndI(fieldValue).Value.IsZero() {
				andExpr = nil
				return false
			}
			andExpr = clause.And(andExpr, clause.Eq{
				Column: d.FieldColumn[name],
				Value:  fieldValue,
			})
			return true
		})
		return andExpr
	}))(func(_ int, expression clause.Expression) bool {
		if expression == nil {
			return true
		}
		noExpr = false
		orExpr = clause.Or(orExpr, expression)
		return true
	})
	if noExpr {
		return
	}

	ntx := tx.Session(&gorm.Session{NewDB: true, SkipHooks: true}).
		Table(d.DBTable)

	ntx = _SkipWriteBeforeCount.Set(ntx).
		Where(orExpr)

	// where clause 2. tenant
	if d.Tenant != nil {
		ntx.Where(d.Tenant.ClauseEq())
	}

	// where clause 3. soft_delete
	if d.DeletedAt != nil {
		// maybe not required,
		// check SkipHooks whether effect on soft_delete
	}

	// where clause 4. tx.Clause
	if txClause, ok := clause_checker.WhereClause(tx); ok {
		ntx.Where(clause.Not(txClause))
	}

	// do Count
	var cnt int64
	err := ntx.Count(&cnt).Error
	if err != nil {
		tx.Logger.Error(tx.Statement.Context, "before write, do field duplicated check, error: %s", err.Error())
		return
	}
	if cnt > 0 {
		fdErr := fieldDupErr{
			dbTable: d.DBTable,
			dbName:  util.MapKeys(d.ColumnField),
		}
		if d.Tenant != nil {
			fdErr.tenantDBName = d.Tenant.DBName
			fdErr.tenantValue = d.Tenant.Value
		}
		_ = tx.AddError(fdErr)
	}
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
	return _IndI(err).Type == _fieldDupErrRType
}

var _fieldDupErrRType = _IndI(fieldDupErr{}).Type

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
