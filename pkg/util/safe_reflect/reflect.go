package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
)

var (
	_nilValue = reflect.Value{} // comparable
)

type TypVal struct {
	// final Value and final Type
	Typ reflect.Type
	Val reflect.Value

	// original Type and original Value
	typ reflect.Type
	val reflect.Value
}

func Of(v any) TypVal {
	typOf := reflect.TypeOf(v)
	valOf := reflect.ValueOf(v)
	return TypVal{
		Typ: typOf,
		Val: valOf,
		typ: typOf,
		val: valOf,
	}
}

func IndirectOf(v any) TypVal {
	tv := Of(v)
	tv.noPointer()
	return tv
}

func indirectTV(v any) (typ reflect.Type, val reflect.Value) {
	val = noPointer(reflect.ValueOf(v))
	typ = val.Type()
	return
}

func indirectT(v any) (typ reflect.Type) {
	return underlying(reflect.TypeOf(v))
}

func indirectV(v any) (val reflect.Value) {
	return noPointer(reflect.ValueOf(v))
}

func direct(v reflect.Value) TypVal {
	typOf := v.Type()
	return TypVal{
		Typ: typOf,
		Val: v,
		typ: typOf,
		val: v,
	}
}

func indirect(v reflect.Value) TypVal {
	dv := noPointer(v)
	return TypVal{
		Typ: dv.Type(),
		Val: dv,
		typ: v.Type(),
		val: v,
	}
}

func (tv TypVal) FieldLen() int {
	return fieldLen(tv.Val)
}

// CanDirectAssign only use Type not use flag, a bit of incoming rule
func (tv TypVal) CanDirectAssign() bool {
	return util.ElemIn(tv.typ.Kind(),
		reflect.Chan,
		reflect.Interface,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice,
	)
}

func fieldLen(v reflect.Value) int {
	switch v.Kind() {
	case reflect.Struct:
		return v.NumField()
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return v.Len()
	default:
		return 0
	}
}

// String to kv string format: `type: value`
func String(v any) string {
	return ""
}

// Marshal to json string format: `type: value`
func Marshal(v any) []byte { return nil }

// Copy to new a same instance with v
func Copy(v any) any {
	return nil
}

func HasField(v any, fieldName string, fieldVal any) (ok bool) {
	tv := IndirectOf(v)
	switch tv.Typ.Kind() {
	case reflect.Struct:
		structField, exist := tv.Typ.FieldByName(fieldName)
		ok = exist && structField.Type == reflect.TypeOf(fieldVal)

	case reflect.Array, reflect.Slice:
		structField, exist := tv.Typ.Elem().FieldByName(fieldName)
		ok = exist && structField.Type == reflect.TypeOf(fieldVal)

	case reflect.Map:
		ok = tv.mapCanOpt2(fieldName, fieldVal) && tv.mapKeyExist(reflect.ValueOf(fieldName))

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
	case reflect.Struct:
		tv.StructSetFields(fields)

	case reflect.Array:
		tv.ArraySetStructFields(fields)

	case reflect.Slice:
		tv.SliceSetStructFields(fields)

	case reflect.Map:
		tv.MapAssign(fieldName, fieldVal)

	default:

	}
}

func SetFields(v any, fields map[string]any) {
	tv := IndirectOf(v)

	switch tv.Typ.Kind() {
	case reflect.Struct:
		tv.StructSetFields(fields)

	case reflect.Array:
		tv.ArraySetStructFields(fields)

	case reflect.Slice:
		tv.SliceSetStructFields(fields)

	case reflect.Map:
		for fieldName, fieldVal := range fields {
			tv.MapAssign(fieldName, fieldVal)
		}

	default:

	}
}
