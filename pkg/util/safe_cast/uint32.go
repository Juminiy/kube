// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

import "math"

func U32toI8(u32 uint32) int8 {
	if u32 > math.MaxInt8 {
		castOverflowErrorF("uint32", "int8", u32)
		return InvalidI8
	}
	return int8(u32)
}

func U32toI64(u32 uint32) int64 {
	return int64(u32)
}

func U32toU8(u32 uint32) uint8 {
	if u32 > math.MaxUint8 {
		castOverflowErrorF("uint32", "uint8", u32)
		return InvalidU8
	}
	return uint8(u32)
}

func U32toU64(u32 uint32) uint64 {
	return uint64(u32)
}

func U32toI(u32 uint32) int {
	return int(u32)
}

func U32toI16(u32 uint32) int16 {
	if u32 > math.MaxInt16 {
		castOverflowErrorF("uint32", "int16", u32)
		return InvalidI16
	}
	return int16(u32)
}

func U32toI32(u32 uint32) int32 {
	if u32 > math.MaxInt32 {
		castOverflowErrorF("uint32", "int32", u32)
		return InvalidI32
	}
	return int32(u32)
}

func U32toU(u32 uint32) uint {
	return uint(u32)
}

func U32toU16(u32 uint32) uint16 {
	if u32 > math.MaxUint16 {
		castOverflowErrorF("uint32", "uint16", u32)
		return InvalidU16
	}
	return uint16(u32)
}

func U32toU32(u32 uint32) uint32 {
	return uint32(u32)
}

func U32toUPtr(u32 uint32) uintptr {
	return uintptr(u32)
}
