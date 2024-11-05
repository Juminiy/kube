// Package mockv2 was generated
package mockv2

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/spf13/cast"
	"maps"
)

const defaultKey = "~default~"

func specialRule(tagkv safe_reflect.TagKV) any {
	for tag := range tagkv {
		switch tag {
		case "uuid", "UUID":
			return stringFunc["uuid"]()

		default:
			return nil
		}
	}
	return nil
}

var _defaultRule = util.MapsMerge(
	boolRule,
	intRule,
	uintRule,
	floatRule,
	timeRule,
	stringRule,
)

type Rule struct {
	t safe_reflect.TagKV
	r rule
}

func newRule(tagkv safe_reflect.TagKV) *Rule {
	return &Rule{
		t: tagkv,
		r: maps.Clone(_defaultRule),
	}
}

func (r *Rule) parse() *Rule {
	for rk, rv := range r.t {
		switch rk {
		case "range", // int, uint, float, time.Time
			"len",                    // string
			"now", "before", "after": // time.Time
			r.parseRange(rv)

		case "min", "max": // int, uint, float
			r.parseNum(rk, rv)

		case "positive", // int, uint, float
			"negative",      // int, float
			"regexp",        // string
			"uuid",          // string
			"alpha",         // string
			"numeric",       // string
			"symbol",        // string
			"binary", "bin", // string
			"octal", "oct", // string
			"hexadecimal", "hex", //string
			"timestamp": //string
			r.parseString(rk)

		case "char", // string
			"enum": // any
			r.parseEnum(rk, rv)
		}
	}

	return r
}

// parseRange
// +support-tag: range, len
// +example:
// vala~valb;
// vala~;
// ~valb;
func (r *Rule) parseRange(rv string) {
	if len(rv) == 0 {
		return
	}
	minMax := split(rv, util.Tilde)
	onlyMax := rv[0] == '~'

	minval, maxval := "", ""
	switch len(minMax) {
	case 2:
		minval, maxval = minMax[0], minMax[1]

	case 1:
		if onlyMax {
			maxval = minMax[0]
		} else {
			minval = minMax[0]
		}
	}

	for _, fn := range r.r.applyRangeFns() {
		fn(minval, maxval)
	}

}

// +support-tag: min, max
// +example: k1:114514;k2:9527
func (r *Rule) parseNum(tag, rv string) {
	i64 := cast.ToInt64(rv)
	switch tag {
	case "min":
		r.r.applyMin(i64)

	case "max":
		r.r.applyMax(i64)
	}
}

// +support-tag: enum, char
// +example: enum:1,2,3;char:c,~,);
func (r *Rule) parseEnum(tag, rv string) {
	vals := split(rv, ",")
	switch tag {
	case "enum":
		r.r["enum"] = vals

	case "char":
		r.r.applyStringCharset(util.StringSlice2RuneSlice(vals)...)

	}
}

func (r *Rule) parseString(tag string) {
	r.r.applyStringTag(tag)
}

var _defaultValue = map[tKind]any{
	tBool:   defaultBool(),
	tInt:    cast.ToInt(defaultInt()),
	tI8:     cast.ToInt8(defaultI8()),
	tI16:    cast.ToInt16(defaultInt()),
	tI32:    cast.ToInt32(defaultInt()),
	tI64:    defaultInt(),
	tUint:   cast.ToUint(defaultUint()),
	tU8:     cast.ToUint8(defaultU8()),
	tU16:    cast.ToUint16(defaultUint()),
	tU32:    cast.ToUint32(defaultUint()),
	tU64:    defaultUint(),
	tUPtr:   uintptr(defaultUint()),
	tF32:    cast.ToFloat32(defaultFloat()),
	tF64:    defaultFloat(),
	tString: defaultString(),
}

func (r *Rule) value() map[tKind]any {
	if Default {
		return _defaultValue
	}
	val := maps.Clone(_defaultValue)
	r.r.setValue(val)
	return val
}
