package safe_validator

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/spf13/cast"
	"strings"
)

// +param: tagv
// +example:
// rkind: kInt ~ kF64
// tagv			| byRange
// range:10		| 10~10
// range:~20	| 0~20
// range:30~ 	| 30~math.MaxInt
// range:20~60 	| 20~60
// range:-1~100	| 0~100
// range:-5~-10	| error
// len:11~2		| error
func (f fieldOf) validRange(tagv string) error {
	rangeParsed := parseRange(tagv)
	if !rangeParsed.valid {
		return f.rangeFormatErr(tagv)
	}

	var fl, fr float64
	switch f.rkind {
	case kInt:
		fl, fr = castIPairF64(util.MinInt, util.MaxInt)
	case kI8:
		fl, fr = castIPairF64(util.MinInt8, util.MaxInt8)
	case kI16:
		fl, fr = castIPairF64(util.MinInt16, util.MaxInt16)
	case kI32:
		fl, fr = castIPairF64(util.MinInt32, util.MaxInt32)
	case kI64:
		fl, fr = castIPairF64(util.MinInt64, util.MaxInt64)
	case kUint:
		fl, fr = castUPairF64(util.MinUint, util.MaxUint)
	case kU8:
		fl, fr = castUPairF64(util.MinUint8, util.MaxUint8)
	case kU16:
		fl, fr = castUPairF64(util.MinUint16, util.MaxUint16)
	case kU32:
		fl, fr = castUPairF64(util.MinUint32, util.MaxUint32)
	case kU64, kUPtr:
		fl, fr = castUPairF64(util.MinUint64, util.MaxUint64)
	case kF32:
		fl, fr = float64(util.SmallestNonzeroFloat32), float64(util.MaxFloat32)
	case kF64:
		fl, fr = util.SmallestNonzeroFloat64, util.MaxFloat64
	default:
		panic(errTagKindCheckErr)
	}
	rangeParsed.setLimit(&fl, &fr)

	var ok bool
	switch f.rkind {
	case kInt, kI8, kI16, kI32, kI64:
		ok = util.InRange(cast.ToInt64(f.val), *rangeParsed.intL, *rangeParsed.intR)
	case kUint, kU8, kU16, kU32, kU64, kUPtr:
		ok = util.InRange(cast.ToUint64(f.val), *rangeParsed.uintL, *rangeParsed.uintR)
	case kF32, kF64:
		ok = util.InRange(cast.ToFloat64(f.val), *rangeParsed.floatL, *rangeParsed.floatR)
	}

	if !ok {
		return f.rangeValidErr(tagv)
	}
	return nil
}

func (f fieldOf) rangeFormatErr(tagv string) error {
	return fmt.Errorf(errTagFormatFmt, f.name, rangeOf, tagv)
}

func (f fieldOf) rangeValidErr(tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("is not in range (%s)", tagv))
}

type rangeLR struct {
	intL   *int64
	intR   *int64
	uintL  *uint64
	uintR  *uint64
	floatL *float64
	floatR *float64
	valid  bool
}

func (r *rangeLR) set(rl, rr *float64) {
	r.valid = true
	if rl != nil {
		r.intL = util.New(cast.ToInt64(*rl))
		r.uintL = util.New(cast.ToUint64(*rl))
		r.floatL = util.New(*rl)
	}
	if rr != nil {
		r.intR = util.New(cast.ToInt64(*rr))
		r.uintR = util.New(cast.ToUint64(*rr))
		r.floatR = util.New(*rr)
	}
}

func (r *rangeLR) setLimit(rl, rr *float64) {
	if rl != nil {
		r.intL = util.PtrPairMax(r.intL, util.New(cast.ToInt64(*rl)))
		r.uintL = util.PtrPairMax(r.uintL, util.New(cast.ToUint64(*rl)))
		r.floatL = util.PtrPairMax(r.floatL, rl)
	}
	if rr != nil {
		r.intR = util.PtrPairMin(r.intR, util.New(cast.ToInt64(*rr)))
		r.uintR = util.PtrPairMin(r.uintR, util.New(cast.ToUint64(*rr)))
		r.floatR = util.PtrPairMin(r.floatR, rr)
	}
}

func parseRange(lenRange string) (rangeLr rangeLR) {
	rangelr := strings.Split(lenRange, "~")

	switch len(rangelr) {
	case 1:
		numValid := len(rangelr[0]) > 0
		if !numValid {
			return
		}
		numTry := cast.ToFloat64(rangelr[0])
		rangeLr.set(util.New(numTry), util.New(numTry))

	case 2:
		num0Valid, num1Valid := len(rangelr[0]) > 0, len(rangelr[1]) > 0
		num0Try, num1Try := cast.ToFloat64(rangelr[0]), cast.ToFloat64(rangelr[1])
		if !num0Valid && !num1Valid {
			return
		} else if num0Valid && num1Valid {
			if num0Try > num1Try {
				return
			}
			rangeLr.set(util.New(num0Try), util.New(num1Try))
		} else if num0Valid {
			rangeLr.set(util.New(num0Try), nil)
		} else { // num1Valid
			rangeLr.set(nil, util.New(num1Try))
		}

	default:
		return
	}
	return
}
