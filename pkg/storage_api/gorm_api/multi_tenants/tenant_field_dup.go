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
	"strings"
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

func (f Field) Clause() clause.Expression {
	var expr clause.Expression = clause_checker.TrueExpr()
	if f.Value != nil {
		expr = f.ClauseEq()
	} else if len(f.Values) > 0 {
		expr = f.ClauseIn()
	}
	return expr
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

func (f Field) ClauseIn() clause.IN {
	return clause.IN{
		Column: clause.Column{
			Table: f.DBTable,
			Name:  f.DBName,
		},
		Values: f.Values,
	}
}

type FieldDup struct {
	Tenant      *Tenant
	Clauses     []clause.Interface
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
	slices.All(schema.Fields)(func(_ int, field *gormschema.Field) bool {
		if mt, ok := field.Tag.Lookup(cfg.TagKey); ok {
			if keys, ok := util.MapElemOk(_Tag(mt), cfg.TagUniqueKey); ok {
				columnField[field.DBName] = field.Name
				slices.All(strings.Split(keys, ","))(func(_ int, key string) bool {
					if key == "-" { // ignore field
						return false
					} else if len(key) > 0 {
						groups[key] = append(groups[key], field.Name)
					} else {
						groups[field.Name] = []string{field.Name}
					}
					return true
				})

			}
		}
		return true
	})
	if len(groups) == 0 {
		return nil
	}

	return &FieldDup{
		Tenant:      cfg.TenantInfo(tx),
		Clauses:     schema.QueryClauses,
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
	if util.ElemIn(tx.Statement.ReflectValue.Kind(), reflect.Array, reflect.Slice) &&
		!GetSessionConfig(cfg, tx).ComplexFieldDup {
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
		rval := _IndI(dest)
		switch rval.Type.Kind() {
		case reflect.Struct:
			d.FieldValue = rval.StructValues()

		case reflect.Map:
			d.ColumnValue = rval.MapValues()

		default: // ignore case
		}
	}
	d.simple(tx)
}

func (d *FieldDup) simple(tx *gorm.DB) {
	if len(d.Groups) == 0 {
		return
	}
	if len(d.FieldValue) == 0 &&
		len(d.ColumnValue) == 0 {
		return
	} else if len(d.FieldValue) == 0 {
		d.FieldValue = lo.MapKeys(d.ColumnValue, func(_ any, column string) string {
			return d.ColumnField[column]
		})
	} else if len(d.ColumnValue) == 0 {
		d.ColumnValue = lo.MapKeys(d.FieldValue, func(_ any, name string) string {
			return d.FieldColumn[name]
		})
	}

	// where clause 1. values
	orExpr, noExpr := d.rowValuesExpr()
	if noExpr {
		return
	}

	ntx := tx.Session(&gorm.Session{NewDB: true, SkipHooks: true}).
		Table(d.DBTable)

	ntx = _SkipQueryCallback.Set(ntx).
		Where(orExpr)

	// where clause 2. tenant
	if d.Tenant != nil {
		ntx.Where(d.Tenant.Clause())
	}

	// where clause 3. soft_delete or other schema(table) clauses
	slices.All(d.Clauses)(func(_ int, c clause.Interface) bool {
		ntx.Statement.AddClause(c)
		return true
	})

	// where clause 4. tx.Clause
	if txClause, ok := clause_checker.WhereClause(tx); ok {
		ntx.Where(clause.Not(txClause))
	}

	// do Count
	var cnt int64
	err := ntx.Count(&cnt).Error
	if err != nil {
		tx.Logger.Error(tx.Statement.Context, "before create or update, do field duplicated check, error: %s", err.Error())
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

// rowValuesExpr
// support one group, multiple groups
// each group one field, multiple fields
// each group if one or more fields reflect.Value.IsZero(), the group will be omitted
// if no groups, the count will be omitted
func (d *FieldDup) rowValuesExpr() (orExpr clause.Expression, noExpr bool) {
	orExpr = clause_checker.FalseExpr()
	noExpr = true
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
	return
}

func (d *FieldDup) complex(tx *gorm.DB) {
	if len(d.Groups) == 0 {
		return
	}
}

func (d *FieldDup) rowsValuesExpr(tx *gorm.DB) (expr clause.Expression, ok bool) {
	return
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
