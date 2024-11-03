package mock

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/zerobuf"
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

const lowerStr = "abcdefghijklmnopqrstuvwxyz"
const upperStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaStr = lowerStr + upperStr
const numericStr = "0123456789"
const alphaNumericStr = alphaStr + numericStr
const binStr = "01"
const octStr = "01234567"
const hexStr = "01234567890abcdefABCDEF"
const specialStr = "@#$%&?|!(){}<>=*+-_:;,."
const specialSafeStr = "!@.-_*" // https://github.com/1Password/spg/pull/22
const spaceStr = " "
const allStr = lowerStr + upperStr + numericStr + specialStr + spaceStr
const vowels = "aeiou"
const hashtag = '#'
const questionmark = '?'
const dash = '-'
const base58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const minUint = 0
const maxUint = ^uint(0)
const minInt = -maxInt - 1
const maxInt = int(^uint(0) >> 1)
const is32bit = ^uint(0)>>32 == 0

func stringByRunes(r []rune, size int) string {
	buf := zerobuf.Get()

	for range size {
		buf.WriteString(string(randT(r...)))
	}

	bufStr := buf.UnsafeString()
	buf.Free()
	return bufStr
}
