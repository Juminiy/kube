package safe_parse

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_cast"
	"strconv"
	"time"
)

func ParseBool(s string) (v bool, ok bool) {
	v, err := strconv.ParseBool(s)
	ok = err == nil
	if !ok {
		return
	}
	return
}

func ParseInt(s string) (v int, ok bool) {
	i64, err := strconv.ParseInt(s, 10, 64)
	ok = err == nil && util.InRange(i64, int64(util.MinInt), int64(util.MaxInt))
	if !ok {
		return
	}
	return safe_cast.I64toI(i64), ok
}

func ParseInt8(s string) (v int8, ok bool) {
	i64, err := strconv.ParseInt(s, 10, 64)
	ok = err == nil && util.InRange(i64, int64(util.MinInt8), int64(util.MaxInt8))
	if !ok {
		return
	}
	return safe_cast.I64toI8(i64), ok
}

func ParseInt16(s string) (v int16, ok bool) {
	i64, err := strconv.ParseInt(s, 10, 64)
	ok = err == nil && util.InRange(i64, int64(util.MinInt16), int64(util.MaxInt16))
	if !ok {
		return
	}
	return safe_cast.I64toI16(i64), ok
}

func ParseInt32(s string) (v int32, ok bool) {
	i64, err := strconv.ParseInt(s, 10, 64)
	ok = err == nil && util.InRange(i64, int64(util.MinInt32), int64(util.MaxInt32))
	if !ok {
		return
	}
	return safe_cast.I64toI32(i64), ok
}

func ParseInt64(s string) (v int64, ok bool) {
	i64, err := strconv.ParseInt(s, 10, 64)
	ok = err == nil && util.InRange(i64, util.MinInt64, util.MaxInt64)
	if !ok {
		return
	}
	return safe_cast.I64toI64(i64), ok
}

func ParseUint(s string) (v uint, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, uint64(util.MinUint), uint64(util.MaxUint))
	if !ok {
		return
	}
	return safe_cast.U64toU(u64), ok
}

func ParseUint8(s string) (v uint8, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, uint64(util.MinUint8), uint64(util.MaxUint8))
	if !ok {
		return
	}
	return safe_cast.U64toU8(u64), ok
}

func ParseUint16(s string) (v uint16, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, uint64(util.MinUint16), uint64(util.MaxUint16))
	if !ok {
		return
	}
	return safe_cast.U64toU16(u64), ok
}

func ParseUint32(s string) (v uint32, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, uint64(util.MinUint32), uint64(util.MaxUint32))
	if !ok {
		return
	}
	return safe_cast.U64toU32(u64), ok
}

func ParseUint64(s string) (v uint64, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, util.MinUint64, util.MaxUint64)
	if !ok {
		return
	}
	return safe_cast.U64toU64(u64), ok
}

func ParseUintptr(s string) (v uintptr, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, util.MinUint64, util.MaxUint64)
	if !ok {
		return
	}
	return safe_cast.U64toUPtr(u64), ok
}

func ParseFloat32(s string) (v float32, ok bool) {
	f64, err := strconv.ParseFloat(s, 64)
	ok = err == nil && f64InRangeF32(f64)
	if !ok {
		return
	}
	return safe_cast.F64tof32(f64), ok
}

func ParseFloat64(s string) (v float64, ok bool) {
	f64, err := strconv.ParseFloat(s, 64)
	return f64, err == nil
}

func f64InRangeF32(v float64) bool {
	return util.InRange(v, util.MinFloat32Overflow, util.MaxFloat32Overflow)
}

func ParseComplex64(s string) (v complex64, ok bool) {
	c128, err := strconv.ParseComplex(s, 128)
	ok = err == nil && f64InRangeF32(real(c128)) && f64InRangeF32(imag(c128))
	if !ok {
		return
	}
	return complex(safe_cast.F64tof32(real(c128)), safe_cast.F64tof32(imag(c128))), ok
}

func ParseComplex(s string) (v complex128, ok bool) {
	c128, err := strconv.ParseComplex(s, 128)
	return c128, err == nil
}

func ParseTime(s string) (v time.Time, ok bool) {
	v, err := time.Parse(time.DateTime, s)
	return v, err == nil
}
