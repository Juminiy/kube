package safe_validator

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
)

func (cfg *Config) Map(v any) bool {
	ok, _ := cfg.mapOkE(v)
	return ok
}

func (cfg *Config) MapE(v any) error {
	_, err := cfg.mapOkE(v)
	return err
}

func (cfg *Config) mapOkE(v any) (ok bool, err error) {
	tv := indir(v)
	if tv.Val.Kind() != kMap {
		return true, errValNotMap
	}

	if tv.Val.Len() == 0 {
		return true, nil
	}

	errHandle := util.NewErrHandle()
	for mapIter := tv.Val.MapRange(); mapIter.Next(); {
		elemv := wrapv(indirv(mapIter.Value()))
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

var errValNotMap = errors.New("value type is not array")
var errMapElemNotStruct = errors.New("map elem type is not struct")
