package safe_reflect

import "reflect"

// Pointer API
// +desc is for pointer, or pointer to pointer, or p to ppp....

const (
	_noPtrLoopMax = 16 // safe_guard
)

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
	for i := 0; v.Kind() == Ptr && i < _noPtrLoopMax; i++ {
		v = v.Elem()
	}
	return v
}

func noPtrOk(v reflect.Value) (reflect.Value, bool) {
	pointed := false
	for i := 0; v.Kind() == Ptr && i < _noPtrLoopMax; i++ {
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
	for i := 0; t.Kind() == Ptr && i < _noPtrLoopMax; i++ {
		t = t.Elem()
	}
	return t
}

func underOk(t reflect.Type) (reflect.Type, bool) {
	ok := false
	for i := 0; t.Kind() == Ptr && i < _noPtrLoopMax; i++ {
		t, ok = t.Elem(), true
	}
	return t, ok
}

func underlyingEqual(t0, t1 reflect.Type) bool {
	return underlying(t0) == underlying(t1)
}

func pointerType(v any, ptrLevel int) reflect.Type {
	vTyp := directT(v)
	for i := 0; i < ptrLevel && i < _noPtrLoopMax; i++ {
		vTyp = reflect.PointerTo(vTyp)
	}
	return vTyp
}

// WARNING: may cause dead-loop
func pointerCast(v any, ptrLevel int) any {
	for i := 0; i < ptrLevel && i < _noPtrLoopMax; i++ {
		v = &v
	}
	return v
}

// dereference _ -> _
// dereference * -> *
// dereference ** -> *
// dereference ***... -> *
func onePointer(v reflect.Value) reflect.Value {
	preV := v
	for i := 0; v.Kind() == Ptr && i < _noPtrLoopMax; i++ {
		preV = v
		v = v.Elem()
	}
	return preV
}

// Deprecated
// unused, none-sense
func cast2Pointer(v any, ptrLevel int) any {
	if v == nil {
		return nil
	}

	var vPtr = v
	for i := 0; i < ptrLevel && i < _noPtrLoopMax; i++ {
		vPtr = &vPtr
	}
	return vPtr
}

// Deprecated
// unused, none-sense
func interfacePointer(v reflect.Value) reflect.Value {
	for i := 0; (v.Kind() == Any || v.Kind() == Ptr) && i < _noPtrLoopMax; i++ {
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
