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
// kF64 	| enum:3.33,1.22,2.11	| special judge
// kPtr 	| enum:x,xx,xxx			| indir
func (f fieldOf) validEnum(tagv string) error {
	if util.ElemIn(f.rkind, kF32, kF64) {
		return f.validEnumFloat(tagv)
	}

	if !util.MapOk(
		lo.SliceToMap(strings.Split(tagv, ","), func(item string) (string, est) { return item, _est }),
		f.str) {
		return f.enumValidErr(tagv)
	}
	return nil
}

func (f fieldOf) validEnumFloat(tagv string) error {
	return nil
}

func (f fieldOf) enumFormatErr(tagv string) error {
	return fmt.Errorf(errTagFormatFmt, f.name, enumOf, tagv)
}

func (f fieldOf) enumValidErr(tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("is not in enums: (%s)", tagv))
}

func (f fieldOf) enumValidErrFloat(tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("is not in enums: (%s), float may lose precision for misjudge, recommand float to use range", tagv))
}
