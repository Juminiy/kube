package mock

import (
	"github.com/brianvoe/gofakeit/v7"
	"math"
)

type UintFunc func() uint64

var uintFunc = map[string]UintFunc{}

var defaultUint = func() uint64 {
	return uint64(gofakeit.UintN(intDefault))
}

// type decl is all int
var uintRule = rule{
	"uint:min": uint64(0),
	"uint:max": uint64(math.MaxUint),
	"u8:min":   uint64(0),
	"u8:max":   uint64(math.MaxUint8),
	"u16:min":  uint64(0),
	"u16:max":  uint64(math.MaxUint16),
	"u32:min":  uint64(0),
	"u32:max":  uint64(math.MaxUint32),
	"u64:min":  uint64(0),
	"u64:max":  uint64(math.MaxUint64),
}
