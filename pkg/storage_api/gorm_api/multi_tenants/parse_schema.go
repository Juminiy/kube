package multi_tenants

import (
	"database/sql"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	gormschema "gorm.io/gorm/schema"
	"slices"
	"time"
)

func (cfg *Config) ParseSchema(tx *gorm.DB) {
	stmt := tx.Statement
	if cfg.UseTableParseSchema && len(stmt.Table) != 0 &&
		(stmt.Schema == nil || // no Schema
			(stmt.Schema != nil && (len(stmt.Schema.Name) == 0 || // has unnamed Schema
				func() bool {
					parsedSchema, ok := cfg.cacheStore.Load(cfg.graspSchemaKey(stmt.Table))
					return ok && parsedSchema.(*gormschema.Schema) != stmt.Schema
				}()))) { // destSchema != parsedSchema
		if zeroV, ok := cfg.cacheStore.Load(cfg.graspModelKey(stmt.Table)); ok {
			if err := stmt.Parse(zeroV); err != nil {
				tx.Logger.Error(stmt.Context, "use table parse schema error: %s", err.Error())
			}
		}
	}
}

func (cfg *Config) GraspSchema(tx *gorm.DB, zeroList ...any) {
	cfg.UseTableParseSchema = true
	slices.All(zeroList)(func(_ int, zeroV any) bool {
		stmt := tx.Statement
		err := stmt.Parse(zeroV)
		if err != nil {
			tx.Logger.Warn(stmt.Context, "user table grasp schema error: %s", err.Error())
		} else if stmt.Schema != nil {
			cfg.cacheStore.Store(cfg.graspSchemaKey(stmt.Schema.Table), stmt.Schema)
			cfg.cacheStore.Store(cfg.graspModelKey(stmt.Schema.Table), _IndI(zeroV).Interface())
		}
		return true
	})

}

func (cfg *Config) graspSchemaKey(tableName string) string {
	return util.StringJoin(":", cfg.PluginName, "grasp_schema", tableName)
}

func (cfg *Config) graspModelKey(tableName string) string {
	return util.StringJoin(":", cfg.PluginName, "grasp_model", tableName)
}

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

// Deprecated: use clause.Interface instead
func DeletedAt(schema *gormschema.Schema) *Field { // maybe not required
	deletedAt := schema.LookUpField("DeletedAt")
	if deletedAt == nil {
		deletedAt = schema.LookUpField("deleted_at")
		if deletedAt == nil { // pkg soft_delete
			return nil
		}
	}
	return util.New(FieldFromSchema(deletedAt))
}

func ValidTime(src gorm.DeletedAt, dest time.Time) gorm.DeletedAt {
	if !src.Valid {
		return gorm.DeletedAt(sql.NullTime{Valid: true, Time: dest})
	}
	return src
}
