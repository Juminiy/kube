package reflect

import (
	"reflect"
)

type TypVal struct {
	Typ reflect.Type
	Val reflect.Value
}

// String to kv string(type: value)
func String(v any) string {
	return ""
}

// Marshal to json string(type: value)
func Marshal(v any) []byte { return nil }

var (
	_nilValue = reflect.Value{}
)

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
