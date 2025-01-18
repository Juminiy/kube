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

type FieldDup struct {
	Tenant      *Tenant
	Clauses     []clause.Interface
	DBTable     string
	FieldColumn map[string]string
	ColumnField map[string]string
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
		(&rowValues{
			FieldValue: rval.StructValues(),
			FieldDup:   d,
		}).simple(tx)

	case reflect.Map:
		(&rowValues{
			FieldValue: rval.MapValues(),
			FieldDup:   d,
		}).simple(tx)

	case reflect.Slice, reflect.Array:
		(&rowsValues{
			FieldDup: d,
		}).complex(tx)

	default: // ignore case
	}
}

func (d *FieldDup) Update(tx *gorm.DB) {
	dest := tx.Statement.Dest
	switch columnValue := dest.(type) {
	case map[string]any:
		(&rowValues{
			ColumnValue: columnValue,
			FieldDup:    d,
		}).simple(tx)

	case *map[string]any:
		(&rowValues{
			ColumnValue: *columnValue,
			FieldDup:    d,
		}).simple(tx)

	default:
		rval := _IndI(dest)
		switch rval.Type.Kind() {
		case reflect.Struct:
			(&rowValues{
				FieldValue: rval.StructValues(),
				FieldDup:   d,
			}).simple(tx)

		case reflect.Map:
			(&rowValues{
				ColumnValue: rval.MapValues(),
				FieldDup:    d,
			}).simple(tx)

		default: // ignore case
		}
	}
}

func (d *FieldDup) doCount(tx *gorm.DB, orExpr clause.Expression) {
	ntx := tx.Session(&gorm.Session{NewDB: true, SkipHooks: true}).
		Table(d.DBTable)

	// where clause 1. orExpr
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

type rowValues struct {
	FieldValue  map[string]any
	ColumnValue map[string]any
	*FieldDup
}

func (d *rowValues) simple(tx *gorm.DB) {
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

	orExpr, noExpr := d.expr()
	if noExpr {
		return
	}
	d.doCount(tx, orExpr)
}

// rowValuesExpr
// support one group, multiple groups
// each group one field, multiple fields
// each group if one or more fields reflect.Value.IsZero(), the group will be omitted
// if no groups, the count will be omitted
func (d *rowValues) expr() (orExpr clause.Expression, noExpr bool) {
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

type rowsValues struct {
	List []rowValues
	*FieldDup
}

func (d *rowsValues) complex(tx *gorm.DB) {
	if len(d.Groups) == 0 {
		return
	}
	rval := _Ind(tx.Statement.ReflectValue)
	if !util.ElemIn(rval.T.Indirect().Kind(), reflect.Struct, reflect.Map) {
		return
	}
	d.List = lo.Map(rval.Values(), func(item map[string]any, _ int) rowValues {
		return rowValues{
			FieldValue: item,
			FieldDup:   d.FieldDup,
		}
	})

	orExpr, noExpr := d.expr()
	if noExpr {
		return
	}
	d.doCount(tx, orExpr)
}

func (d *rowsValues) expr() (orExpr clause.Expression, noExpr bool) {
	orExpr = clause_checker.FalseExpr()
	noExpr = true
	slices.All(d.List)(func(_ int, values rowValues) bool {
		subOrExpr, noOK := values.expr()
		if noOK {
			return true
		}
		noExpr = false
		orExpr = clause.Or(orExpr, subOrExpr)
		return true
	})
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
