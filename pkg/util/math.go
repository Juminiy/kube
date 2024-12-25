package util

import (
	"math"
)

// Typed Float
const (
	MaxFloat32             = float32(math.MaxFloat32)
	MinFloat32             = float32(-math.MaxFloat32)
	SmallestNonzeroFloat32 = float32(math.SmallestNonzeroFloat32)
	MaxFloat32Overflow     = float64(math.MaxFloat32)
	MinFloat32Overflow     = float64(-math.MaxFloat32)
	MaxFloat64             = float64(math.MaxFloat64)
	MinFloat64             = float64(-math.MaxFloat64)
	SmallestNonzeroFloat64 = float64(math.SmallestNonzeroFloat64)
)

// Typed Int, Uint
const (
	MaxInt     = int(math.MaxInt)
	MinInt     = int(math.MinInt)
	MaxInt8    = int8(math.MaxInt8)
	MinInt8    = int8(math.MinInt8)
	MaxInt16   = int16(math.MaxInt16)
	MinInt16   = int16(math.MinInt16)
	MaxInt32   = int32(math.MaxInt32)
	MinInt32   = int32(math.MinInt32)
	MaxInt64   = int64(math.MaxInt64)
	MinInt64   = int64(math.MinInt64)
	MaxUint    = uint(math.MaxUint)
	MinUint    = uint(0)
	MaxUint8   = uint8(math.MaxUint8)
	MinUint8   = uint8(0)
	MaxUint16  = uint16(math.MaxUint16)
	MinUint16  = uint16(0)
	MaxUint32  = uint32(math.MaxUint32)
	MinUint32  = uint32(0)
	MaxUint64  = uint64(math.MaxUint64)
	MinUint64  = uint64(0)
	MaxUintptr = uintptr(math.MaxUint)
	MinUintptr = uintptr(0)
)

// Typed InRange
func InRange[T Number](v, l, r T) bool {
	return l <= v && v <= r
}

func IsOdd(n int) bool {
	return n%2 == 1
}

const (
	MagicNumber = 114514
)
