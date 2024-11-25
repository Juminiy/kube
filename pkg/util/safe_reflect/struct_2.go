package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
	"strings"
)

// StructParseTagKV2
// +param: an app
// +result: FieldTagKV
func (tv TypVal) StructParseTagKV2(app string) (fieldTagKv FieldTagKV) {
	if tv.noPointer().Kind() != Struct {
		return
	}

	return structParseTagKV2(tv.Typ, app)
}

func structParseTagKV2(typ reflect.Type, app string) (fieldTagKv FieldTagKV) {
	fieldTagKv = make(FieldTagKV, typ.NumField())
	for i := range typ.NumField() {
		field := typ.Field(i)

		fieldTagKv[field.Name] = parseTagKV2(field.Tag.Get(app))
	}
	return
}

func parseTagKV2(tagValue string) (tagKv TagKV) {
	tagKv = make(TagKV, util.MagicMapCap)

	tagKv.parseSemicolonAndColon2(tagValue)

	return
}

// parseTagKVSemicolonAndColon
// +example:
// `app:"k1:v1;k2:v2;k3:v3;key;val;k4:v4:v5:v6"`
// +result:
// k1 -> v1
// k2 -> v2
// k3 -> v3
// key -> ""
// val -> ""
// no k4
func (tagKv TagKV) parseSemicolonAndColon2(tagValue string) {
	kvs := strings.Split(tagValue, ";")
	for _, kv := range kvs {
		if len(kv) == 0 { // skip ""
			continue
		}
		keyVal := strings.Split(kv, ":")
		switch len(keyVal) {
		case 0:

		case 1:
			tagKv[keyVal[0]] = ""

		case 2:
			tagKv[keyVal[0]] = keyVal[1]

		default:

		}
	}
}
