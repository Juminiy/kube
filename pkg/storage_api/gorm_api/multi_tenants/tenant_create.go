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

	tInfo := cfg.TenantInfo(tx)
	if tInfo == nil {
		return
	}
	_Ind(tx.Statement.ReflectValue).SetField(map[string]any{
		tInfo.Name: tInfo.Value, // FieldName
		/*tInfo.DBName: tInfo.Value, // DBName*/
	})
}

func (cfg *Config) AfterCreate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}

	if !GetSessionConfig(cfg, tx).AfterCreateShowTenant {
		tInfo := cfg.TenantInfo(tx)
		if tInfo == nil {
			return
		}
		_Ind(tx.Statement.ReflectValue).SetField(map[string]any{
			tInfo.Name: nil,
		})
	}

	// write back MapType's autoIncrement primaryKey values
	if sch := tx.Statement.Schema; sch != nil && _Ind(tx.Statement.ReflectValue).T.Indirect().Kind() == reflect.Map {
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
			_Ind(tx.Statement.ReflectValue).Set(reflect.ValueOf(dstMap))

		default: // ignore
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
