// Package safe_reflect

// API operation

// Type API

// (1). Type Get

// (2). Type Make

// (3). Type Wrapper

// Value API

// (1). Value Get

// (2). Value Make

// (3). Value Set Self

// (4). Value Set FieldValue, ElemValue

// (5). Value Wrapper

package safe_reflect

import (
	"reflect"
)

type TypVal struct {
	// final Value and final Type
	Typ reflect.Type
	Val reflect.Value

	// original Type and original Value
	typ reflect.Type
	//val reflect.Value
}

func Of(v any) TypVal {
	typOf := directT(v)
	valOf := directV(v)
	return TypVal{
		Typ: typOf,
		Val: valOf,
		typ: typOf,
		//val: valOf,
	}
}

func IndirectOf(v any) TypVal {
	tv := Of(v)
	tv.noPointer()
	return tv
}

// HasField
// check type:
// 1. struct that do not allow embedded field
// 2. array elem is struct
// 3. slice elem is struct
// 4. map[string]any
// 5. latter: []map[string]any
func HasField(v any, fieldName string, fieldVal any) (ok bool) {
	tv := IndirectOf(v)
	switch tv.Typ.Kind() {
	case Struct:
		structField, exist := tv.Typ.FieldByName(fieldName)
		ok = exist && structField.Type == directT(fieldVal)

	case Arr, Slice:
		underElemTyp := underlying(tv.Typ.Elem())
		if underElemTyp.Kind() != Struct {
			return false
		}
		structField, exist := underElemTyp.FieldByName(fieldName)
		ok = exist && structField.Type == directT(fieldVal)

	case Map:
		ok = tv.mapCanOpt2(fieldName, fieldVal) && tv.mapKeyExist(directV(fieldName))

	default:
		ok = false
	}
	return
}

func HasFields(v any, fields map[string]any) bool {
	for fieldName, fieldVal := range fields {
		if !HasField(v, fieldName, fieldVal) {
			return false
		}
	}
	return true
}

func SetField(v any, fieldName string, fieldVal any) {
	tv := IndirectOf(v)

	fields := map[string]any{fieldName: fieldVal}
	switch tv.Typ.Kind() {
	case Struct:
		tv.StructSetFields(fields)

	case Arr:
		tv.ArraySetStructFields(fields)

	case Slice:
		tv.SliceSetStructFields(fields)

	case Map:
		tv.MapAssign(fieldName, fieldVal)

	default:

	}
}

func SetFields(v any, fields map[string]any) {
	tv := IndirectOf(v)

	switch tv.Typ.Kind() {
	case Struct:
		tv.StructSetFields(fields)

	case Arr:
		tv.ArraySetStructFields(fields)

	case Slice:
		tv.SliceSetStructFields(fields)

	case Map:
		for fieldName, fieldVal := range fields {
			tv.MapAssign(fieldName, fieldVal)
		}

	default:

	}
}

// GetTags
// +example:
//
//	type T1 struct {
//		Name string `gorm:"column:name"`
//		Desc string `gorm:"column:desc"`
//	}
//
// +param: v=T1{}, app="gorm", key="column"
// +result: {"name", "desc"}
func GetTags(v any, app, key string) []string {
	tv := IndirectOf(v)

	switch tv.Typ.Kind() {
	case Struct:
		return tv.StructParseTag(app).GetValList(key)

	case Arr, Slice:
		return tv.SliceOrArrayStructGetTagVal(app, key)

	default:
		return nil
	}
}

// comment of unused and unrealized function
/*// String to kv string format: `type: value`
func String(v any) string {
	return ""
}

// Marshal to json string format: `type: value`
func Marshal(v any) []byte { return nil }

// Copy to new a same instance with v
func Copy(v any) any {
	return nil
}*/
