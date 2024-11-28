package safe_validator

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
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
		var err error
		switch tagk {
		case enumOf:
			err = f.validEnum(tagv)
		case notNil:
			err = f.validNotNil()
		case notZero:
			err = f.validNotZero()
		case rangeOf:
			err = f.validRange(tagv)
		case lenOf:
			err = f.validLen(tagv)
		case ruleOf:
			err = f.validRule(tagv)
		case regexOf:
			err = f.validRegex(tagv)
		case defaultOf:
			f.setDefault(tagv)
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (f fieldOf) indirect(tag string) (fieldOf, bool) {
	cloneF := f
	if cloneF.cfg.IndirectValue {
		cloneF.rval = indirv(cloneF.rval)
		cloneF.rkind = cloneF.rval.Kind()
		cloneF.val = cloneF.rval.Interface()
		ok := tagApplyKind(_apply, tag, cloneF.rkind)
		if !ok {
			return cloneF, false
		}
	}
	return cloneF, true
}

func (f fieldOf) errPointerNil(tagk, tagv string) error {
	if f.rval.Kind() == kPtr && f.rval.IsNil() {
		return fmt.Errorf(errPtrNilFmt, f.name, tagk, tagv)
	}
	return nil
}
