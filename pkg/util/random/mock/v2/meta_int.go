// Package mockv2 was generated
package mockv2

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/spf13/cast"
	"math"
)

type IntFunc func() int64

var intFunc = map[string]IntFunc{
	defaultKey: defaultInt,
}

var defaultInt = func() int64 {
	return int64(gofakeit.IntN(intDefault))
}

var defaultI8 = func() int64 {
	return int64(gofakeit.IntRange(math.MinInt8, math.MaxInt8))
}

const (
	intDefault = 1<<15 - 1
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

func (r *rule) applyInt(minval, maxval string) {
	(*r)["int:min"], (*r)["int:max"] = rangeOfInt64(minval, maxval, math.MinInt, math.MaxInt, math.MaxInt)
	(*r)["i8:min"], (*r)["i8:max"] = rangeOfInt64(minval, maxval, math.MinInt8, math.MaxInt8, math.MaxInt8)
	(*r)["i16:min"], (*r)["i16:max"] = rangeOfInt64(minval, maxval, math.MinInt16, math.MaxInt16, math.MaxInt16)
	(*r)["i32:min"], (*r)["i32:max"] = rangeOfInt64(minval, maxval, math.MinInt32, math.MaxInt32, math.MaxInt32)
	(*r)["i64:min"], (*r)["i64:max"] = rangeOfInt64(minval, maxval, math.MinInt64, math.MaxInt64, math.MaxInt64)
}

// get a valid pair range value
// from a given range: minof, maxof
// and default range: mindef, maxdef
func rangeOfInt64(minof, maxof string, mindef, maxdef, maxoob int64) (ofmin, ofmax int64) {
	ofmin, ofmax = mindef, maxdef
	minnil, maxnil := len(minof) == 0, len(maxof) == 0
	if !minnil {
		ofmin = cast.ToInt64(minof)
	}
	if !maxnil {
		ofmax = cast.ToInt64(maxof)
	}
	if ofmin < mindef || ofmax > maxdef { // invalid tag value
		return
	}

	switch {
	case !minnil && !maxnil && ofmax < ofmin: // invalid tag value
		return

	case !minnil && maxnil:

	}

	if ofmax > maxoob {
		ofmax = maxoob
	}
	return
}
