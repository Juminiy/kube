package safe_validator

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/spf13/cast"
	"reflect"
)

// Struct
// compatible with fiber.StructValidator
func (cfg *Config) Struct(v any) bool {
	ok, _ := cfg.structOkE(v)
	return ok
}

func (cfg *Config) StructE(v any) error {
	_, err := cfg.structOkE(v)
	return err
}

func (cfg *Config) structOkE(v any) (ok bool, err error) {
	parsed := cfg.parseStruct(v)
	if parsed == nil {
		return false, errValNotStruct
	}
	ok, err = parsed.valid(), parsed.All()
	return
}

var errValNotStruct = errors.New("value type is not struct")

type structOf struct {
	FieldRTyp  map[string]reflect.Type
	FieldRVal  map[string]reflect.Value
	FieldVal   map[string]any
	FieldTagKv safe_reflect.FieldTagKV
	*util.ErrHandle

	cfg *Config
}

func (cfg *Config) parseStruct(v any) *structOf {
	tv := indir(v)
	if tv.Typ.Kind() != kStruct {
		return nil
	}
	return &structOf{
		FieldRTyp:  tv.StructFieldsType(),
		FieldRVal:  tv.StructFieldValueAll(),
		FieldVal:   tv.Struct2MapAll(),
		FieldTagKv: tv.StructParseTagKV2(cfg.Tag),
		ErrHandle:  util.NewErrHandle(),
		cfg:        cfg,
	}
}

func (s *structOf) valid() bool {
	if s == nil {
		return true
	}

	for name, typ := range s.FieldRTyp {
		field := fieldOf{
			name:  name,
			rkind: typ.Kind(),
			rval:  s.FieldRVal[name],
			val:   s.FieldVal[name],
			str:   cast.ToString(s.FieldVal[name]),
			tag:   s.FieldTagKv[name],
			cfg:   s.cfg,
		}
		if s.Has(field.valid()) && s.cfg.OnErrorStop {
			return false
		}
	}
	return !s.Has()
}
