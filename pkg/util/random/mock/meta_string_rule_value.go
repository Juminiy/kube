package mock

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

type StringFunc func() string

/*
 * stringFunc
 */
var stringFunc = map[string]StringFunc{
	defaultKey: defaultString,
	"uuid":     uuid.NewString,
}

var defaultString = func() string {
	return gofakeit.LetterN(stringDefaultMaxLen)
}

const (
	stringDefaultMinLen = 1
	stringDefaultMaxLen = 16
	stringDefaultRatio  = 2
	stringMaxLen        = util.Ki
)

// rule to make stringFunc
var stringRule = rule{
	"string:char":    []rune{},
	"string:len:min": stringDefaultMinLen,
	"string:len:max": stringDefaultMaxLen,
}

func (r *rule) applyStringLen(minlen, maxlen string) {
	lenmin, lenmax := rangeOfInt64(minlen, maxlen, stringDefaultMinLen, stringDefaultMaxLen, stringMaxLen)
	(*r)["string:len:min"], (*r)["string:len:max"] = int(lenmin), int(lenmax)

}

func (r *rule) applyStringCharset(charset ...rune) {
	(*r)["string:char"] = append((*r)["string:char"].([]rune), charset...)
}

func (r *rule) applyStringTag(tagkey string) {
	switch tagkey {
	case "regexp":

	case "uuid":

	case "timestamp":

	case "alpha":
		r.applyStringCharset(util.String2RuneSlice(alphaStr)...)

	case "numeric":
		r.applyStringCharset(util.String2RuneSlice(numericStr)...)

	case "symbol":
		r.applyStringCharset(util.String2RuneSlice(specialSafeStr)...)

	case "binary", "bin":
		r.applyStringCharset(util.String2RuneSlice(binStr)...)

	case "octal", "oct":
		r.applyStringCharset(util.String2RuneSlice(octStr)...)

	case "hexadecimal", "hex":
		r.applyStringCharset(util.String2RuneSlice(hexStr)...)

	}

}

func (r *rule) stringValue(val map[tKind]any) {
	r.stringFromRunes(val)
}

func (r *rule) stringFromRunes(val map[tKind]any) {
	size := gofakeit.IntRange(pairToInt((*r)["string:len:min"], (*r)["string:len:max"]))
	runes := (*r)["string:char"].([]rune)
	if len(runes) == 0 {
		r.applyStringCharset(util.String2RuneSlice(alphaNumericStr)...)
	}
	val[tString] = randStringByRunes((*r)["string:char"].([]rune), size)
}
