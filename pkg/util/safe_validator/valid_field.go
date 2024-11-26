package safe_validator

import (
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
}

func (f fieldOf) valid() error {
	for _, tagk := range _tagPrior {
		// _timeTyp for special judge

		if !util.MapOk(f.tag, tagk) || // app tag no tagk desc
			!tagApplyKind(tagk, f.rkind) { // tagk not apply field kind
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
