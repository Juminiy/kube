package reflect

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

func (tv TypVal) FieldLen() int {
	return fieldLen(tv.Val)
}

func (tv TypVal) CanAssign() bool {
	return util.ElemIn(tv.typ.Kind(),
		reflect.Chan,
		reflect.Func,
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
