package safe_reflectv2

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
)

const (
	Invalid reflect.Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

type Type struct {
	reflect.Type
}

func FromValue(v Value) Type {
	t := Type{}
	if v.IsValid() {
		t.Type = v.Type()
	}
	return t
}

func FromWrap(rt reflect.Type) Type {
	return Type{Type: rt}
}

func (t Type) isBool() bool {
	return t.Kind() == Bool
}

func (t Type) isBoolLike() bool {
	return t.isBool()
}

func (t Type) isInt() bool {
	return util.InRange(t.Kind(), Int, Int64)
}

func (t Type) isIntLike() bool { return t.isInt() }

func (t Type) isUint() bool {
	return util.InRange(t.Kind(), Uint, Uintptr)
}

func (t Type) isUintLike() bool {
	return t.isUint()
}

func (t Type) isFloat() bool {
	return util.ElemIn(t.Kind(), Float32, Float64)
}

func (t Type) isFloatLike() bool {
	return t.isFloat()
}

func (t Type) isNumber() bool {
	return util.InRange(t.Kind(), Int, Float64)
}

func (t Type) isNumberLike() bool {
	return t.isNumber()
}

func (t Type) isString() bool {
	return t.Kind() == String
}

// fmt.Stringer
func (t Type) isStringLike() bool {
	return t.isString() || t.isBytes() || t.isNumberLike()
}

// []byte
func (t Type) isBytes() bool {
	return t.Kind() == Slice && t.Elem().Kind() == Uint8
}
