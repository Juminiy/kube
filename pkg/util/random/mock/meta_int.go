package mock

import (
	"github.com/brianvoe/gofakeit/v7"
	"math"
)

type IntFunc func() int64

var intFunc = map[string]IntFunc{}

var defaultInt = func() int64 {
	return int64(gofakeit.IntN(intDefault))
}

const (
	intDefault = 1 << 16
)

// type decl is all int
var intRule = rule{
	"int:min": int64(math.MinInt),
	"int:max": int64(math.MaxInt),
	"i8:min":  int64(math.MinInt8),
	"i8:max":  int64(math.MaxInt8),
	"i16:min": int64(math.MinInt16),
	"i16:max": int64(math.MaxInt16),
	"i32:min": int64(math.MinInt32),
	"i32:max": int64(math.MaxInt32),
	"i64:min": int64(math.MinInt64),
	"i64:max": int64(math.MaxInt64),
}
