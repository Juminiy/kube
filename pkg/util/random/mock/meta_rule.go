package mock

import "github.com/spf13/cast"

type rule map[string]any

func (r *rule) applyInt(minval, maxval string) {

}

func (r *rule) applyUint(minval, maxval string) {

}

func (r *rule) applyMin(minval any) {

}

func (r *rule) applyMax(maxval any) {

}

func (r *rule) applyStringLen(minlen, maxlen string) {
	var lenmin, lenmax = stringDefaultMinLen, stringDefaultMaxLen
	var minnil, maxnil = len(minlen) == 0, len(maxlen) == 0
	if !minnil {
		lenmin = cast.ToInt(minlen)
	}
	if !maxnil {
		lenmax = cast.ToInt(maxlen)
	}
	if lenmin <= 0 || lenmax <= 0 {
		return
	} // invalid tag value

	switch {
	case !minnil && !maxnil && lenmax < lenmin:
		return // invalid tag value

	case !minnil && maxnil:
		if lenmin > lenmax {
			lenmax = stringDefaultRatio * lenmin
		}
	}

	if lenmax > stringMaxLen {
		lenmax = stringMaxLen
	}
	(*r)["string:len:min"] = lenmin
	(*r)["string:len:max"] = lenmax

}

func (r *rule) applyStringCharset(charset ...rune) {
	(*r)["string:char"] = append((*r)["string:char"].([]rune), charset...)
}

func (r *rule) applyTime(lval, rval string) {

}

func (r *rule) applyRangeFns() []func(minval, maxval string) {
	return []func(string, string){
		r.applyInt,
		r.applyUint,
		r.applyStringLen,
		r.applyTime,
	}
}
