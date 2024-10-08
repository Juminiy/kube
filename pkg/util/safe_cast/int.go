// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

import "math"

func ItoUPtr(i int) uintptr {
	if i < 0 {
		castNegativeErrorF("int", "uintptr", i)
		return InvalidUPtr
	}
	return uintptr(i)
}

func ItoI(i int) int {
	return int(i)
}

func ItoI16(i int) int16 {
	if i > math.MaxInt16 {
		castOverflowErrorF("int", "int16", i)
		return InvalidI16
	}
	if i < math.MinInt16 {
		castOverflowErrorF("int", "int16", i)
		return InvalidI16
	}
	return int16(i)
}

func ItoI32(i int) int32 {
	if i > math.MaxInt32 {
		castOverflowErrorF("int", "int32", i)
		return InvalidI32
	}
	if i < math.MinInt32 {
		castOverflowErrorF("int", "int32", i)
		return InvalidI32
	}
	return int32(i)
}

func ItoU(i int) uint {
	if i < 0 {
		castNegativeErrorF("int", "uint", i)
		return InvalidU
	}
	return uint(i)
}

func ItoU16(i int) uint16 {
	if i < 0 {
		castNegativeErrorF("int", "uint16", i)
		return InvalidU16
	}
	if i > math.MaxUint16 {
		castOverflowErrorF("int", "uint16", i)
		return InvalidU16
	}
	return uint16(i)
}

func ItoU32(i int) uint32 {
	if i < 0 {
		castNegativeErrorF("int", "uint32", i)
		return InvalidU32
	}
	if i > math.MaxUint32 {
		castOverflowErrorF("int", "uint32", i)
		return InvalidU32
	}
	return uint32(i)
}

func ItoI8(i int) int8 {
	if i > math.MaxInt8 {
		castOverflowErrorF("int", "int8", i)
		return InvalidI8
	}
	if i < math.MinInt8 {
		castOverflowErrorF("int", "int8", i)
		return InvalidI8
	}
	return int8(i)
}

func ItoI64(i int) int64 {
	return int64(i)
}

func ItoU8(i int) uint8 {
	if i < 0 {
		castNegativeErrorF("int", "uint8", i)
		return InvalidU8
	}
	if i > math.MaxUint8 {
		castOverflowErrorF("int", "uint8", i)
		return InvalidU8
	}
	return uint8(i)
}

func ItoU64(i int) uint64 {
	if i < 0 {
		castNegativeErrorF("int", "uint64", i)
		return InvalidU64
	}
	return uint64(i)
}
