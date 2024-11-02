package mock

import (
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"reflect"
	"strings"
)

// struct app tag
const (
	mockTag = "mock"
)

type tKind reflect.Kind

// alias of reflect.Kind
const (
	tInvalid tKind = iota
	tBool
	tInt
	tI8
	tI16
	tI32
	tI64
	tUint
	tU8
	tU16
	tU32
	tU64
	tUPtr
	tF32
	tF64
	tC64
	tC128
	tArr
	tChan
	tFunc
	tAny
	tMap
	tPtr
	tSlice
	tString
	tStruct
	tUnsafePtr
)

// define
const (
	tTime tKind = iota + 27
)

// short for safe_reflect.IndirectOf
var indir = safe_reflect.IndirectOf

var split = strings.Split
