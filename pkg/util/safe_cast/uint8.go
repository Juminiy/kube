// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

import "math"

func U8toI32(u8 uint8) int32 {
	return int32(u8)
}

func U8toU(u8 uint8) uint {
	return uint(u8)
}

func U8toU16(u8 uint8) uint16 {
	return uint16(u8)
}

func U8toU32(u8 uint8) uint32 {
	return uint32(u8)
}

func U8toUPtr(u8 uint8) uintptr {
	return uintptr(u8)
}

func U8toI(u8 uint8) int {
	return int(u8)
}

func U8toI16(u8 uint8) int16 {
	return int16(u8)
}

func U8toU8(u8 uint8) uint8 {
	return uint8(u8)
}

func U8toU64(u8 uint8) uint64 {
	return uint64(u8)
}

func U8toI8(u8 uint8) int8 {
	if u8 > math.MaxInt8 {
		castOverflowErrorF("uint8", "int8", u8)
		return InvalidI8
	}
	return int8(u8)
}

func U8toI64(u8 uint8) int64 {
	return int64(u8)
}
