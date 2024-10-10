package safe_reflect

import "reflect"

// Pointer API is for pointer, or pointer to pointer, or p to ppp....

func (tv *TypVal) noPointer() reflect.Value {
	tv.Val = deref2NoPointer(tv.Val)
	tv.Typ = tv.Val.Type()
	return tv.Val
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

func deref2Underlying(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

func underlyingTypeEq(t0, t1 reflect.Type) bool {
	return deref2Underlying(t0) == deref2Underlying(t1)
}

// unused, none-sense yet
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
	for range capV {
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
	return _nilValue
}
