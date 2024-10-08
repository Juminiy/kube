package safe_cast

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
)

const (
	errCastNegative2Unsigned  = "cast negative-value to unsigned-type"
	errCastValueRangeOverflow = "cast wide-range-value to thin-range-type overflow"
)

func castErrorF(fromTyp, toTyp string, v any, err string) {
	stdlog.ErrorF("cast %s(%v) to %s, error: %s", fromTyp, v, toTyp, err)
}

func castNegativeErrorF(fromTyp, toTyp string, v any) {
	castErrorF(fromTyp, toTyp, v, errCastNegative2Unsigned)
}

func castOverflowErrorF(fromTyp, toTyp string, v any) {
	castErrorF(fromTyp, toTyp, v, errCastValueRangeOverflow)
}

const (
	InvalidI    int     = 0
	InvalidI8   int8    = 0
	InvalidI16  int16   = 0
	InvalidI32  int32   = 0
	InvalidI64  int64   = 0
	InvalidU    uint    = 0
	InvalidU8   uint8   = 0
	InvalidU16  uint16  = 0
	InvalidU32  uint32  = 0
	InvalidU64  uint64  = 0
	InvalidUPtr uintptr = 0
	InvalidF32  float32 = 0
	InvalidF64  float64 = 0
)

//go:generate go run codegen/safe_cast_codegen.go
//go:generate gofmt -w -s .
