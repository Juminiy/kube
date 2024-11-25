package safe_validator

import (
	"github.com/Juminiy/kube/pkg/util"
)

// apply tag keys
const (
	// notNil:
	//	case Chan, Func, Interface, Map, Pointer, Slice, UnsafePointer: return !v.IsNil()
	// 	default: skip tag
	// notNil cannot notZero
	notNil = "not_nil"

	// notZero:
	//	case Chan, Func, Interface, Map, Pointer, Slice, UnsafePointer: return !v.IsNil()
	// 	default: !IsZero()
	// notZero can notNil
	notZero   = "not_zero"
	lenOf     = "len"
	rangeOf   = "range"
	enumOf    = "enum"
	ruleOf    = "rule"
	regexOf   = "regex"
	defaultOf = "default"
)

// const
var _tagPrior = []string{enumOf, notNil, notZero, rangeOf, lenOf, ruleOf, regexOf, defaultOf}

var apply = map[string]map[kind]est{
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
	for tag, kinds := range apply {
		switch {
		case util.MapOk(kinds, kNumber):
			util.MapInsert(kinds, util.MapElem(_extTyp, kNumber)...)
			fallthrough
		case util.MapOk(kinds, kLikeStr):
			util.MapInsert(kinds, util.MapElem(_extTyp, kLikeStr)...)
			fallthrough
		case util.MapOk(kinds, kDirectCompare):
			util.MapInsert(kinds, util.MapElem(_extTyp, kDirectCompare)...)
			fallthrough
		case util.MapOk(kinds, kLikePtr):
			util.MapInsert(kinds, util.MapElem(_extTyp, kLikePtr)...)
			fallthrough
		case util.MapOk(kinds, kAll):
			util.MapInsert(kinds, util.MapElem(_extTyp, kAll)...)
		case util.MapOk(kinds, kBool):
			util.MapDelete(kinds, kBool)
		}
		util.MapDelete(kinds, kC64, kC128, kInvalid)
		apply[tag] = kinds
	}
}

func tagApplyKind(tag string, k kind) bool {
	return util.MapOk(util.MapElem(apply, tag), k)
}
