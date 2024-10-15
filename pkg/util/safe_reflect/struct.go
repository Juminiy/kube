package safe_reflect

import (
	"reflect"
	"strings"
)

// Struct API
// +param underlyingIsStructVal type and its attribute type can indirect
// +desc Struct is value

func (tv TypVal) StructSet(underlyingIsStructVal any) {
	v := tv.noPointer()

	if v.Kind() != Struct || !v.CanSet() {
		return
	}

	indirSrcT, indirSrcV := indirectTV(underlyingIsStructVal)
	if tv.Typ == indirSrcT && v.CanSet() {
		v.Set(indirSrcV)
	}
}

func (tv TypVal) StructSetFields(fields map[string]any) {
	v := tv.noPointer()

	if v.Kind() != Struct || !v.CanSet() {
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
		fieldIndex := tv.structFieldIndexByName(fieldName)
		if len(fieldIndex) == 0 {
			continue
		}
		indirDst := indirect(v.FieldByIndex(fieldIndex))
		if indirSrcT == indirDst.Typ && indirDst.Val.CanSet() {
			indirDst.Val.Set(indirSrcV) // dst can indirect
		}
	}
}

func (tv TypVal) structFieldIndexByName(fieldName string) []int {
	if tv.noPointer().Kind() != Struct {
		return nil
	}
	for i := range tv.Typ.NumField() {
		if fTyp := tv.Typ.Field(i); fTyp.Name == fieldName {
			return fTyp.Index
		}
	}
	return nil
}

// StructParseTag
// +example
// `app1:"tag_val1" app2:"tag_val2" app3:"tag_val3"`
func (tv TypVal) StructParseTag(app string) (tagMap TagMap) {
	v := tv.noPointer()

	if v.Kind() != Struct {
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

// ParseGetVal
// +example
// `gorm:"column:user_name;type:varchar(128);comment:user's name, account's name"`
// +example
// `app:"unique:1;union_unique:0;field:name;"`
// `app:"unique:0;union_unique:1;field:name_part1;follow:-"`
// `app:"unique:0;union_unique:1;field:name_part1;follow:+"`
func (m TagMap) ParseGetVal(field, key string) string {
	if len(m) == 0 {
		return ""
	}
	// column:user_name;type:varchar(128);comment:user's name, account's name
	kvStrs := strings.Split(m[field], ";")
	for kvI := range kvStrs {
		//kvPairs := strings.Split(kvStrs[kvI], ":")
		//if len(kvPairs) == 2 && kvPairs[0] == key {
		//	return kvPairs[1]
		//}

		// optimize to apply example like `app:"unique:1;union_unique:0;field:name;follow::"`
		kvIStr := kvStrs[kvI]
		firstColonIndex := strings.Index(kvIStr, ":")
		if firstColonIndex != -1 && // find split-Colon(:)
			kvIStr[:firstColonIndex] == key && // find key match
			len(kvIStr) >= firstColonIndex+1 { // can get value
			return kvIStr[firstColonIndex+1:]
		}
	}
	return ""
}

func StructMake(fields []FieldDesc) any {
	structTyp := structType(fields)
	if structTyp == nil {
		return nil
	}
	return reflect.New(structTyp).Elem().Interface()
}

func structType(fields []FieldDesc) reflect.Type {
	structFields := make([]reflect.StructField, len(fields))
	fieldSet := make(map[string]struct{}, len(fields))
	for i := range fields {
		if _, ok := fieldSet[fields[i].Name]; ok {
			return nil
		}
		structFields[i] = fields[i].StructField()
		fieldSet[fields[i].Name] = struct{}{}
	}
	return reflect.StructOf(structFields)
}

type FieldDesc struct {
	Name  string
	Value any
	Tag   reflect.StructTag
}

func (f FieldDesc) StructField() reflect.StructField {
	return reflect.StructField{
		Name: f.Name,
		Type: directT(f.Value),
		Tag:  f.Tag,
	}
}
