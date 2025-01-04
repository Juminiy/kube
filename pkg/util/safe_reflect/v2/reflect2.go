package safe_reflectv2

import (
	"reflect"
)

// type Value variable var is v
// type Type variable var is t
// type any variable var is i
// type reflect.Value variable var is rv
// type reflect.Type variable var is rt

func Indirect(i any) Value {
	return Direct(i).indirect()
}

func IndirectRV(rv reflect.Value) reflect.Value {
	return indirect(rv, false)
}

func Direct(i any) Value {
	return Value{Value: direct(i)}
}

func Set(src, dst any) {
	Indirect(dst).SetI(src)
}

func SetLike(src, dst any) {
	Indirect(dst).SetILike(src)
}

// reference to: encoding/json
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
