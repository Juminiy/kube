package safe_validator

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/spf13/cast"
	"reflect"
)

func Struct(v any) bool {
	return parseStruct(v).valid()
}

type structOf struct {
	FieldRTyp  map[string]reflect.Type
	FieldRVal  map[string]reflect.Value
	FieldVal   map[string]any
	FieldTagKv safe_reflect.FieldTagKV
	CanSet     bool
	util.ErrHandle
}

func parseStruct(v any) *structOf {
	tv := indir(v)
	if tv.Typ.Kind() != kStruct {
		return nil
	}
	return &structOf{
		FieldRTyp:  tv.StructFieldsType(),
		FieldRVal:  tv.StructFieldValueAll(),
		FieldVal:   tv.Struct2MapAll(),
		FieldTagKv: tv.StructParseTagKV2(_tag),
		CanSet:     tv.StructCanSet(),
	}
}

func (s *structOf) valid() bool {
	if s == nil {
		return true
	}

	for name, typ := range s.FieldRTyp {
		field := fieldOf{
			rkind: typ.Kind(),
			rval:  s.FieldRVal[name],
			val:   s.FieldVal[name],
			str:   cast.ToString(s.FieldVal[name]),
			tag:   s.FieldTagKv[name],
		}
		s.Has(field.valid())
	}
	return s.Has()
}
