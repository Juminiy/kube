package safe_validator

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"regexp"
	"strings"
)

// apply tag keys
// return true, error is nil
// return false, error not nil
// never panic
const (
	invalidTagK = "-"

	// notNil:
	// case kChan, kFunc, kAny, kMap, kPtr, kSlice, kUnsafePtr: return !v.IsNil()
	// default: skip tag
	// compatible with reflect.Value.IsNil() notNil cannot notZero
	notNil = "not_nil"
	isNil  = "nil"

	// notZero:
	// case kChan, kFunc, kAny, kMap, kPtr, kSlice, kUnsafePtr: return !v.IsNil()
	// default: !v.IsZero()
	// compatible with reflect.Value.IsZero() notZero can notNil
	// when Config.IndirectValue=true case kPtr: indir
	notZero = "not_zero"
	isZero  = "zero"

	// lenOf
	// case kArr, kChan, kMap, kSlice, kString: return util.InRange(rangeL, v.Len(), rangeR)
	// default: skip tag
	// when Config.IndirectValue=true case kPtr: indir
	lenOf  = "len"
	notLen = "not_len"

	// rangeOf
	// case kInt ~ kF64: return util.InRange(rangeL, v.Len(), rangeR)
	// default: skip tag
	// when Config.IndirectValue=true case kPtr: indir
	rangeOf  = "range"
	notRange = "not_range"

	// enumOf
	// case kBool ~ kF64, kString: return util.ElemIn(v, parsedEnums...)
	// default: skip tag
	// when Config.IndirectValue=true case kPtr: indir
	enumOf  = "enum"
	notEnum = "not_enum"

	// ruleOf
	// case kString: valid rule
	// default: skip tag
	// when Config.IndirectValue=true case kPtr: indir
	ruleOf  = "rule"
	notRule = "not_rule"

	// regexOf
	// case kString: valid regex
	// default: skip tag
	// when Config.IndirectValue=true case kPtr: indir
	regexOf  = "regex"
	notRegex = "not_regex"

	// defaultOf
	// case v.CanSet() && v.IsZero(): v.Set()
	defaultOf = "default"
)

var (
	_parseTagKPrefixRegexp = regexp.MustCompile(`(not_|!)*`)
)

// len, not_not_len, !!len, not_!len, !not_len -> len
// !len, not_len, !!!len, !not_!len -> not_len
// tagk suffix not in (notNil, enumOf, notZero, rangeOf, lenOf, ruleOf, regexOf, defaultOf) -> -
func parseTagK(tagk string) string {
	tagk = strings.TrimSpace(tagk)
	var matchSuffix = invalidTagK
	for _, tagK := range _prior {
		if strings.HasSuffix(tagk, tagK) {
			matchSuffix = tagSuffix(tagK)
			break
		}
	}
	if util.ElemIn(matchSuffix,
		invalidTagK, defaultOf) {
		return matchSuffix
	}
	matchPrefix := strings.TrimSuffix(tagk, matchSuffix)
	if !_parseTagKPrefixRegexp.MatchString(matchPrefix) {
		return invalidTagK
	}
	if util.IsOdd(strings.Count(matchPrefix, "not_") + strings.Count(matchPrefix, "!")) {
		return "not_" + matchSuffix
	} else {
		return matchSuffix
	}
}

func tagSuffix(tagk string) string {
	return util.StringTrimPrefix(tagk, "not_", "!")
}

// readOnly
var _prior = []string{
	notNil, isNil,
	enumOf, notEnum,
	notZero, isZero,
	rangeOf, notRange,
	lenOf, notLen,
	ruleOf, notRule,
	regexOf, notRegex,
	defaultOf}

var _readVTagK = _prior[2:14]

type tagApplyKindT map[string]map[kind]est

// readOnly
var _apply = tagApplyKindT{
	notNil:    {kLikePtr: _est},
	notZero:   {kAll: _est, kOmitBool: _est},
	lenOf:     {kArr: _est, kChan: _est, kMap: _est, kSlice: _est, kString: _est},
	rangeOf:   {kNumber: _est},
	enumOf:    {kDirectCompare: _est},
	ruleOf:    {kLikeStr: _est},
	regexOf:   {kLikeStr: _est},
	defaultOf: {kDirectCompare: _est},
}

func init() {
	for _, tagK := range _prior {
		for _, tagKReserve := range _prior {
			if tagK != tagKReserve && tagSuffix(tagK) == tagSuffix(tagKReserve) {
				if util.MapOk(_apply, tagK) && !util.MapOk(_apply, tagKReserve) {
					_apply[tagKReserve] = _apply[tagK]
					break
				} else if !util.MapOk(_apply, tagK) && util.MapOk(_apply, tagKReserve) {
					_apply[tagK] = _apply[tagKReserve]
					break
				}
			}
		}
	}
	for tag, kinds := range _apply {
		if util.MapOk(kinds, kNumber) {
			util.MapInsert(kinds, util.MapElem(_extTyp, kNumber)...)
		}
		if util.MapOk(kinds, kLikeStr) {
			util.MapInsert(kinds, util.MapElem(_extTyp, kLikeStr)...)
		}
		if util.MapOk(kinds, kDirectCompare) {
			util.MapInsert(kinds, util.MapElem(_extTyp, kDirectCompare)...)
		}
		if util.MapOk(kinds, kLikePtr) {
			util.MapInsert(kinds, util.MapElem(_extTyp, kLikePtr)...)
		}
		if util.MapOk(kinds, kAll) {
			util.MapInsert(kinds, util.MapElem(_extTyp, kAll)...)
		}
		if util.MapOk(kinds, kBool) {
			util.MapDelete(kinds, kBool)
		}
		util.MapDelete(kinds, kC64, kC128, kInvalid)
		_apply[tag] = kinds
	}
}

func tagApplyKind(apply tagApplyKindT, tag string, k kind) bool {
	return util.MapOk(util.MapElem(apply, tag), k)
}

var errTagKindCheckErr = errors.New("tag apply kind check error")

const errTagConflictFmt = "format invalid, field: (%s), tagKey: (%s) conflict with tagKey: (%s)"
const errTagFormatFmt = "format invalid, field: (%s), tagKey: (%s), tagVal: (%s)"
const errValInvalidFmt = "value invalid, field: (%s), value: (%v) %s"
const errPtrNilFmt = "value invalid, field: (%s), tagKey: (%s), tagVal: (%s), by pointer but pointer is nil"

func errIsTagFormat(err error) bool {
	return err != nil && strings.Contains(err.Error(), "format invalid")
}
