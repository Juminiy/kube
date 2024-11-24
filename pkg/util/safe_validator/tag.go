package safe_validator

import "reflect"

type option struct {
	notNil  bool
	notZero bool
}

// apply tag keys
const (
	// notNil:
	//	case Chan, Func, Interface, Map, Pointer, Slice, UnsafePointer: return !v.IsNil()
	// 	default: skip tag
	notNil = "not_nil"

	// notZero:
	//	case Chan, Func, Interface, Map, Pointer, Slice, UnsafePointer: return !v.IsNil()
	// 	default: !IsZero()
	notZero   = "not_zero"
	lenOf     = "len"
	rangeOf   = "range"
	enumOf    = "enum"
	ruleOf    = "rule"
	regexOf   = "regex"
	defaultOf = "default"
)

type kind = reflect.Kind

const (
	kTime kind = 27
)

type est = struct{}

var _est = est{}

var apply = map[string]map[kind]est{
	notNil: map[kind]est{reflect.Chan: _est},
}
