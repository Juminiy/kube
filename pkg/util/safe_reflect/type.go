package safe_reflect

import (
	"fmt"
	"reflect"
	"time"

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

const (
	EFace = Any
	IFace = Any
)

// CanDirectAssign only use Type not use flag, a bit of incoming rule
func (tv TypVal) CanDirectAssign() bool {
	if tv.typ == nil {
		return false
	}
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
	return CanDirectCompare(tv.Typ)
}

func CanDirectCompare(typ reflect.Type) bool {
	return util.ElemIn(typ.Kind(),
		Bool,
		Int, I8, I16, I32, I64,
		Uint, U8, U16, U32, U64, UPtr,
		F32, F64,
		String,
	)
}

// StructType
// Struct and Ptr Struct
// Slice Struct, Arr Struct, Ptr Slice Ptr Struct, Ptr Arr Ptr Struct
// Map [K] Struct, Ptr Map [K] Ptr Struct
// Any cast
func StructType(v any) (typ reflect.Type, ok bool) {
	typ = indirectT(v)
loopOf:
	switch typ.Kind() {
	case Struct:
		return typ, true

	case Slice, Arr, Map:
		typ = underlying(typ.Elem())
		goto loopOf

	default:
		return nil, false
	}
}

func StructGetTag2(v any, app0, key0, app1, key1 string) (tag0, tag1 TagVV, ok bool) {
	typ, ok := StructType(v)
	if !ok {
		return
	}

	ok = true
	tag0, tag1 = structParseTag2(typ, app0, key0, app1, key1)
	return
}

var errorType = directT((*error)(nil)).Elem()
var stringerType = directT((*fmt.Stringer)(nil)).Elem()
var timeType = directT(time.Time{})

func rValueType(v reflect.Value) reflect.Type {
	if v.IsValid() {
		return v.Type()
	}
	return nil
}
