package safe_validator

import (
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
)

var indir = safe_reflect.IndirectOf
var indirv = safe_reflect.IndirectRVal

//func castIPairF64[I ~int | int8 | int16 | int32 | int64](v0, v1 I) (float64, float64) {
//	return safe_cast.ItoF64(v0), safe_cast.ItoF64(v1)
//}
//
//func castUPairF64[U ~uint | uint8 | uint16 | uint32 | uint64 | uintptr](v0, v1 U) (float64, float64) {
//	return safe_cast.UtoF64(v0), safe_cast.UtoF64(v1)
//}
