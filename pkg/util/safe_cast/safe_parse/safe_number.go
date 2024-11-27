package safe_parse

import "github.com/Juminiy/kube/pkg/util"

// Number no precision lose Parse From string
// if parse invalid field is nil
type Number struct {
	u8   *uint8
	u16  *uint16
	u32  *uint32
	u64  *uint64
	i8   *int8
	i16  *int16
	i32  *int32
	i64  *int64
	f32  *float32
	f64  *float64
	s    string
	i    *int
	u    *uint
	uptr *uintptr
}

func Parse(s string) (num Number) {
	num.s = s
	i, ok := ParseIntOk(s)
	if ok {
		num.i = util.New(i)
	}
	i8, ok := ParseInt8Ok(s)
	if ok {
		num.i8 = util.New(i8)
	}
	i16, ok := ParseInt16Ok(s)
	if ok {
		num.i16 = util.New(i16)
	}
	i32, ok := ParseInt32Ok(s)
	if ok {
		num.i32 = util.New(i32)
	}
	i64, ok := ParseInt64Ok(s)
	if ok {
		num.i64 = util.New(i64)
	}
	u, ok := ParseUintOk(s)
	if ok {
		num.u = util.New(u)
	}
	u8, ok := ParseUint8Ok(s)
	if ok {
		num.u8 = util.New(u8)
	}
	u16, ok := ParseUint16Ok(s)
	if ok {
		num.u16 = util.New(u16)
	}
	u32, ok := ParseUint32Ok(s)
	if ok {
		num.u32 = util.New(u32)
	}
	u64, ok := ParseUint64Ok(s)
	if ok {
		num.u64 = util.New(u64)
	}
	uptr, ok := ParseUintptrOk(s)
	if ok {
		num.uptr = util.New(uptr)
	}
	f32, ok := ParseFloat32Ok(s)
	if ok {
		num.f32 = util.New(f32)
	}
	f64, ok := ParseFloat64Ok(s)
	if ok {
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
