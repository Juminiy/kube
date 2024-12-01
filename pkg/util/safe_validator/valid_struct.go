package safe_validator

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/samber/lo"
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
	if tv.Val.Kind() != kStruct {
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
		if len(s.FieldTagKv[name]) == 0 {
			continue
		}
		field := fieldOf{
			name:  name,
			rkind: typ.Kind(),
			rval:  s.FieldRVal[name],
			val:   s.FieldVal[name],
			str:   toString(s.FieldVal[name]),
			tag:   s.FieldTagKv[name],
			cfg:   s.cfg,
		}
		field.tag = lo.MapKeys(field.tag, func(_tagv string, _tagk string) (_tagkNew string) {
			return parseTagK(_tagk)
		})
		if tagk0, tagk1 := field.tagConflict(); len(tagk0) > 0 && len(tagk1) > 0 {
			if s.cfg.IgnoreTagError {
				continue
			}
			s.Has(field.errTagConflict(tagk0, tagk1))
		}

		if err := field.valid(); err != nil {
			if errIsTagFormat(err) && s.cfg.IgnoreTagError {
				continue
			}
			s.Has(err)
		}
		if s.Has() && s.cfg.OnErrorStop {
			return false
		}
	}
	return !s.Has()
}
