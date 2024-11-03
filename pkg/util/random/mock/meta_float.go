package mock

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/spf13/cast"
	"math"
)

type FloatFunc func() float64

var floatFunc = map[string]FloatFunc{
	defaultKey: defaultFloat,
}

var defaultFloat = func() float64 {
	return gofakeit.Float64Range(floatDefaultMin, floatDefaultMax)
}

const (
	floatDefaultMin = float64(-intDefault)
	floatDefaultMax = float64(intDefault)
)

// type decl is all float64
var floatRule = rule{
	"float32:min": math.SmallestNonzeroFloat32,
	"float32:max": math.MaxFloat32,
	"float64:min": math.SmallestNonzeroFloat64,
	"float64:max": math.MaxFloat64,
}

func (r *rule) applyFloat(minval, maxval string) {
	(*r)["float32:min"], (*r)["float32:max"] = rangeOfFloat64(minval, maxval, math.SmallestNonzeroFloat32, math.MaxFloat32, math.MaxFloat32)
	(*r)["float64:min"], (*r)["float64:max"] = rangeOfFloat64(minval, maxval, math.SmallestNonzeroFloat64, math.MaxFloat64, math.MaxFloat64)
}

// get a valid pair range value
// from a given range: minof, maxof
// and default range: mindef, maxdef
func rangeOfFloat64(minof, maxof string, mindef, maxdef, maxoob float64) (ofmin, ofmax float64) {
	ofmin, ofmax = mindef, maxdef
	minnil, maxnil := len(minof) == 0, len(maxof) == 0
	if !minnil {
		ofmin = cast.ToFloat64(minof)
	}
	if !maxnil {
		ofmax = cast.ToFloat64(maxof)
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
