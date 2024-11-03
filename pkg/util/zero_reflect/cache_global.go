package zero_reflect

import (
	runtimeutil "github.com/Juminiy/kube/pkg/util/runtime"
	"reflect"
	"sync"
)

var _global = sync.Map{}

func TypeOf(v any) reflect.Type {
	switch _noPointerLevel {
	case _noPointer:
		return indirectTyp(v)

	case _structSpec:
		return structIndirectTyp(v)

	case _mustComparable:
		return comparableIndirectTyp(v)

	default:
		return directTyp(v)
	}
}

// put v direct type to pool
func directTyp(v any) reflect.Type {
	vTypKey := runtimeutil.EFaceOf(&v).Type()
	vTypCache, ok := _global.Load(vTypKey)
	if ok {
		return vTypCache.(reflect.Type)
	}
	vTyp := reflect.TypeOf(v)
	_global.Store(vTypKey, vTyp)
	return vTyp
}

// indirectTyp
// 1. put v NoPointer type to pool
// (1). *...*T -> T
func indirectTyp(v any) reflect.Type {
	return directTyp(v)
}

// structIndirectTyp
// 1. (1). as indirectTyp(1)
// 2. put v must struct type to pool
// (2). *...*Struct 			 -> Struct, (as (1))
// (3). *...*[]*...*T 			 -> []*...*T, (T not Struct)
// (4). *...*[num]*...*T 		 -> [num]*...*T, (T not Struct)
// (5). *...*[]*...*Struct 		 -> Struct
// (6). *...*[num]*...*Struct 	 -> Struct
// (7). map[K]E, E ~ (1),(3),(4) -> (1),(3),(4)
// (8). map[K]E, E ~ (2),(5),(6) -> Struct
func structIndirectTyp(v any) reflect.Type {
	return directTyp(v)
}

// comparableIndirectTyp
// 1. (1). as indirectTyp(1)
// 2. (2,5,6,8) as structIndirectTyp(2,5,6,8)
// 3. put v must comparable type to pool
// (9). any -> T, T must be in (reflect.Kind 1~16, Func, String, Struct)
func comparableIndirectTyp(v any) reflect.Type {
	return directTyp(v)
}
