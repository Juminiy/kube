package safe_validator

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
)

// apply tag keys
const (
	// notNil:
	// case kChan, kFunc, kAny, kMap, kPtr, kSlice, kUnsafePtr: return !v.IsNil()
	// default: skip tag
	// notNil cannot notZero
	notNil = "not_nil"

	// notZero:
	// case kChan, kFunc, kAny, kMap, kPtr, kSlice, kUnsafePtr: return !v.IsNil()
	// default: !IsZero()
	// notZero can notNil
	notZero = "not_zero"

	// lenOf
	// case kArr, kChan, kMap, kSlice, kString: return util.InRange(rangeL, v.Len(), rangeR)
	// default: skip tag
	lenOf = "len"

	// rangeOf
	rangeOf = "range"

	// enumOf
	enumOf = "enum"

	// ruleOf
	// allow: kString
	ruleOf = "rule"

	// regexOf
	// allow: kString
	regexOf = "regex"

	// defaultOf
	// case v.CanSet() && v.IsZero()
	defaultOf = "default"
)

// readOnly
var _prior = []string{notNil, enumOf, notZero, rangeOf, lenOf, ruleOf, regexOf, defaultOf}

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
	defaultOf: {kAll: _est},
}

func init() {
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

const errTagFormatFmt = "format invalid, field: (%s), tagKey: (%s), tagVal: (%s)"
const errValInvalidFmt = "value invalid, field: (%s), value: (%v) %s"
const errPtrNilFmt = "pointer to value invalid, field: (%s), tagKey: (%s), tagVal: (%s), pointer is nil"
