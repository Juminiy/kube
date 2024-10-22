package safe_reflect

import "reflect"

// interface is equal to any to say: recommend any

// WARNING: maybe cause dead-loop, be careful when use them

const (
	_unpackLoopMax = 8 // safe_guard
)

// unpackOf
// do not export
func unpackOf(v any) TypVal {
	tv := Of(v)
	tv.unpack()
	return tv
}

func (tv *TypVal) unpack() reflect.Value {
	v, packed := unpackOk(tv.Val)
	if packed {
		tv.Val = v
		tv.Typ = v.Type()
	}
	return v
}

func unpack(v reflect.Value) reflect.Value {
	for i := 0; valueCanElem(v) && i < _unpackLoopMax; i++ {
		v = v.Elem()
	}
	return v
}

func unpackOk(v reflect.Value) (reflect.Value, bool) {
	packed := false
	for i := 0; valueCanElem(v) && i < _unpackLoopMax; i++ {
		v, packed = v.Elem(), true
	}
	return v, packed
}

func unpackV(v any) reflect.Value {
	return unpack(directV(v))
}

func unpackT(v any) reflect.Type {
	return unpackV(v).Type()
}

func unpackEqual(v0, v1 any) bool {
	return unpackV(v0) == unpackV(v1)
}
