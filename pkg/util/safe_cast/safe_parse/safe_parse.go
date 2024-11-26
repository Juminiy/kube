package safe_parse

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_cast"
	"strconv"
)

func ParseBoolOk(s string) (v bool, ok bool) {
	v, err := strconv.ParseBool(s)
	ok = err == nil
	return
}

func ParseIntOk(s string) (v int, ok bool) {
	i64, err := strconv.ParseInt(s, 10, 64)
	ok = err == nil && util.InRange(i64, int64(util.MinInt), int64(util.MaxInt))
	return safe_cast.I64toI(i64), ok
}

func ParseInt8Ok(s string) (v int8, ok bool) {
	i64, err := strconv.ParseInt(s, 10, 64)
	ok = err == nil && util.InRange(i64, int64(util.MinInt8), int64(util.MaxInt8))
	return safe_cast.I64toI8(i64), ok
}

func ParseInt16Ok(s string) (v int16, ok bool) {
	i64, err := strconv.ParseInt(s, 10, 64)
	ok = err == nil && util.InRange(i64, int64(util.MinInt16), int64(util.MaxInt16))
	return safe_cast.I64toI16(i64), ok
}

func ParseInt32Ok(s string) (v int32, ok bool) {
	i64, err := strconv.ParseInt(s, 10, 64)
	ok = err == nil && util.InRange(i64, int64(util.MinInt32), int64(util.MaxInt32))
	return safe_cast.I64toI32(i64), ok
}

func ParseInt64Ok(s string) (v int64, ok bool) {
	i64, err := strconv.ParseInt(s, 10, 64)
	ok = err == nil && util.InRange(i64, util.MinInt64, util.MaxInt64)
	return safe_cast.I64toI64(i64), ok
}

func ParseUintOk(s string) (v uint, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, uint64(util.MinUint), uint64(util.MaxUint))
	return safe_cast.U64toU(u64), ok
}

func ParseUint8Ok(s string) (v uint8, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, uint64(util.MinUint8), uint64(util.MaxUint8))
	return safe_cast.U64toU8(u64), ok
}

func ParseUint16Ok(s string) (v uint16, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, uint64(util.MinUint16), uint64(util.MaxUint16))
	return safe_cast.U64toU16(u64), ok
}

func ParseUint32Ok(s string) (v uint32, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, uint64(util.MinUint32), uint64(util.MaxUint32))
	return safe_cast.U64toU32(u64), ok
}

func ParseUint64Ok(s string) (v uint64, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, util.MinUint64, util.MaxUint64)
	return safe_cast.U64toU64(u64), ok
}

func ParseUintptrOk(s string) (v uintptr, ok bool) {
	u64, err := strconv.ParseUint(s, 10, 64)
	ok = err == nil && util.InRange(u64, util.MinUint64, util.MaxUint64)
	return safe_cast.U64toUPtr(u64), ok
}

func ParseFloat32Ok(s string) (v float32, ok bool) {
	f64, err := strconv.ParseFloat(s, 64)
	ok = err == nil && util.InRange(f64, float64(util.SmallestNonzeroFloat32), float64(util.MaxFloat32))
	return safe_cast.F64tof32(f64), ok
}

func ParseFloat64Ok(s string) (v float64, ok bool) {
	f64, err := strconv.ParseFloat(s, 64)
	ok = err == nil && util.InRange(f64, util.SmallestNonzeroFloat64, util.MaxFloat64)
	return f64, ok
}
