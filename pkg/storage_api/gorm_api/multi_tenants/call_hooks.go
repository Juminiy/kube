package multi_tenants

import (
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/samber/lo"
	"gorm.io/gorm"
	gormschema "gorm.io/gorm/schema"
	"maps"
	"reflect"
	"slices"
)

// referred from: callbacks.callMethod
func CallHooks(db *gorm.DB, fc func(any, *gorm.DB) bool) {
	ntx := db.Session(&gorm.Session{NewDB: true})

	switch db.Statement.Dest.(type) {
	case map[string]any, *map[string]any:
		if structValue := _IndI(db.Statement.Model); structValue.CanAddr() { // *T -> T
			fc(structValue.Addr().Interface(), ntx)
		}

	case *[]map[string]any:
		structSlice := _IndI(db.Statement.Model).V // *[]*T -> []*T
		for i := 0; i < structSlice.Len(); i++ {
			if addrValue := reflect.Indirect(structSlice.Index(i)); addrValue.CanAddr() { // *T -> T
				fc(addrValue.Addr().Interface(), ntx)
			}
		}

	default: // ignore
	}

}

func setUpDestMapStmtModel(tx *gorm.DB, sch *gormschema.Schema) {
	//if tx.Statement.Model == tx.Statement.Dest
	switch _IndI(tx.Statement.Dest).T.Kind() {
	case reflect.Slice: // *[]map[string]any, []map[string]any
		// only for create
		tx.Statement.Model = sch.MakeSlice().Interface() // *[]*T

	case reflect.Map: // *map[string]any, map[string]any
		if modelInd := _IndI(tx.Statement.Model); modelInd.T.Kind() == reflect.Struct &&
			modelInd.CanAddr() { // Model is *T, **T, ...
			// do nothing
		} else { // for create
			tx.Statement.Model = reflect.New(sch.ModelType).Interface() // *T
			if modelInd.T.Kind() == reflect.Struct && !modelInd.IsZero() {
				_IndI(tx.Statement.Model).Set(modelInd.Value)
			}
		}

	default: // ignore
	}

	switch mapValue := tx.Statement.Dest.(type) {
	// Model *T
	case map[string]any:
		_IndI(tx.Statement.Model).StructSet(toFieldValue(sch, mapValue))

	case *map[string]any:
		_IndI(tx.Statement.Model).StructSet(toFieldValue(sch, *mapValue))

		// Model *[]*T
	case *[]map[string]any:
		structSlice := _IndI(tx.Statement.Model)
		slices.All(*mapValue)(func(_ int, m map[string]any) bool {
			newElem := reflect.New(sch.ModelType)         // *T
			_Ind(newElem).StructSet(toFieldValue(sch, m)) // *T <- m
			structSlice.SliceAppend(newElem.Interface())  // Model = append(Model, *T)
			return true
		})

	default: //ignore
	}
}

func scanModelToDestMap(tx *gorm.DB) {
	switch destValue := tx.Statement.Dest.(type) {
	case map[string]any:
		scanModelValueToDestValue(_IndI(tx.Statement.Model).StructToMap(), destValue)

	case *map[string]any:
		scanModelValueToDestValue(_IndI(tx.Statement.Model).StructToMap(), *destValue)

	case *[]map[string]any:
		slices.All(_IndI(tx.Statement.Model).SliceStructValues())(func(i int, m map[string]any) bool {
			scanModelValueToDestValue(m, (*destValue)[i])
			return true
		})

	default: // ignore
	}
}

func scanModelValueToDestValue(modelValue, destValue map[string]any) {
	maps.All(modelValue)(func(field string, modelFv any) bool {
		if destFv, ok := destValue[field]; ok && reflect.ValueOf(modelFv).IsZero() {
			delete(destValue, field)
		} else if (!ok || reflect.ValueOf(destFv).IsZero()) &&
			safe_reflect.CanDirectCompare(reflect.TypeOf(modelFv)) &&
			!reflect.ValueOf(modelFv).IsZero() {
			destValue[field] = modelFv
		}
		return true
	})
}

func scanDestMapToModel(tx *gorm.DB) {
	// omit embedded fields
	switch destValue := tx.Statement.Dest.(type) {
	case map[string]any:
		_IndI(tx.Statement.Model).StructSet(destValue)

	case *map[string]any:
		_IndI(tx.Statement.Model).StructSet(*destValue)

	case *[]map[string]any:
		modelSlice := _IndI(tx.Statement.Model)
		slices.All(*destValue)(func(i int, m map[string]any) bool {
			_Ind(modelSlice.Index(i)).StructSet(m)
			return true
		})
	}
}

func toFieldValue(sch *gormschema.Schema, values map[string]any) map[string]any {
	return lo.MapKeys(values, func(_ any, columnOrField string) string {
		field := sch.LookUpField(columnOrField)
		if field != nil {
			return field.Name
		}
		return ""
	})
}

func toColumnValue(sch *gormschema.Schema, values map[string]any) map[string]any {
	return lo.MapKeys(values, func(_ any, columnOrField string) string {
		field := sch.LookUpField(columnOrField)
		if field != nil {
			return field.DBName
		}
		return ""
	})
}
