package safe_validator

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_cast/safe_parse"
	"strings"
)

// +param: tagv
// +example:
// rkind: kInt ~ kF64
// tagv			| byRange
// range:10		| 10~10
// range:~20	| kInt: math.MinInt~20, kUint: 0~20, kF64: -math.MaxFloat64~20, etc.
// range:30~ 	| kInt: 30~math.MaxInt, kUint: 30~math.MaxUint, kF64: 30~math.MaxFloat64, etc.
// range:20~60 	| 20~60
// range:-1~100	| 0~100
// range:-5~-10	| error
// len:11~2		| error
func (f fieldOf) validRange(tagv string) error {
	rangeParsed := parseRange(tagv)
	if !rangeParsed.valid {
		return f.rangeFormatErr(tagv)
	}
	rangeParsed.setLimitKind(f.rkind)

	if !f.valueInRange(rangeParsed) {
		return f.rangeValidErr(tagv)
	}
	return nil
}

func (f fieldOf) rangeValidErr(tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("is not in range (%s)", tagv))
}

func (f fieldOf) validRangeNot(tagv string) error {
	rangeParsed := parseRange(tagv)
	if !rangeParsed.valid {
		return f.rangeFormatErr(tagv)
	}
	rangeParsed.setLimitKind(f.rkind)

	if f.valueInRange(rangeParsed) {
		return f.rangeNotValidErr(tagv)
	}
	return nil
}

func (f fieldOf) rangeNotValidErr(tagv string) error {
	return fmt.Errorf(errValInvalidFmt, f.name, f.val,
		fmt.Sprintf("is not in range (%s)", tagv))
}

func parseRange(lenRange string) (rangeLr rangeLR) {
	rangelr := strings.Split(lenRange, "~")

	switch len(rangelr) {
	case 1:
		rangeLr.set(rangelr[0], rangelr[0])

	case 2:
		rangeLr.set(rangelr[0], rangelr[1])

	default:
		return
	}
	return
}

func (f fieldOf) rangeFormatErr(tagv string) error {
	return fmt.Errorf(errTagFormatFmt, f.name, rangeOf, tagv)
}

func (f fieldOf) valueInRange(rangeParsed rangeLR) (ok bool) {
	switch f.rkind {
	case kInt, kI8, kI16, kI32, kI64:
		ok = util.InRange(f.rval.Int(), *rangeParsed.intL, *rangeParsed.intR)
	case kUint, kU8, kU16, kU32, kU64, kUPtr:
		ok = util.InRange(f.rval.Uint(), *rangeParsed.uintL, *rangeParsed.uintR)
	case kF32, kF64:
		ok = util.InRange(f.rval.Float(), *rangeParsed.floatL, *rangeParsed.floatR)
	default:
		panic(errTagKindCheckErr)
	}
	return
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

func (r *rangeLR) set(rl, rr string) {
	parserl := safe_parse.ParseNumber(rl)
	r.intL = parserl.I64()
	r.uintL = parserl.U64()
	r.floatL = parserl.F64()

	parserr := safe_parse.ParseNumber(rr)
	r.intR = parserr.I64()
	r.uintR = parserr.U64()
	r.floatR = parserr.F64()

	r.valid = r.floatL != nil || r.floatR != nil
}

func (r *rangeLR) setLimit(rl, rr string) {
	parserl := safe_parse.ParseNumber(rl)
	r.intL = util.PtrPairMax(r.intL, parserl.I64())
	r.uintL = util.PtrPairMax(r.uintL, parserl.U64())
	r.floatL = util.PtrPairMax(r.floatL, parserl.F64())

	parserr := safe_parse.ParseNumber(rr)
	r.intR = util.PtrPairMin(r.intR, parserr.I64())
	r.uintR = util.PtrPairMin(r.uintR, parserr.U64())
	r.floatR = util.PtrPairMin(r.floatR, parserr.F64())
}

func (r *rangeLR) setLimitInt(rl, rr int64) {
	r.intL = util.PtrPairMax(r.intL, &rl)
	r.intR = util.PtrPairMin(r.intR, &rr)
}

func (r *rangeLR) setLimitUint(rl, rr uint64) {
	r.uintL = util.PtrPairMax(r.uintL, &rl)
	r.uintR = util.PtrPairMin(r.uintR, &rr)
}

func (r *rangeLR) setLimitFloat(rl, rr float64) {
	r.floatL = util.PtrPairMax(r.floatL, &rl)
	r.floatR = util.PtrPairMin(r.floatR, &rr)
}

func (r *rangeLR) setLimitKind(kd kind) {
	switch kd {
	case kInt:
		r.setLimitInt(int64(util.MinInt), int64(util.MaxInt))
	case kI8:
		r.setLimitInt(int64(util.MinInt8), int64(util.MaxInt8))
	case kI16:
		r.setLimitInt(int64(util.MinInt16), int64(util.MaxInt16))
	case kI32:
		r.setLimitInt(int64(util.MinInt32), int64(util.MaxInt32))
	case kI64:
		r.setLimitInt(util.MinInt64, util.MaxInt64)
	case kUint:
		r.setLimitUint(uint64(util.MinUint), uint64(util.MaxUint))
	case kU8:
		r.setLimitUint(uint64(util.MinUint8), uint64(util.MaxUint8))
	case kU16:
		r.setLimitUint(uint64(util.MinUint16), uint64(util.MaxUint16))
	case kU32:
		r.setLimitUint(uint64(util.MinUint32), uint64(util.MaxUint32))
	case kU64, kUPtr:
		r.setLimitUint(util.MinUint64, util.MaxUint64)
	case kF32:
		r.setLimitFloat(float64(util.MinFloat32), float64(util.MaxFloat32))
	case kF64:
		r.setLimitFloat(util.MinFloat64, util.MaxFloat64)
	default:
		panic(errTagKindCheckErr)
	}
}
