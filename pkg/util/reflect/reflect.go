package reflect

import (
	"reflect"
)

func underlyingStruct(v reflect.Value) bool {
	return deref2NoPointer(v).Kind() == reflect.Struct
}

func fieldSize(v reflect.Value) int {
	switch v.Kind() {
	case reflect.Struct:
		return v.NumField()
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return v.Len()
	default:
		return 0
	}
}

// dereference _ -> _
// dereference * -> _
// dereference ** -> _
// dereference ***... -> _
func deref2NoPointer(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Pointer {
		v = reflect.Indirect(v)
	}
	return v
}

// dereference _ -> _
// dereference * -> *
// dereference ** -> *
// dereference ***... -> *
func deref2OnePointer(v reflect.Value) reflect.Value {
	preV := v
	for v.Kind() == reflect.Pointer {
		preV = v
		v = reflect.Indirect(v)
	}
	return preV
}

// unused, none-sense
func cast2Pointer(v any, capV int) any {
	if v == nil {
		return nil
	}

	var vPtr = v
	for _ = range capV {
		vPtr = &vPtr
	}
	return vPtr
}

// unused, none-sense
func derefInterfacePointer(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Pointer {
		switch v.Kind() {
		case reflect.Interface:
			vInst := v.Interface()
			return reflect.ValueOf(vInst)

		case reflect.Pointer:
			v = deref2NoPointer(v)

		default:
			return v
		}
	}
	return reflect.Value{}
}

var (
	_nilValue = reflect.Value{}
)

func mapKeyExistAssign(v reflect.Value, mapKey, mapVal any) {
	v = deref2NoPointer(v)
	if v.Kind() != reflect.Map || v.IsNil() {
		return
	}

	mapKeyV, mapValV := reflect.ValueOf(mapKey), reflect.ValueOf(mapVal)
	if oldMapVal := v.MapIndex(mapKeyV); oldMapVal == _nilValue {
		return
	}

	v.SetMapIndex(mapKeyV, mapValV)
}

func mapDryAssign(v reflect.Value, mapKey, mapVal any) {
	v = deref2NoPointer(v)
	if v.Kind() != reflect.Map || v.IsNil() {
		return
	}

	v.SetMapIndex(reflect.ValueOf(mapKey), reflect.ValueOf(mapVal))
}

func mapDryDelete(v reflect.Value, mapKey any) {
	mapDryAssign(v, mapKey, nil)
}
