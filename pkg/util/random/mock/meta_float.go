package mock

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/spf13/cast"
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
	"float32:min": util.MinFloat32,
	"float32:max": util.MaxFloat32,
	"float64:min": util.MinFloat64,
	"float64:max": util.MaxFloat64,
}

func (r *rule) applyFloat(minval, maxval string) {
	(*r)["float32:min"], (*r)["float32:max"] = rangeOfFloat64(minval, maxval, util.MinFloat32Overflow, util.MaxFloat32Overflow, util.MaxFloat32Overflow)
	(*r)["float64:min"], (*r)["float64:max"] = rangeOfFloat64(minval, maxval, util.MinFloat64, util.MaxFloat64, util.MaxFloat64)
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
