package multi_tenants

import (
	"github.com/samber/lo"
	"gorm.io/gorm"
	gormschema "gorm.io/gorm/schema"
	"reflect"
	"slices"
)

func (cfg *Config) BeforeCreate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := GetSessionConfig(cfg, tx)

	if !sCfg.DisableFieldDup {
		cfg.FieldDupCheck(tx, false)
		if tx.Error != nil {
			return
		}
	}

	// TODO: add Create Map with Default And NotNull Values
	// for example: gorm.Model.CreatedAt, gorm.Model.UpdatedAt
	if _, ok := hasSchemaAndDestIsMap(tx); ok {

	}

	if tInfo := cfg.TenantInfo(tx); tInfo != nil {
		_Ind(tx.Statement.ReflectValue).SetField(map[string]any{
			tInfo.Field.Name: tInfo.Field.Value, // FieldName
		})
	}
}

func (cfg *Config) AfterCreate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}

	// write back MapType's autoIncrement primaryKey values
	if sch, ok := hasSchemaAndDestIsMap(tx); ok {
		autoIncPk := lo.Filter(sch.PrimaryFields, func(item *gormschema.Field, _ int) bool {
			return item.AutoIncrement
		})

		// this func can be many choices: addAutoIncPkNameByDBName
		autoIncPkFunc := replaceAutoIncPkDBNameToName

		// Create Map gorm can write back primaryKey values
		// but Map[key] is DBName(ColumnName) not Name(FieldName)
		// Map Type Support in gorm.Scan
		// 1. Create(map[string]any{})
		// 2. Create(&map[string]any{})
		// 3. Create(&[]map[string]any{})
		switch mapValue := tx.Statement.Dest.(type) {
		case map[string]any:
			autoIncPkFunc(autoIncPk, mapValue, mapValue)

		case *map[string]any:
			autoIncPkFunc(autoIncPk, *mapValue, *mapValue)

		case *[]map[string]any:
			mapSz := len(*mapValue) / 2
			dstMap, srcMap := (*mapValue)[:mapSz], (*mapValue)[mapSz:]
			slices.All(dstMap)(func(i int, m map[string]any) bool {
				autoIncPkFunc(autoIncPk, m, srcMap[i])
				return true
			})
			tx.Statement.ReflectValue.Set(reflect.ValueOf(dstMap))

		default: // ignore
		}
	}

	if !GetSessionConfig(cfg, tx).AfterCreateShowTenant {
		if tInfo := cfg.TenantInfo(tx); tInfo != nil {
			_Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tInfo.Field.Name: nil, // FieldName
			})
		}
	}
}

// Replace Create Map Key:
// (Map[DBName] -> Value) To (Map[Name] -> Value)
func replaceAutoIncPkDBNameToName(autoIncPk []*gormschema.Field, dstMap, srcMap map[string]any) {
	slices.All(autoIncPk)(func(_ int, field *gormschema.Field) bool {
		if srcV, ok := srcMap[field.DBName]; ok {
			delete(dstMap, field.DBName)
			dstMap[field.Name] = srcV
		}
		return true
	})
}

// Add Create Map Key:
// (Map[Name] -> Value) By (Map[DBName] -> Value)
func addAutoIncPkNameByDBName(autoIncPk []*gormschema.Field, dstMap, srcMap map[string]any) {
	slices.All(autoIncPk)(func(_ int, field *gormschema.Field) bool {
		if srcV, ok := srcMap[field.DBName]; ok {
			//delete(dstMap, field.DBName)
			dstMap[field.Name] = srcV
		}
		return true
	})
}

func hasSchemaAndDestIsMap(tx *gorm.DB) (sch *gormschema.Schema, ok bool) {
	sch = tx.Statement.Schema
	return sch,
		sch != nil &&
			_Ind(tx.Statement.ReflectValue).T.Indirect().Kind() == reflect.Map
}
