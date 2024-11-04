package mock

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/Juminiy/kube/pkg/util/zero_reflect"
	"github.com/spf13/cast"
	"reflect"
	"strings"
	"time"
)

// struct app tag
const (
	mockTag = "mock"
)

type tKind reflect.Kind

// alias of reflect.Kind
const (
	tInvalid tKind = iota
	tBool
	tInt
	tI8
	tI16
	tI32
	tI64
	tUint
	tU8
	tU16
	tU32
	tU64
	tUPtr
	tF32
	tF64
	tC64
	tC128
	tArr
	tChan
	tFunc
	tAny
	tMap
	tPtr
	tSlice
	tString
	tStruct
	tUnsafePtr
)

// define
const (
	tTime tKind = iota + 27
)

func (k tKind) zero() any {
	switch k {
	case tBool:
		return false
	case tInt:
		return int(0)
	case tI8:
		return int8(0)
	case tI16:
		return int16(0)
	case tI32:
		return int32(0)
	case tI64:
		return int64(0)
	case tUint:
		return uint(0)
	case tU8:
		return uint8(0)
	case tU16:
		return uint16(0)
	case tU32:
		return uint32(0)
	case tU64:
		return uint64(0)
	case tUPtr:
		return uintptr(0)
	case tF32:
		return float32(0)
	case tF64:
		return float64(0)
	default:
		return ""
	}
}

func (k tKind) cast(v any) any {
	switch k {
	case tBool:
		return cast.ToBool(v)
	case tInt:
		return cast.ToInt(v)
	case tI8:
		return cast.ToInt8(v)
	case tI16:
		return cast.ToInt16(v)
	case tI32:
		return cast.ToInt32(v)
	case tI64:
		return cast.ToInt64(v)
	case tUint:
		return cast.ToUint(v)
	case tU8:
		return cast.ToUint8(v)
	case tU16:
		return cast.ToUint16(v)
	case tU32:
		return cast.ToUint32(v)
	case tU64:
		return cast.ToUint64(v)
	case tUPtr:
		return uintptr(cast.ToInt64(v))
	case tF32:
		return cast.ToFloat32(v)
	case tF64:
		return cast.ToFloat64(v)
	default:
		return cast.ToString(v)
	}
}

func (k tKind) isNum() bool {
	return util.ElemIn(k,
		tBool,
		tInt, tI8, tI16, tI32, tI64,
		tUint, tU8, tU16, tU32, tU64, tUPtr,
		tF32, tF64,
	)
}

func (k tKind) isMeta() bool {
	return util.ElemIn(k,
		tBool,
		tInt, tI8, tI16, tI32, tI64,
		tUint, tU8, tU16, tU32, tU64, tUPtr,
		tF32, tF64,
		tString,
	)
}

// short for safe_reflect.IndirectOf
var indir = safe_reflect.IndirectOf

var typeof = zero_reflect.TypeOf

var split = strings.Split

var _timeTyp = safe_reflect.Of(time.Now()).Typ

var pairToStr = func(v0, v1 any) (string, string) {
	return cast.ToString(v0), cast.ToString(v1)
}

var pairToInt = func(v0, v1 any) (int, int) {
	return cast.ToInt(v0), cast.ToInt(v1)
}

var pairToUInt = func(v0, v1 any) (uint, uint) {
	return cast.ToUint(v0), cast.ToUint(v1)
}

var pairToInt64 = func(v0, v1 any) (int64, int64) {
	return cast.ToInt64(v0), cast.ToInt64(v1)
}

var pairToUInt64 = func(v0, v1 any) (uint64, uint64) {
	return cast.ToUint64(v0), cast.ToUint64(v1)
}

var pairToF32 = func(v0, v1 any) (float32, float32) {
	return cast.ToFloat32(v0), cast.ToFloat32(v1)
}

var pairToF64 = func(v0, v1 any) (float64, float64) {
	return cast.ToFloat64(v0), cast.ToFloat64(v1)
}

func castFunc(src, dst any) any {
	switch dst.(type) {
	case uint64:
		return cast.ToUint64(src)
	case int64:
		return cast.ToInt64(src)
	case float64:
		return cast.ToFloat64(src)
	default:
		return cast.ToString(src)
	}
}
