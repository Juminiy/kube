package safe_validator

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_cast"
)

// +param: tagv
// +example:
// rkind: kArr, kChan, kMap, kSlice, kString
// byRange limit is int [0, math.MaxInt]
// tagv			| byRange
// len:10		| 10~10
// len:~20		| 0~20
// len:30~ 		| 30~math.MaxInt
// len:20~60 	| 20~60
// len:-1~100	| 0~100
// len:-5~-10	| error
// len:11~2		| error
func (f fieldOf) validLen(tagv string) error {
	lenRangeParsed := parseRange(tagv)
	if !lenRangeParsed.valid {
		return f.lenFormatErr(tagv)
	}

	lenRangeParsed.setLimit(util.NewFloat64(0), util.NewFloat64(safe_cast.ItoF64(util.MaxInt)))
	if rvlen := f.rval.Len(); !util.InRange(
		rvlen, safe_cast.I64toI(*lenRangeParsed.intL), safe_cast.I64toI(*lenRangeParsed.intR)) {
		return f.lenValidErr(rvlen, tagv)
	}
	return nil
}

func (f fieldOf) lenFormatErr(tagv string) error {
	return fmt.Errorf(errTagFormatFmt, f.name, lenOf, tagv)
}

func (f fieldOf) lenValidErr(vlen int, tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("len: %d not in range: (%s)", vlen, tagv))
}