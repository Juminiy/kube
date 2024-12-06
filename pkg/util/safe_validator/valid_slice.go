package safe_validator

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
)

func (cfg *Config) Slice(v any) bool {
	ok, _ := cfg.sliceOkE(v)
	return ok
}

func (cfg *Config) SliceE(v any) error {
	_, err := cfg.sliceOkE(v)
	return err
}

func (cfg *Config) sliceOkE(v any) (ok bool, err error) {
	tv := indir(v)
	if tv.Val.Kind() != kSlice {
		return true, errValNotSlice
	}

	return cfg.arrayLikeOkE(tv)
}

var errValNotSlice = errors.New("value type is not slice")
var errSliceElemNotStruct = errors.New("slice elem type is not struct")

func (cfg *Config) arrayLikeOkE(tv safe_reflect.TypVal) (ok bool, err error) {
	if tv.Val.Len() == 0 {
		return true, nil
	}

	errHandle := util.NewErrHandle()
	for i := range tv.Val.Len() {
		elemv := wrapv(indirv(tv.Val.Index(i)))
		if elemv.Val.Kind() != kStruct {
			continue
		}
		elemOf := cfg.parseStructOf(elemv)
		ok, err = elemOf.valid(), elemOf.All()

		if errIsTagFormat(err) && cfg.IgnoreTagError {
			continue
		}
		if errHandle.Has(err) && cfg.OnErrorStop {
			return
		}
	}

	return errHandle.Has(), errHandle.All()
}
