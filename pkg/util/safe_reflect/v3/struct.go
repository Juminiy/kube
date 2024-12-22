package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"reflect"
	"strings"
	"sync"
)

type Fields map[string]reflect.StructField
type Types map[string]reflect.Type
type Tags map[string]Tag
type Tag map[string]string

func (t T) StructTags(tagKey string) Tags {
	return lo.MapValues(t.StructFields(),
		func(field reflect.StructField, name string) Tag {
			return parseTagVal(field.Tag.Get(tagKey))
		})
}

func parseTagVal(s string) (tag Tag) {
	defer func() { util.MapDelete(tag, "") }()
	tag = parseTagValKV(s)
	if len(tag) > 0 {
		if _, ok := tag[s]; len(tag) == 1 && ok {
			goto tagValL
		}
		return
	}
tagValL:
	tag = parseTagValL(s)
	return
}

// k:"k1:v1;k2:v2;k3;k4"  -> Tag{"k1": "v1", "k2": "v2", "k3": "", "k4": ""}
// k:"k1:-;k2:v1,v2;"	  -> Tag{"k1": "-", "k2": "v1,v2"}
// k:":v1;k1:;:;k2:v2:v3" -> Tag{"k1": "", "k2": "v2:v3"}
func parseTagValKV(s string) Tag {
	kvs := strings.Split(s, ";")
	tag := make(Tag, len(kvs))
	for _, kv := range kvs {
		if len(kv) == 0 {
			continue
		}
		colonI := strings.Index(kv, ":")
		if colonI == -1 {
			tag[kv] = ""
		} else if colonI == 0 {

		} else {
			tag[kv[:colonI]] = kv[colonI+1:]
		}
	}
	return tag
}

// json:"name,omitempty" -> Tag{"name": "", "omitempty": ""}
// json:",inline"		 -> Tag{"inline": ""}
func parseTagValL(s string) Tag {
	return lo.SliceToMap(strings.Split(s, ","), func(e string) (string, string) {
		return e, ""
	})
}

func (t T) StructTypes() Types {
	return lo.MapValues(t.StructFields(),
		func(field reflect.StructField, key string) reflect.Type {
			return field.Type
		})
}

func (t T) StructFields() Fields {
	if fields, ok := _Fields.Load(t.Type); ok {
		return fields.(Fields)
	}
	fields := t.structFields()
	_Fields.Store(t.Type, fields)
	return fields
}

func (t T) structFields() Fields {
	if t.Kind() != reflect.Struct {
		return nil
	}
	fields := make(Fields, t.NumField())
	for i := range t.NumField() {
		field := t.Field(i)
		fields[field.Name] = field
	}
	return fields
}

var _Fields = sync.Map{}

func (v V) StructSet(nv map[string]any) {
	if v.Kind() != reflect.Struct || !v.CanSet() {
		return
	}
	for name, val := range nv {
		field := v.FieldByName(name)
		if field == ZeroValue() || !field.CanSet() {
			continue
		}
		V2Wrap(field).SetILike(val)
	}
}

type Values map[string]reflect.Value

func (tv Tv) StructToMap() map[string]any {
	return tv.StructValues()
}

func (tv Tv) StructValues() map[string]any {
	return lo.MapValues(tv.structValues(),
		func(rv reflect.Value, name string) any {
			return Any(rv)
		})
}

func (tv Tv) structValues() Values {
	t, v := tv.T, tv.V
	if t.Kind() != reflect.Struct {
		return nil
	}
	return lo.MapValues(t.StructFields(),
		func(field reflect.StructField, name string) reflect.Value {
			return v.FieldByName(name)
		})
}
