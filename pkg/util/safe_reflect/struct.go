package safe_reflect

import (
	"reflect"
	"strings"
)

// Struct API underlyingIsStructVal type and its attribute type can indirect
// reflect.Struct is value

func (tv TypVal) StructSet(underlyingIsStructVal any) {
	v := tv.noPointer()
	if v.Kind() != reflect.Struct {
		return
	}

	vTyp, vVal := indirectTV(underlyingIsStructVal)
	if tv.Typ == vTyp && v.CanSet() {
		v.Set(vVal)
	}
}

func (tv TypVal) StructSetFields(fields map[string]any) {
	v := tv.noPointer()
	if v.Kind() != reflect.Struct {
		return
	}

	for fieldName, fieldVal := range fields {
		srcTyp, srcVal := indirectTV(fieldVal)
		dstVal := deref2NoPointer(v.FieldByName(fieldName))
		if dstVal.Type() == srcTyp && dstVal.CanSet() {
			dstVal.Set(srcVal)
		}
	}
}

// ParseStructTag
// +example
// `app1:"tag_val1" app2:"tag_val2" app3:"tag_val3"`
func (tv TypVal) ParseStructTag(app string) (tagMap TagMap) {
	v := tv.noPointer()
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
