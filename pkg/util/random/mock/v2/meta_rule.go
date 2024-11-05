// Package mockv2 was generated
package mockv2

import (
	"github.com/brianvoe/gofakeit/v7"
	"math"
)

type rule map[string]any

/*
 * apply rule
 */
func (r *rule) applyRangeFns() []func(minval, maxval string) {
	return []func(string, string){
		r.applyInt,
		r.applyUint,
		r.applyFloat,
		r.applyStringLen,
		r.applyTime,
	}
}

func (r *rule) applyMin(minval int64) {
	minvalstr, maxvalstr := pairToStr(minval, math.MaxInt64)
	r.applyInt(minvalstr, maxvalstr)
	r.applyUint(minvalstr, maxvalstr)
	r.applyFloat(minvalstr, maxvalstr)
}

func (r *rule) applyMax(maxval int64) {
	minvalstr, maxvalstr := pairToStr(math.MinInt64, maxval)
	r.applyInt(minvalstr, maxvalstr)
	r.applyUint(minvalstr, maxvalstr)
	r.applyFloat(minvalstr, maxvalstr)
}

/*
 * setValue from rule
 */
func (r *rule) setValue(val map[tKind]any) {
	r.rangeValue(val)  // priority 2
	r.stringValue(val) // priority 1
	r.enumValue(val)   // priority 0
}

func (r *rule) rangeValue(val map[tKind]any) {
	for kind := tInt; kind <= tTime; kind++ {
		var v any
		switch kind {
		case tInt:
			v = gofakeit.IntRange(pairToInt((*r)["int:min"], (*r)["int:max"]))
		case tI8:
			v = gofakeit.IntRange(pairToInt((*r)["i8:min"], (*r)["i8:max"]))
		case tI16:
			v = gofakeit.IntRange(pairToInt((*r)["i16:min"], (*r)["i16:max"]))
		case tI32:
			v = gofakeit.IntRange(pairToInt((*r)["i32:min"], (*r)["i32:max"]))
		case tI64:
			v = gofakeit.IntRange(pairToInt((*r)["i64:min"], (*r)["i64:max"]))
		case tUint:
			v = gofakeit.UintRange(pairToUInt((*r)["uint:min"], (*r)["uint:max"]))
		case tU8:
			v = gofakeit.UintRange(pairToUInt((*r)["u8:min"], (*r)["u8:max"]))
		case tU16:
			v = gofakeit.UintRange(pairToUInt((*r)["u16:min"], (*r)["u16:max"]))
		case tU32:
			v = gofakeit.UintRange(pairToUInt((*r)["u32:min"], (*r)["u32:max"]))
		case tU64:
			v = gofakeit.UintRange(pairToUInt((*r)["u64:min"], (*r)["u64:max"]))
		case tUPtr:
			v = gofakeit.UintRange(pairToUInt((*r)["u64:min"], (*r)["u64:max"]))
		case tF32:
			v = gofakeit.Float32Range(pairToF32((*r)["float32:min"], (*r)["float32:max"]))
		case tF64:
			v = gofakeit.Float64Range(pairToF64((*r)["float64:min"], (*r)["float64:max"]))
		}

		val[kind] = kind.cast(v)
		//switch kind {
		//case tInt, tI8, tI16, tI32, tI64:
		//	val[kind] = castFunc(v, int64(0))
		//case tUint, tU8, tU16, tU32, tU64, tUPtr:
		//	val[kind] = castFunc(v, uint64(0))
		//case tF32, tF64:
		//	val[kind] = castFunc(v, float64(0))
		//}
	}
}

func (r *rule) enumValue(val map[tKind]any) {
	if v, ok := (*r)["enum"]; ok {
		strs := v.([]string)
		if len(strs) > 0 {
			for kind := tInt; kind <= tTime; kind++ {
				val[kind] = kind.cast(randT(strs...))
			}
		}
	}
}
