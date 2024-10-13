package safe_reflect

import (
	"reflect"

	"github.com/Juminiy/kube/pkg/util"
)

func IsBool(v reflect.Value) bool {
	return underlyingIsBool(v)
}

func underlyingIsBool(v reflect.Value) bool {
	return noPointer(v).Kind() == Bool
}

func IsInt(v reflect.Value) bool {
	return underlyingIsInt(v)
}

func underlyingIsInt(v reflect.Value) bool {
	return util.ElemIn(noPointer(v).Kind(),
		Int, I8, I16, I32, I64)
}

func IsUint(v reflect.Value) bool {
	return underlyingIsUint(v)
}

func underlyingIsUint(v reflect.Value) bool {
	return util.ElemIn(noPointer(v).Kind(),
		Uint, U8, U16, U32, U64, UPtr)
}

func IsFloat(v reflect.Value) bool {
	return underlyingIsFloat(v)
}

func underlyingIsFloat(v reflect.Value) bool {
	return util.ElemIn(noPointer(v).Kind(),
		F32, F64)
}

func IsArray(v reflect.Value) bool {
	return underlyingIsArray(v)
}

func underlyingIsArray(v reflect.Value) bool {
	return noPointer(v).Kind() == Arr
}

func IsChan(v reflect.Value) bool {
	return underlyingIsChan(v)
}

func underlyingIsChan(v reflect.Value) bool {
	return noPointer(v).Kind() == Chan
}

func IsFunc(v reflect.Value) bool {
	return underlyingIsFunc(v)
}

func underlyingIsFunc(v reflect.Value) bool {
	return noPointer(v).Kind() == Func
}

func IsMap(v reflect.Value) bool {
	return underlyingIsMap(v)
}

func underlyingIsMap(v reflect.Value) bool {
	return noPointer(v).Kind() == Map
}

func IsSlice(v reflect.Value) bool {
	return underlyingIsSlice(v)
}

func underlyingIsSlice(v reflect.Value) bool {
	return noPointer(v).Kind() == Slice
}

func IsString(v reflect.Value) bool {
	return underlyingIsString(v)
}

func underlyingIsString(v reflect.Value) bool {
	return noPointer(v).Kind() == String
}

func IsStruct(v reflect.Value) bool {
	return underlyingIsStruct(v)
}

func underlyingIsStruct(v reflect.Value) bool {
	return noPointer(v).Kind() == Struct
}
