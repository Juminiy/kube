package safe_validator

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_cast"
	"math"
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
	lenRangeParsed.setLimitInt(int64(0), int64(math.MaxInt))

	if !f.valueLenInRange(lenRangeParsed) {
		return f.lenInValidErr(f.rval.Len(), tagv)
	}
	return nil
}

func (f fieldOf) lenInValidErr(vlen int, tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("len: (%d) not in range: (%s)", vlen, tagv))
}

func (f fieldOf) validLenNot(tagv string) error {
	lenRangeParsed := parseRange(tagv)
	if !lenRangeParsed.valid {
		return f.lenFormatErr(tagv)
	}
	lenRangeParsed.setLimitInt(int64(0), int64(math.MaxInt))

	if f.valueLenInRange(lenRangeParsed) {
		return f.lenNotInValidErr(f.rval.Len(), tagv)
	}
	return nil
}

func (f fieldOf) lenNotInValidErr(vlen int, tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("len: (%d) in range: (%s)", vlen, tagv))
}

func (f fieldOf) lenFormatErr(tagv string) error {
	return fmt.Errorf(errTagFormatFmt, f.name, lenOf, tagv)
}

func (f fieldOf) valueLenInRange(lenRangeParsed rangeLR) bool {
	return util.InRange(safe_cast.ItoI64(f.rval.Len()),
		*lenRangeParsed.intL, *lenRangeParsed.intR)
}
