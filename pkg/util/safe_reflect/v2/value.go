package safe_reflectv2

import (
	"reflect"
)

type Value struct {
	reflect.Value
	i any
}

func Wrap(rv reflect.Value) Value {
	v := Value{Value: rv}
	if rv.CanInterface() {
		v.i = rv.Interface()
	}
	return v
}

func Wrap2(rv reflect.Value, i any) Value {
	return Value{Value: rv, i: i}
}

func (v Value) isNil() bool {
	return v.IsNil()
}

func Set(src, dst any) {
	Indirect(dst).SetI(src)
}

func SetLike(src, dst any) {
	Indirect(dst).SetILike(src)
}

// encoding/json
// null:
//	case Interface, Pointer, Map, Slice: v.SetZero()
// true, false:
//	case Bool: v.SetBool(value)
// 	case EFace: v.Set(rvalue)
// "":
//	case Bytes: v.SetBytes(b)
//	case String: v.SetString(s)
// 	case EFace: v.Set(rvalue)
// number:
// 	case String: v.SetString(s)
// 	case EFace: v.Set(n)
// 	case Number: v.SetInt(n), v.SetUint(n), v.SetFloat(n)
