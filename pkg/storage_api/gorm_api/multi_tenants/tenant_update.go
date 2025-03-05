package multi_tenants

import (
	"errors"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/clause_checker"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	gormschema "gorm.io/gorm/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"maps"
	"reflect"
	"slices"
)

var ErrUpdateTenantAllNotAllowed = errors.New("update tenant all rows or global update is not allowed")

func (cfg *Config) BeforeUpdate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := GetSessionConfig(cfg, tx)

	if !sCfg.UpdateAllowTenantAll && !tx.AllowGlobalUpdate {
		if clause_checker.NoWhereClause(tx) {
			_ = tx.AddError(ErrUpdateTenantAllNotAllowed)
			return
		}
	}

	if sCfg.UpdateMapSetPkToClause {
		beforeUpdateMapDeletePkAndSetPkToClause(tx)
	}

	if sCfg.UpdateOmitMapZeroElemKey {
		beforeUpdateMapDeleteZeroValueColumn(tx)
	}

	if sCfg.UpdateOmitMapUnknownKey {
		beforeUpdateMapDeleteUnknownColumn(tx)
	}

	if !sCfg.DisableFieldDup {
		cfg.FieldDupCheck(tx, true)
		if tx.Error != nil {
			return
		}
	}

	if sCfg.UpdateMapCallHooks {
		beforeUpdateMapCallHook(tx)
	}

	cfg.tenantClause(tx, false)
}

func (cfg *Config) AfterUpdate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
}

func beforeUpdateMapDeletePkAndSetPkToClause(tx *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(tx); ok {
		mapValue := _IndI(tx.Statement.Dest).MapValues()
		slices.All(sch.PrimaryFields)(func(_ int, pF *gormschema.Field) bool {
			if mapElem, ok := util.MapElemOk(mapValue, pF.DBName); ok {
				mapElemRv := reflect.ValueOf(mapElem)
				if mapElemRv.IsValid() && !mapElemRv.IsZero() {
					tx.Statement.AddClause(clause_checker.ClauseFieldEq(pF, mapElem))
				}
				_IndI(tx.Statement.Dest).MapSetField(map[string]any{pF.DBName: nil})
			}
			return true
		})
	}
}

func beforeUpdateMapDeleteUnknownColumn(tx *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(tx); ok {
		dbNames := sets.New(sch.DBNames...)
		notInDBNames := make([]string, 0, len(dbNames)/4)
		switch mapValue := tx.Statement.Dest.(type) {
		case map[string]any:
			maps.All(mapValue)(func(dbName string, _ any) bool {
				if !dbNames.Has(dbName) {
					notInDBNames = append(notInDBNames, dbName)
				}
				return true
			})
			slices.All(notInDBNames)(func(_ int, dbName string) bool {
				delete(mapValue, dbName)
				return true
			})

		case *map[string]any:
			maps.All(*mapValue)(func(dbName string, _ any) bool {
				if !dbNames.Has(dbName) {
					notInDBNames = append(notInDBNames, dbName)
				}
				return true
			})
			slices.All(notInDBNames)(func(_ int, dbName string) bool {
				delete(*mapValue, dbName)
				return true
			})

		default: // ignore
		}
	}
}

func beforeUpdateMapDeleteZeroValueColumn(tx *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(tx); ok {
		mapValue := _IndI(tx.Statement.Dest).MapValues()
		slices.All(sch.Fields)(func(_ int, field *gormschema.Field) bool {
			if mapElem, ok := util.MapElemOk(mapValue, field.DBName); ok {
				mapElemRv := reflect.ValueOf(mapElem)
				if mapElemRv.IsValid() && mapElemRv.IsZero() {
					_IndI(tx.Statement.Dest).MapSetField(map[string]any{field.DBName: nil})
				}
			}
			return true
		})
	}
}

// no need to call Model for Hooks,
// gorm will do: callbacks.SetupUpdateReflectValue
// we only need to do:
//  1. db.Statement.Model and set before
//  2. set Config.BeforeUpdate before callbacks.SetupUpdateReflectValue
//
// detail in: Config.Initialize
// referred from: callbacks.BeforeUpdate
func beforeUpdateMapCallHook(db *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(db); ok &&
		!db.Statement.SkipHooks && sch.BeforeUpdate {
		setUpDestMapStmtModel(db, sch)
		/*CallHooks(db, func(v any, tx *gorm.DB) bool {
			if beforeUpdateI, ok := v.(callbacks.BeforeUpdateInterface); ok {
				_ = db.AddError(beforeUpdateI.BeforeUpdate(tx))
				return true
			}
			return false
		})*/
	}
}

// no need to call Model for Hooks, gorm will do: callbacks.SetupUpdateReflectValue
// referred from: callbacks.AfterUpdate
func afterUpdateMapCallHook(db *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(db); ok &&
		!db.Statement.SkipHooks && sch.AfterUpdate {
		/*CallHooks(db, func(v any, tx *gorm.DB) bool {
			if afterUpdateI, ok := v.(callbacks.AfterUpdateInterface); ok {
				_ = db.AddError(afterUpdateI.AfterUpdate(tx))
				return true
			}
			return false
		})*/
	}
}
