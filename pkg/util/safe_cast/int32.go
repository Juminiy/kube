// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

import "math"

func I32toU8(i32 int32) uint8 {
	if i32 < 0 {
		castNegativeErrorF("int32", "uint8", i32)
		return InvalidU8
	}
	if i32 > math.MaxUint8 {
		castOverflowErrorF("int32", "uint8", i32)
		return InvalidU8
	}
	return uint8(i32)
}

func I32toU64(i32 int32) uint64 {
	if i32 < 0 {
		castNegativeErrorF("int32", "uint64", i32)
		return InvalidU64
	}
	return uint64(i32)
}

func I32toI8(i32 int32) int8 {
	if i32 > math.MaxInt8 {
		castOverflowErrorF("int32", "int8", i32)
		return InvalidI8
	}
	if i32 < math.MinInt8 {
		castOverflowErrorF("int32", "int8", i32)
		return InvalidI8
	}
	return int8(i32)
}

func I32toI64(i32 int32) int64 {
	return int64(i32)
}

func I32toI32(i32 int32) int32 {
	return int32(i32)
}

func I32toU(i32 int32) uint {
	if i32 < 0 {
		castNegativeErrorF("int32", "uint", i32)
		return InvalidU
	}
	return uint(i32)
}

func I32toU16(i32 int32) uint16 {
	if i32 < 0 {
		castNegativeErrorF("int32", "uint16", i32)
		return InvalidU16
	}
	if i32 > math.MaxUint16 {
		castOverflowErrorF("int32", "uint16", i32)
		return InvalidU16
	}
	return uint16(i32)
}

func I32toU32(i32 int32) uint32 {
	if i32 < 0 {
		castNegativeErrorF("int32", "uint32", i32)
		return InvalidU32
	}
	return uint32(i32)
}

func I32toUPtr(i32 int32) uintptr {
	if i32 < 0 {
		castNegativeErrorF("int32", "uintptr", i32)
		return InvalidUPtr
	}
	return uintptr(i32)
}

func I32toI(i32 int32) int {
	return int(i32)
}

func I32toI16(i32 int32) int16 {
	if i32 > math.MaxInt16 {
		castOverflowErrorF("int32", "int16", i32)
		return InvalidI16
	}
	if i32 < math.MinInt16 {
		castOverflowErrorF("int32", "int16", i32)
		return InvalidI16
	}
	return int16(i32)
}
