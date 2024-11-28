package safe_validator

import (
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"reflect"
	"time"
)

type kind = reflect.Kind

// kind mirror of reflect.Kind
const (
	kInvalid kind = iota
	kBool
	kInt
	kI8
	kI16
	kI32
	kI64
	kUint
	kU8
	kU16
	kU32
	kU64
	kUPtr
	kF32
	kF64
	kC64
	kC128
	kArr
	kChan
	kFunc
	kAny
	kMap
	kPtr
	kSlice
	kString
	kStruct
	kUnsafePtr
)

// kind extended
const (
	kTime = iota + kUnsafePtr + 1
	kOmitBool
	kOmitComplex
	kNumber
	kLikeStr
	kDirectCompare
	kLikePtr
	kAll
)

// _extTyp
// readOnly
var _extTyp = map[kind][]kind{
	kNumber:        {2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},             // kInt ~ kF64
	kLikeStr:       {kString},                                                // kString, ByteSlice([]byte), fmt.Stringer
	kDirectCompare: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, kString}, // kNumber, kString
	kLikePtr:       {kChan, kFunc, kMap, kPtr, kUnsafePtr, kAny, kSlice},
	kAll:           {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26},
}

// type empty struct
type est = struct{}

// value empty struct
var _est = est{}
var _timeTyp = safe_reflect.Of(time.Time{}).Typ
