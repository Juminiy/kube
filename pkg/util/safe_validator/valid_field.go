package safe_validator

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/spf13/cast"
	"reflect"
)

type fieldOf struct {
	name  string
	rkind kind
	rval  reflect.Value
	val   any
	str   string
	tag   safe_reflect.TagKV

	cfg *Config
}

func (f fieldOf) valid() error {
	for _, tagk := range _prior {
		// _timeTyp for special judge

		if !util.MapOk(f.tag, tagk) || // app tag no tagk desc
			!tagApplyKind(f.cfg.apply, tagk, f.rkind) { // tagk not apply field kind
			continue
		}

		tagv := f.tag[tagk]
		cloneF, skip, err := f.tagApplyIndirect(tagk, tagv)
		if err != nil {
			return err
		}
		if skip {
			continue
		}
		switch tagk {
		case enumOf:
			err = cloneF.validEnum(tagv)
		case notNil:
			err = cloneF.validNotNil()
		case notZero:
			err = cloneF.validNotZero()
		case rangeOf:
			err = cloneF.validRange(tagv)
		case lenOf:
			err = cloneF.validLen(tagv)
		case ruleOf:
			err = cloneF.validRule(tagv)
		case regexOf:
			err = cloneF.validRegex(tagv)
		case defaultOf:
			cloneF.setDefault(tagv)
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (f fieldOf) tagApplyIndirect(tagk, tagv string) (cloneF fieldOf, skip bool, err error) {
	if f.rkind != kPtr || util.ElemIn(tagk,
		notNil, defaultOf) {
		return f, false, nil
	}
	if !util.ElemIn(tagk,
		enumOf, lenOf, notZero, rangeOf, regexOf, ruleOf) {
		return f, true, nil
	}
	if err = f.errPointerNil(tagk, tagv); err != nil {
		return f, false, err
	}
	cloneF, skip = f.indirect(tagk)
	return cloneF, skip, nil // skip indirect-value type mismatch tag
}

func (f fieldOf) indirect(tag string) (cloneF fieldOf, skip bool) {
	cloneF = f
	if cloneF.cfg.IndirectValue &&
		cloneF.rkind == kPtr {
		cloneF.rval = indirv(cloneF.rval)
		cloneF.rkind = cloneF.rval.Kind()
		cloneF.val = cloneF.rval.Interface()
		cloneF.str = cast.ToString(cloneF.val)
		ok := tagApplyKind(_apply, tag, cloneF.rkind)
		if ok {
			return cloneF, false
		} else {
			return f, true
		}
	}
	return f, true
}

func (f fieldOf) errPointerNil(tagk, tagv string) error {
	if f.rkind == kPtr && f.rval.IsNil() {
		return fmt.Errorf(errPtrNilFmt, f.name, tagk, tagv)
	}
	return nil
}
