// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

import "math"

func UPtrtoI8(uptr uintptr) int8 {
	if uptr > math.MaxInt8 {
		castOverflowErrorF("uintptr", "int8", uptr)
		return InvalidI8
	}
	return int8(uptr)
}

func UPtrtoI64(uptr uintptr) int64 {
	if uptr > math.MaxInt64 {
		castOverflowErrorF("uintptr", "int64", uptr)
		return InvalidI64
	}
	return int64(uptr)
}

func UPtrtoU8(uptr uintptr) uint8 {
	if uptr > math.MaxUint8 {
		castOverflowErrorF("uintptr", "uint8", uptr)
		return InvalidU8
	}
	return uint8(uptr)
}

func UPtrtoU64(uptr uintptr) uint64 {
	return uint64(uptr)
}

func UPtrtoU32(uptr uintptr) uint32 {
	if uptr > math.MaxUint32 {
		castOverflowErrorF("uintptr", "uint32", uptr)
		return InvalidU32
	}
	return uint32(uptr)
}

func UPtrtoUPtr(uptr uintptr) uintptr {
	return uintptr(uptr)
}

func UPtrtoI(uptr uintptr) int {
	if uptr > math.MaxInt {
		castOverflowErrorF("uintptr", "int", uptr)
		return InvalidI
	}
	return int(uptr)
}

func UPtrtoI16(uptr uintptr) int16 {
	if uptr > math.MaxInt16 {
		castOverflowErrorF("uintptr", "int16", uptr)
		return InvalidI16
	}
	return int16(uptr)
}

func UPtrtoI32(uptr uintptr) int32 {
	if uptr > math.MaxInt32 {
		castOverflowErrorF("uintptr", "int32", uptr)
		return InvalidI32
	}
	return int32(uptr)
}

func UPtrtoU(uptr uintptr) uint {
	return uint(uptr)
}

func UPtrtoU16(uptr uintptr) uint16 {
	if uptr > math.MaxUint16 {
		castOverflowErrorF("uintptr", "uint16", uptr)
		return InvalidU16
	}
	return uint16(uptr)
}
