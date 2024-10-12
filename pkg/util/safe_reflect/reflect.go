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
	typOf := directT(v)
	valOf := directV(v)
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
	return indirectT(v), indirectV(v)
}

func indirectT(v any) (typ reflect.Type) {
	return underlying(directT(v))
}

func indirectV(v any) (val reflect.Value) {
	return noPointer(directV(v))
}

func indirect(v reflect.Value) TypVal {
	indirV := noPointer(v)
	return TypVal{
		Typ: indirV.Type(),
		Val: indirV,
		typ: v.Type(),
		val: v,
	}
}

func directTV(v any) (typ reflect.Type, val reflect.Value) {
	return directT(v), directV(v)
}

func directT(v any) (typ reflect.Type) { return reflect.TypeOf(v) }

func directV(v any) (val reflect.Value) { return reflect.ValueOf(v) }

func direct(v reflect.Value) TypVal {
	t := v.Type()
	return TypVal{
		Typ: t,
		Val: v,
		typ: t,
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
		reflect.Map,
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

func HasField(v any, fieldName string, fieldVal any) (ok bool) {
	tv := IndirectOf(v)
	switch tv.Typ.Kind() {
	case reflect.Struct:
		structField, exist := tv.Typ.FieldByName(fieldName)
		ok = exist && structField.Type == directT(fieldVal)

	case reflect.Array, reflect.Slice:
		structField, exist := tv.Typ.Elem().FieldByName(fieldName)
		ok = exist && structField.Type == directT(fieldVal)

	case reflect.Map:
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
