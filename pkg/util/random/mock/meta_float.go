package mock

import (
	"github.com/brianvoe/gofakeit/v7"
	"math"
)

type FloatFunc func() float64

var floatFunc = map[string]FloatFunc{}

var defaultFloat = func() float64 {
	return gofakeit.Float64Range(floatDefaultMin, floatDefaultMax)
}

const (
	floatDefaultMin = float64(-intDefault)
	floatDefaultMax = float64(intDefault)
)

// type decl is all float64
var floatRule = rule{
	"float32:min": float64(math.SmallestNonzeroFloat32),
	"float32:max": float64(math.MaxFloat32),
	"float64:min": float64(math.SmallestNonzeroFloat64),
	"float64:max": float64(math.MaxFloat64),
}
