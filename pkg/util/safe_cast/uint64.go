// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

import (
	"math"
)

func U64toI8(u64 uint64) int8 {
	if u64 > math.MaxInt8 {
		castOverflowErrorF("uint64", "int8", u64)
		return InvalidI8
	}
	return int8(u64)
}

func U64toI64(u64 uint64) int64 {
	if u64 > math.MaxInt64 {
		castOverflowErrorF("uint64", "int64", u64)
		return InvalidI64
	}
	return int64(u64)
}

func U64toU8(u64 uint64) uint8 {
	if u64 > math.MaxUint8 {
		castOverflowErrorF("uint64", "uint8", u64)
		return InvalidU8
	}
	return uint8(u64)
}

func U64toU64(u64 uint64) uint64 {
	return uint64(u64)
}

func U64toI(u64 uint64) int {
	if u64 > math.MaxInt {
		castOverflowErrorF("uint64", "int", u64)
		return InvalidI
	}
	return int(u64)
}

func U64toI16(u64 uint64) int16 {
	if u64 > math.MaxInt16 {
		castOverflowErrorF("uint64", "int16", u64)
		return InvalidI16
	}
	return int16(u64)
}

func U64toI32(u64 uint64) int32 {
	if u64 > math.MaxInt32 {
		castOverflowErrorF("uint64", "int32", u64)
		return InvalidI32
	}
	return int32(u64)
}

func U64toU(u64 uint64) uint {
	return uint(u64)
}

func U64toU16(u64 uint64) uint16 {
	if u64 > math.MaxUint16 {
		castOverflowErrorF("uint64", "uint16", u64)
		return InvalidU16
	}
	return uint16(u64)
}

func U64toU32(u64 uint64) uint32 {
	if u64 > math.MaxUint32 {
		castOverflowErrorF("uint64", "uint32", u64)
		return InvalidU32
	}
	return uint32(u64)
}

func U64toUPtr(u64 uint64) uintptr {
	return uintptr(u64)
}
