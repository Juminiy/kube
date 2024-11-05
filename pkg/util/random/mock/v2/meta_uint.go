// Package mockv2 was generated
package mockv2

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/spf13/cast"
	"math"
)

type UintFunc func() uint64

var uintFunc = map[string]UintFunc{
	defaultKey: defaultUint,
}

var defaultUint = func() uint64 {
	return uint64(gofakeit.UintN(intDefault))
}

var defaultU8 = func() uint64 {
	return uint64(gofakeit.UintRange(0, math.MaxUint8))
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

func (r *rule) applyUint(minval, maxval string) {
	oldUMin, oldUMax := pairToUInt64((*r)["uint:min"], (*r)["uint:max"])
	(*r)["uint:min"], (*r)["uint:max"] = pairToUInt64(rangeOfUint64(minval, maxval, oldUMin, oldUMax, uint64(math.MaxUint)))

	oldU8Min, oldU8Max := pairToUInt64((*r)["uint:min"], (*r)["uint:max"])
	(*r)["u8:min"], (*r)["u8:max"] = pairToUInt64(rangeOfUint64(minval, maxval, oldU8Min, oldU8Max, uint64(math.MaxUint8)))

	oldU16Min, oldU16Max := pairToUInt64((*r)["uint:min"], (*r)["uint:max"])
	(*r)["u16:min"], (*r)["u16:max"] = pairToUInt64(rangeOfUint64(minval, maxval, oldU16Min, oldU16Max, uint64(math.MaxUint16)))

	oldU32Min, oldU32Max := pairToUInt64((*r)["uint:min"], (*r)["uint:max"])
	(*r)["u32:min"], (*r)["u32:max"] = pairToUInt64(rangeOfUint64(minval, maxval, oldU32Min, oldU32Max, uint64(math.MaxUint32)))

	oldU64Min, oldU64Max := pairToUInt64((*r)["uint:min"], (*r)["uint:max"])
	(*r)["u64:min"], (*r)["u64:max"] = pairToUInt64(rangeOfUint64(minval, maxval, oldU64Min, oldU64Max, uint64(math.MaxUint64)))
	r.applyInt(minval, maxval)
}

// get a valid pair range value
// from a given range: minof, maxof
// and default range: mindef, maxdef
func rangeOfUint64(minof, maxof string, mindef, maxdef, maxoob uint64) (ofmin, ofmax uint64) {
	ofmin, ofmax = mindef, maxdef
	minnil, maxnil := len(minof) == 0, len(maxof) == 0
	if !minnil {
		ofmin = cast.ToUint64(minof)
	}
	if !maxnil {
		ofmax = cast.ToUint64(maxof)
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
