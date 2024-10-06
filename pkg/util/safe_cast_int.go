package util

import (
	"math"
)

const (
	InvalidU64 uint64 = 0
	InvalidI   int    = 0
)

const (
	errCastNegative2Unsigned  = "cast negative-value to unsigned-type"
	errCastValueRangeOverflow = "cast wide-range-value to thin-range-type overflow"
)

func castNegativeErrorF(fromTyp, toTyp string, v any) {
	castErrorF(fromTyp, toTyp, v, errCastNegative2Unsigned)
}

func castOverflowErrorF(fromTyp, toTyp string, v any) {
	castErrorF(fromTyp, toTyp, v, errCastValueRangeOverflow)
}

func ItoU64(i int) uint64 {
	if i < 0 {
		castNegativeErrorF("int", "uint64", i)
		return InvalidU64
	}
	return uint64(i)
}

func U64toI(u64 uint64) int {
	if u64 > math.MaxInt {
		castOverflowErrorF("uint64", "int", u64)
		return InvalidI
	}
	return int(u64)
}
