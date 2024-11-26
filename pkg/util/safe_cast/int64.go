// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

import (
	"github.com/spf13/cast"
	"math"
)

func I64toI32(i64 int64) int32 {
	if i64 > math.MaxInt32 {
		castOverflowErrorF("int64", "int32", i64)
		return InvalidI32
	}
	if i64 < math.MinInt32 {
		castOverflowErrorF("int64", "int32", i64)
		return InvalidI32
	}
	return int32(i64)
}

func I64toU(i64 int64) uint {
	if i64 < 0 {
		castNegativeErrorF("int64", "uint", i64)
		return InvalidU
	}
	return uint(i64)
}

func I64toU16(i64 int64) uint16 {
	if i64 < 0 {
		castNegativeErrorF("int64", "uint16", i64)
		return InvalidU16
	}
	if i64 > math.MaxUint16 {
		castOverflowErrorF("int64", "uint16", i64)
		return InvalidU16
	}
	return uint16(i64)
}

func I64toU32(i64 int64) uint32 {
	if i64 < 0 {
		castNegativeErrorF("int64", "uint32", i64)
		return InvalidU32
	}
	if i64 > math.MaxUint32 {
		castOverflowErrorF("int64", "uint32", i64)
		return InvalidU32
	}
	return uint32(i64)
}

func I64toUPtr(i64 int64) uintptr {
	if i64 < 0 {
		castNegativeErrorF("int64", "uintptr", i64)
		return InvalidUPtr
	}
	return uintptr(i64)
}

func I64toI(i64 int64) int {
	return int(i64)
}

func I64toI16(i64 int64) int16 {
	if i64 > math.MaxInt16 {
		castOverflowErrorF("int64", "int16", i64)
		return InvalidI16
	}
	if i64 < math.MinInt16 {
		castOverflowErrorF("int64", "int16", i64)
		return InvalidI16
	}
	return int16(i64)
}

func I64toU8(i64 int64) uint8 {
	if i64 < 0 {
		castNegativeErrorF("int64", "uint8", i64)
		return InvalidU8
	}
	if i64 > math.MaxUint8 {
		castOverflowErrorF("int64", "uint8", i64)
		return InvalidU8
	}
	return uint8(i64)
}

func I64toU64(i64 int64) uint64 {
	if i64 < 0 {
		castNegativeErrorF("int64", "uint64", i64)
		return InvalidU64
	}
	return uint64(i64)
}

func I64toI8(i64 int64) int8 {
	if i64 > math.MaxInt8 {
		castOverflowErrorF("int64", "int8", i64)
		return InvalidI8
	}
	if i64 < math.MinInt8 {
		castOverflowErrorF("int64", "int8", i64)
		return InvalidI8
	}
	return int8(i64)
}

func I64toI64(i64 int64) int64 {
	return int64(i64)
}

func ItoF32[I ~int | int8 | int16 | int32 | int64](i I) float32 {
	return cast.ToFloat32(cast.ToString(i))
}

func ItoF64[I ~int | int8 | int16 | int32 | int64](i I) float64 {
	return cast.ToFloat64(cast.ToString(i))
}
