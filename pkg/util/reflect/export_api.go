package reflect

import "reflect"

type TypVal struct {
	v any

	Typ reflect.Type
	Val reflect.Value
}

func Parse2NoPointer(v any) (tv TypVal) {
	if v == nil {
		return
	}
	tv = TypVal{v: v}

	tv.Val = deref2NoPointer(reflect.ValueOf(v))
	tv.Typ = tv.Val.Type()
	return
}

// String to kv string(type: value)
func String(v any) string {
	return ""
}

// Marshal to json string(type: value)
func Marshal(v any) string { return "" }
