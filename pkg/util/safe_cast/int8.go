// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.
package safe_cast

func I8toI(i8 int8) int {
	return int(i8)
}

func I8toI16(i8 int8) int16 {
	return int16(i8)
}

func I8toI32(i8 int8) int32 {
	return int32(i8)
}

func I8toU(i8 int8) uint {
	if i8 < 0 {
		castNegativeErrorF("int8", "uint", i8)
		return InvalidU
	}
	return uint(i8)
}

func I8toU16(i8 int8) uint16 {
	if i8 < 0 {
		castNegativeErrorF("int8", "uint16", i8)
		return InvalidU16
	}
	return uint16(i8)
}

func I8toU32(i8 int8) uint32 {
	if i8 < 0 {
		castNegativeErrorF("int8", "uint32", i8)
		return InvalidU32
	}
	return uint32(i8)
}

func I8toUPtr(i8 int8) uintptr {
	if i8 < 0 {
		castNegativeErrorF("int8", "uintptr", i8)
		return InvalidUPtr
	}
	return uintptr(i8)
}

func I8toI8(i8 int8) int8 {
	return int8(i8)
}

func I8toI64(i8 int8) int64 {
	return int64(i8)
}

func I8toU8(i8 int8) uint8 {
	if i8 < 0 {
		castNegativeErrorF("int8", "uint8", i8)
		return InvalidU8
	}
	return uint8(i8)
}

func I8toU64(i8 int8) uint64 {
	if i8 < 0 {
		castNegativeErrorF("int8", "uint64", i8)
		return InvalidU64
	}
	return uint64(i8)
}
