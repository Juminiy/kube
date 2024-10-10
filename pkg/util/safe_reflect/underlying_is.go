package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
)

func IsBool(v reflect.Value) bool {
	return underlyingIsBool(v)
}

func underlyingIsBool(v reflect.Value) bool {
	return deref2NoPointer(v).Kind() == reflect.Bool
}

func IsInt(v reflect.Value) bool {
	return underlyingIsInt(v)
}

func underlyingIsInt(v reflect.Value) bool {
	return util.ElemIn(deref2NoPointer(v).Kind(),
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64)
}

func IsUint(v reflect.Value) bool {
	return underlyingIsUint(v)
}

func underlyingIsUint(v reflect.Value) bool {
	return util.ElemIn(deref2NoPointer(v).Kind(),
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr)
}

func IsFloat(v reflect.Value) bool {
	return underlyingIsFloat(v)
}

func underlyingIsFloat(v reflect.Value) bool {
	return util.ElemIn(deref2NoPointer(v).Kind(), reflect.Float32, reflect.Float64)
}

func IsArray(v reflect.Value) bool {
	return underlyingIsArray(v)
}

func underlyingIsArray(v reflect.Value) bool {
	return deref2NoPointer(v).Kind() == reflect.Array
}

func IsChan(v reflect.Value) bool {
	return underlyingIsChan(v)
}

func underlyingIsChan(v reflect.Value) bool {
	return deref2NoPointer(v).Kind() == reflect.Chan
}

func IsFunc(v reflect.Value) bool {
	return underlyingIsFunc(v)
}

func underlyingIsFunc(v reflect.Value) bool {
	return deref2NoPointer(v).Kind() == reflect.Func
}

func IsMap(v reflect.Value) bool {
	return underlyingIsMap(v)
}

func underlyingIsMap(v reflect.Value) bool {
	return deref2NoPointer(v).Kind() == reflect.Map
}

func IsSlice(v reflect.Value) bool {
	return underlyingIsSlice(v)
}

func underlyingIsSlice(v reflect.Value) bool {
	return deref2NoPointer(v).Kind() == reflect.Slice
}

func IsString(v reflect.Value) bool {
	return underlyingIsString(v)
}

func underlyingIsString(v reflect.Value) bool {
	return deref2NoPointer(v).Kind() == reflect.String
}

func IsStruct(v reflect.Value) bool {
	return underlyingIsStruct(v)
}

func underlyingIsStruct(v reflect.Value) bool {
	return deref2NoPointer(v).Kind() == reflect.Struct
}
