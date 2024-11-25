package safe_validator

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"strings"
)

// +param: tagv
// +example:
// rkind 	| tagv 					| byKind
// kBool  	| enum:true,false		| string
// kInt   	| enum:-1,-2,33,44,55 	| string
// kUint  	| enum:9,10,1111 		| string
// kString	| enum:a,b,cc 			| string
// kF32  	| enum:3.33,1.22,2.11 	| special judge
// fF64 	| enum:3.33,1.22,2.11	| special judge
func (f fieldOf) validEnum(tagv string) error {
	if util.ElemIn(f.rkind, kF32, kF64) {
		return validEnumFloat()
	}
	if !util.MapOk(
		lo.SliceToMap(strings.Split(tagv, ","), func(item string) (string, est) { return item, _est }),
		f.str) {
		return enumValidErr(tagv, f.rval)
	}
	return nil
}

func validEnumFloat() error {
	return nil
}

func enumFormatErr(enums string) error {
	return fmt.Errorf("enum format error: (%s)", enums)
}

func enumValidErr(enums string, v any) error {
	return fmt.Errorf("enum valid error: %v not in (%s)", v, enums)
}

func enumValidErrFloat(enums string, v any) error {
	return fmt.Errorf("enum valid error: %v not in (%s), float may lose precision for misjudge, recommand float to use range", v, enums)
}
