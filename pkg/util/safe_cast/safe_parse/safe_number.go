package safe_parse

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
)

// Number no precision lose Parse From string
// if parse invalid field is nil
type Number struct {
	// parsed v
	i    *int
	i8   *int8
	i16  *int16
	i32  *int32
	i64  *int64
	u    *uint
	u8   *uint8
	u16  *uint16
	u32  *uint32
	u64  *uint64
	uptr *uintptr
	f32  *float32
	f64  *float64

	// unparsed s
	s string
}

// ParseNumber
// a safe Number representation
func ParseNumber(s string) (num Number) {
	num.s = s
	if i, ok := ParseInt(s); ok {
		num.i = util.New(i)
	}
	if i8, ok := ParseInt8(s); ok {
		num.i8 = util.New(i8)
	}
	if i16, ok := ParseInt16(s); ok {
		num.i16 = util.New(i16)
	}
	if i32, ok := ParseInt32(s); ok {
		num.i32 = util.New(i32)
	}
	if i64, ok := ParseInt64(s); ok {
		num.i64 = util.New(i64)
	}
	if u, ok := ParseUint(s); ok {
		num.u = util.New(u)
	}
	if u8, ok := ParseUint8(s); ok {
		num.u8 = util.New(u8)
	}
	if u16, ok := ParseUint16(s); ok {
		num.u16 = util.New(u16)
	}
	if u32, ok := ParseUint32(s); ok {
		num.u32 = util.New(u32)
	}
	if u64, ok := ParseUint64(s); ok {
		num.u64 = util.New(u64)
	}
	if uptr, ok := ParseUintptr(s); ok {
		num.uptr = util.New(uptr)
	}
	if f32, ok := ParseFloat32(s); ok {
		num.f32 = util.New(f32)
	}
	if f64, ok := ParseFloat64(s); ok {
		num.f64 = util.New(f64)
	}
	return
}

func (n Number) U8() *uint8 {
	return n.u8
}
func (n Number) U16() *uint16 {
	return n.u16
}
func (n Number) U32() *uint32 {
	return n.u32
}
func (n Number) U64() *uint64 {
	return n.u64
}
func (n Number) I8() *int8 {
	return n.i8
}
func (n Number) I16() *int16 {
	return n.i16
}
func (n Number) I32() *int32 {
	return n.i32
}
func (n Number) I64() *int64 {
	return n.i64
}
func (n Number) F32() *float32 { return n.f32 }
func (n Number) F64() *float64 { return n.f64 }
func (n Number) S() string     { return n.s }
func (n Number) I() *int {
	return n.i
}
func (n Number) U() *uint {
	return n.u
}
func (n Number) UPtr() *uintptr {
	return n.uptr
}

// Get
// +param: kind reflect.Int ~ reflect.Float64
// +return: v is kind direct value, type: int, uint, float64, ..., etc.
func (n Number) Get(kind reflect.Kind) (v any, ok bool) {
	switch kind {
	case reflect.Int:
		if n.i != nil {
			return *n.i, true
		}
	case reflect.Int8:
		if n.i8 != nil {
			return *n.i8, true
		}
	case reflect.Int16:
		if n.i16 != nil {
			return *n.i16, true
		}
	case reflect.Int32:
		if n.i32 != nil {
			return *n.i32, true
		}
	case reflect.Int64:
		if n.i64 != nil {
			return *n.i64, true
		}
	case reflect.Uint:
		if n.u != nil {
			return *n.u, true
		}
	case reflect.Uint8:
		if n.u8 != nil {
			return *n.u8, true
		}
	case reflect.Uint16:
		if n.u16 != nil {
			return *n.u16, true
		}
	case reflect.Uint32:
		if n.u32 != nil {
			return *n.u32, true
		}
	case reflect.Uint64:
		if n.u64 != nil {
			return *n.u64, true
		}
	case reflect.Uintptr:
		if n.uptr != nil {
			return *n.uptr, true
		}
	case reflect.Float32:
		if n.f32 != nil {
			return *n.f32, true
		}
	case reflect.Float64:
		if n.f64 != nil {
			return *n.f64, true
		}
	default:
		return nil, false
	}
	return
}
