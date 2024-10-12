package safe_reflect

import (
	"reflect"
	"strings"
)

// Struct API
// +param underlyingIsStructVal type and its attribute type can indirect
// +desc reflect.Struct is value

func (tv TypVal) StructSet(underlyingIsStructVal any) {
	v := tv.noPointer()

	if v.Kind() != reflect.Struct || !v.CanSet() {
		return
	}

	indirSrcT, indirSrcV := indirectTV(underlyingIsStructVal)
	if tv.Typ == indirSrcT && v.CanSet() {
		v.Set(indirSrcV)
	}
}

func (tv TypVal) StructSetFields(fields map[string]any) {
	v := tv.noPointer()

	if v.Kind() != reflect.Struct || !v.CanSet() {
		return
	}

	for fieldName, fieldVal := range fields {
		// old-version: worked
		//srcTyp, srcVal := indirectTV(fieldVal)
		//dstTypOfStructField, dstTypOk := v.Type().FieldByName(fieldName)
		//if !dstTypOk || dstTypOfStructField.Type != srcTyp {
		//	continue
		//}
		//dstTv := indirect(v.FieldByName(fieldName))
		//if dstTv.Typ == srcTyp && dstTv.Val.CanSet() {
		//	dstTv.Val.Set(srcVal)
		//}

		// new-version: worked, optimized compare to old-version
		indirSrcT, indirSrcV := indirectTV(fieldVal) // src can indirect
		fieldIndex := tv.fieldIndexByName(fieldName)
		if len(fieldIndex) == 0 {
			continue
		}
		indirDst := indirect(v.FieldByIndex(fieldIndex))
		if indirSrcT == indirDst.Typ && indirDst.Val.CanSet() {
			indirDst.Val.Set(indirSrcV) // dst can indirect
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

	typ := tv.Typ
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

func (tv TypVal) fieldIndexByName(fieldName string) []int {
	if tv.noPointer().Kind() != reflect.Struct {
		return nil
	}
	for i := range tv.Typ.NumField() {
		if fTyp := tv.Typ.Field(i); fTyp.Name == fieldName {
			return fTyp.Index
		}
	}
	return nil
}
