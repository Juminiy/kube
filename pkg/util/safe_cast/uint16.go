// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

import "math"

func U16toI8(u16 uint16) int8 {
	if u16 > math.MaxInt8 {
		castOverflowErrorF("uint16", "int8", u16)
		return InvalidI8
	}
	return int8(u16)
}

func U16toI64(u16 uint16) int64 {
	return int64(u16)
}

func U16toU8(u16 uint16) uint8 {
	if u16 > math.MaxUint8 {
		castOverflowErrorF("uint16", "uint8", u16)
		return InvalidU8
	}
	return uint8(u16)
}

func U16toU64(u16 uint16) uint64 {
	return uint64(u16)
}

func U16toI(u16 uint16) int {
	return int(u16)
}

func U16toI16(u16 uint16) int16 {
	if u16 > math.MaxInt16 {
		castOverflowErrorF("uint16", "int16", u16)
		return InvalidI16
	}
	return int16(u16)
}

func U16toI32(u16 uint16) int32 {
	return int32(u16)
}

func U16toU(u16 uint16) uint {
	return uint(u16)
}

func U16toU16(u16 uint16) uint16 {
	return uint16(u16)
}

func U16toU32(u16 uint16) uint32 {
	return uint32(u16)
}

func U16toUPtr(u16 uint16) uintptr {
	return uintptr(u16)
}
