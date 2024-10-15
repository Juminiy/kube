package safe_reflect

import "reflect"

// Pointer API
// +desc is for pointer, or pointer to pointer, or p to ppp....

func (tv *TypVal) noPointer() reflect.Value {
	v, pointed := noPtrOk(tv.Val)
	if pointed {
		tv.Val = v
		tv.Typ = v.Type()
	}
	return v
}

// noPointer
// dereference _ -> _
// dereference * -> _
// dereference ** -> _
// dereference ***... -> _
func noPointer(v reflect.Value) reflect.Value {
	for v.Kind() == Ptr {
		v = v.Elem()
	}
	return v
}

func noPtrOk(v reflect.Value) (reflect.Value, bool) {
	pointed := false
	for v.Kind() == Ptr {
		v, pointed = v.Elem(), true
	}
	return v, pointed
}

// underlying
// dereference _ -> _
// dereference * -> _
// dereference ** -> _
// dereference ***... -> _
func underlying(t reflect.Type) reflect.Type {
	for t.Kind() == Ptr {
		t = t.Elem()
	}
	return t
}

func underOk(t reflect.Type) (reflect.Type, bool) {
	ok := false
	for t.Kind() == Ptr {
		t, ok = t.Elem(), true
	}
	return t, ok
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
	for v.Kind() == Ptr {
		preV = v
		v = reflect.Indirect(v)
	}
	return preV
}

// unused, none-sense
func cast2Pointer(v any, ptrLevel int) any {
	if v == nil {
		return nil
	}

	var vPtr = v
	for range ptrLevel {
		vPtr = &vPtr
	}
	return vPtr
}

// unused, none-sense
func interfacePointer(v reflect.Value) reflect.Value {
	for v.Kind() == Any ||
		v.Kind() == Ptr {
		switch v.Kind() {
		case Any:
			vInst := v.Interface()
			return directV(vInst)

		case Ptr:
			v = noPointer(v)

		default:
			return v
		}
	}
	return _zeroValue
}

func pointerType(v any, ptrLevel int) reflect.Type {
	vTyp := directT(v)
	for range ptrLevel {
		vTyp = reflect.PointerTo(vTyp)
	}
	return vTyp
}
