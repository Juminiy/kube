// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

import "math"

func UtoI(u uint) int {
	if u > math.MaxInt {
		castOverflowErrorF("uint", "int", u)
		return InvalidI
	}
	return int(u)
}

func UtoI16(u uint) int16 {
	if u > math.MaxInt16 {
		castOverflowErrorF("uint", "int16", u)
		return InvalidI16
	}
	return int16(u)
}

func UtoI32(u uint) int32 {
	if u > math.MaxInt32 {
		castOverflowErrorF("uint", "int32", u)
		return InvalidI32
	}
	return int32(u)
}

func UtoU(u uint) uint {
	return uint(u)
}

func UtoU16(u uint) uint16 {
	if u > math.MaxUint16 {
		castOverflowErrorF("uint", "uint16", u)
		return InvalidU16
	}
	return uint16(u)
}

func UtoU32(u uint) uint32 {
	if u > math.MaxUint32 {
		castOverflowErrorF("uint", "uint32", u)
		return InvalidU32
	}
	return uint32(u)
}

func UtoUPtr(u uint) uintptr {
	return uintptr(u)
}

func UtoI8(u uint) int8 {
	if u > math.MaxInt8 {
		castOverflowErrorF("uint", "int8", u)
		return InvalidI8
	}
	return int8(u)
}

func UtoI64(u uint) int64 {
	if u > math.MaxInt64 {
		castOverflowErrorF("uint", "int64", u)
		return InvalidI64
	}
	return int64(u)
}

func UtoU8(u uint) uint8 {
	if u > math.MaxUint8 {
		castOverflowErrorF("uint", "uint8", u)
		return InvalidU8
	}
	return uint8(u)
}

func UtoU64(u uint) uint64 {
	return uint64(u)
}
