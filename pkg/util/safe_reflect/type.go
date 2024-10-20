package safe_reflect

import (
	"reflect"

	"github.com/Juminiy/kube/pkg/util"
)

// alias of reflect.Kind
const (
	Invalid reflect.Kind = iota
	Bool
	Int
	I8
	I16
	I32
	I64
	Uint
	U8
	U16
	U32
	U64
	UPtr
	F32
	F64
	C64
	C128
	Arr
	Chan
	Func
	Any
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePtr
)

// CanDirectAssign only use Type not use flag, a bit of incoming rule
func (tv TypVal) CanDirectAssign() bool {
	return util.ElemIn(tv.typ.Kind(),
		Chan, Map, Slice,
	)
}

func typeCanElem(t reflect.Type) bool {
	return util.ElemIn(t.Kind(),
		Arr, Chan, Map, Ptr, Slice,
	)
}

func (tv TypVal) CanDirectCompare() bool {
	return util.ElemIn(tv.typ.Kind(),
		Bool,
		Int, I8, I16, I32, I64,
		Uint, U8, U16, U32, U64, UPtr,
		F32, F64,
		String,
	)
}
