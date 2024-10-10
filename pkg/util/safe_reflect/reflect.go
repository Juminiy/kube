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
	valOf := reflect.ValueOf(v)
	typOf := valOf.Type()
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
	val = deref2NoPointer(reflect.ValueOf(v))
	typ = val.Type()
	return
}

func indirectT(v any) (typ reflect.Type) {
	return deref2Underlying(reflect.TypeOf(v))
}

func indirectV(v any) (val reflect.Value) {
	return deref2NoPointer(reflect.ValueOf(v))
}

func of(v reflect.Value) TypVal {
	typOf := v.Type()
	return TypVal{
		Typ: typOf,
		Val: v,
		typ: typOf,
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
