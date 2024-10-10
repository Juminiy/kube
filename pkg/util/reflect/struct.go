package reflect

import (
	"reflect"
	"strings"
)

// ParseStructTag
// +example
// `app1:"tag_val1" app2:"tag_val2" app3:"tag_val3"`
func (tv TypVal) ParseStructTag(app string) (tagMap TagMap) {
	v := deref2NoPointer(tv.Val)

	if v.Kind() != reflect.Struct {
		return
	}

	typ := v.Type()
	tagMap = make(TagMap, typ.NumField())
	for i := range typ.NumField() {
		fieldI := typ.Field(i)
		tagMap[fieldI.Name] = fieldI.Tag.Get(app)
	}
	return
}

// TagMap in an app -> map[field]tag_val
type TagMap map[string]string

// ParseGetTagValV
// +example
// `gorm:"column:user_name;type:varchar(128);comment:user's name, account's name"`
func (m TagMap) ParseGetTagValV(field, key string) string {
	if len(m) == 0 {
		return ""
	}
	// column:user_name;type:varchar(128);comment:user's name, account's name
	kvStrs := strings.Split(m[field], ";")
	for kvI := range kvStrs {
		kvPairs := strings.Split(kvStrs[kvI], ":")
		if len(kvPairs) == 2 && kvPairs[0] == key {
			return kvPairs[1]
		}
	}
	return ""
}

func (tv TypVal) StructSetField(fields map[string]any) {
	tv.noPointer()

	for fieldName, fieldsVal := range fields {
		tv.Val.FieldByName(fieldName).Set(reflect.ValueOf(fieldsVal))
	}
}

func (tv TypVal) structCanOpt(v any) bool {
	return tv.Typ == reflect.TypeOf(v)
}
