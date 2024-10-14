package safe_reflect

import "reflect"

// interface is equal to any to say: recommend any

func unpack(v reflect.Value) reflect.Value {
	for valueCanElem(v) {
		v = v.Elem()
	}
	return v
}

func unpackOk(v reflect.Value) (reflect.Value, bool) {
	packed := false
	for valueCanElem(v) {
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
