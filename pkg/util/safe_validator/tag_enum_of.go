package safe_validator

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_cast/safe_parse"
	"github.com/samber/lo"
	"strings"
)

/*
 * in enums
 */
// validEnum
// +param: tagv
// +example:
// rkind 	| tagv 					| byKind/return
// kBool  	| enum:true,false		| string
// kInt   	| enum:-1,-2,33,44,55 	| string
// kUint  	| enum:9,10,1111 		| string
// kString	| enum:a,b,cc 			| string
// kF32  	| enum:3.33,1.22,2.11 	| special judge
// kF64 	| enum:3.33,1.22,2.11	| special judge
// kPtr 	| enum:x,xx,xxx			| indir
// kU8		| enum:1022,-9999,257	| error or silent(Config.IgnoreTagError)
// kInt 	| enum:rrr,vvv,666		| apply 666
func (f fieldOf) validEnum(tagv string) error {
	enums, ok := f.parseEnums(tagv)
	if !ok {
		return f.parseEnumsErr(tagv)
	}

	if !f.valueInEnums(enums) {
		if util.ElemIn(f.rkind, kF32, kF64) {
			return f.enumInValidFloatErr(tagv)
		} else {
			return f.enumInValidErr(tagv)
		}
	}
	return nil
}

func (f fieldOf) enumInValidErr(tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("is not in enums: (%s)", tagv))
}

func (f fieldOf) enumInValidFloatErr(tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("is not in enums: (%s), float may lose precision for misjudge, recommand float to use range", tagv))
}

/*
 * not in enums
 */
// validEnumNot
// refer to: validEnum
func (f fieldOf) validEnumNot(tagv string) error {
	enums, ok := f.parseEnums(tagv)
	if !ok {
		return f.parseEnumsErr(tagv)
	}

	if f.valueInEnums(enums) {
		if util.ElemIn(f.rkind, kF32, kF64) {
			return f.enumNotInValidFloatErr(tagv)
		} else {
			return f.enumNotInValidErr(tagv)
		}
	}
	return nil
}

func (f fieldOf) enumNotInValidErr(tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("is in enums: (%s)", tagv))
}

func (f fieldOf) enumNotInValidFloatErr(tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("is in enums: (%s), float may lose precision for misjudge, recommand float to use range", tagv))
}

func (f fieldOf) parseEnums(tagv string) ([]string, bool) {
	enums := strings.Split(tagv, ",")
	ok := lo.ContainsBy(enums, func(item string) bool {
		_, ok := safe_parse.Parse(item).Get(f.rkind)
		if !ok {
			_, ok = safe_parse.Parse(item).GetByRT(f.rval.Type())
		}
		return ok
	})
	if !ok {
		return nil, false
	}
	return enums, true
}

func (f fieldOf) parseEnumsErr(tagv string) error {
	return fmt.Errorf(errTagFormatFmt, f.name, enumOf, tagv)
}

func (f fieldOf) valueInEnums(enums []string) bool {
	if util.ElemIn(f.rkind, kF32, kF64) {
		return util.MapOk(
			util.Slice2MapWhen(
				enums, // enum values
				func(item string, index int) bool { // predict func
					var ok bool
					if f.rkind == kF32 {
						_, ok = safe_parse.ParseFloat32(item)
					} else {
						_, ok = safe_parse.ParseFloat64(item)
					}
					return ok
				},
				func(item string) (float64, est) { // transform func
					f64v, _ := safe_parse.ParseFloat64(item)
					return f64v, _est
				},
			),
			f.rval.Float())
	}
	return lo.Contains(enums, f.str)
}
