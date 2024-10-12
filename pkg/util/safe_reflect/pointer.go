package safe_reflect

import "reflect"

// Pointer API
// +desc is for pointer, or pointer to pointer, or p to ppp....

func (tv *TypVal) noPointer() reflect.Value {
	tv.Val = noPointer(tv.Val)
	tv.Typ = tv.Val.Type()
	return tv.Val
}

// dereference _ -> _
// dereference * -> _
// dereference ** -> _
// dereference ***... -> _
func noPointer(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Pointer {
		v = reflect.Indirect(v)
	}
	return v
}

func underlying(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

func underlyingEqual(t0, t1 reflect.Type) bool {
	return underlying(t0) == underlying(t1)
}

// unused, none-sense yet
// dereference _ -> _
// dereference * -> *
// dereference ** -> *
// dereference ***... -> *
func onePointer(v reflect.Value) reflect.Value {
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
	for range capV {
		vPtr = &vPtr
	}
	return vPtr
}

// unused, none-sense
func interfacePointer(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Pointer {
		switch v.Kind() {
		case reflect.Interface:
			vInst := v.Interface()
			return reflect.ValueOf(vInst)

		case reflect.Pointer:
			v = noPointer(v)

		default:
			return v
		}
	}
	return _nilValue
}
