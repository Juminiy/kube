// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

import "math"

func I16toI64(i16 int16) int64 {
	return int64(i16)
}

func I16toU8(i16 int16) uint8 {
	if i16 < 0 {
		castNegativeErrorF("int16", "uint8", i16)
		return InvalidU8
	}
	if i16 > math.MaxUint8 {
		castOverflowErrorF("int16", "uint8", i16)
		return InvalidU8
	}
	return uint8(i16)
}

func I16toU64(i16 int16) uint64 {
	if i16 < 0 {
		castNegativeErrorF("int16", "uint64", i16)
		return InvalidU64
	}
	return uint64(i16)
}

func I16toI8(i16 int16) int8 {
	if i16 > math.MaxInt8 {
		castOverflowErrorF("int16", "int8", i16)
		return InvalidI8
	}
	if i16 < math.MinInt8 {
		castOverflowErrorF("int16", "int8", i16)
		return InvalidI8
	}
	return int8(i16)
}

func I16toI16(i16 int16) int16 {
	return int16(i16)
}

func I16toI32(i16 int16) int32 {
	return int32(i16)
}

func I16toU(i16 int16) uint {
	if i16 < 0 {
		castNegativeErrorF("int16", "uint", i16)
		return InvalidU
	}
	return uint(i16)
}

func I16toU16(i16 int16) uint16 {
	if i16 < 0 {
		castNegativeErrorF("int16", "uint16", i16)
		return InvalidU16
	}
	return uint16(i16)
}

func I16toU32(i16 int16) uint32 {
	if i16 < 0 {
		castNegativeErrorF("int16", "uint32", i16)
		return InvalidU32
	}
	return uint32(i16)
}

func I16toUPtr(i16 int16) uintptr {
	if i16 < 0 {
		castNegativeErrorF("int16", "uintptr", i16)
		return InvalidUPtr
	}
	return uintptr(i16)
}

func I16toI(i16 int16) int {
	return int(i16)
}
