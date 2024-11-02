package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
	"strings"
)

// Struct API
// +param srcV type and its attribute type can indirect
// +desc Struct is value

func (tv TypVal) StructSet(srcV any) {
	v := tv.noPointer()

	if v.Kind() != Struct || !v.CanSet() {
		return
	}

	indirSrcT, indirSrcV := indirectTV(srcV)
	if tv.Typ == indirSrcT && v.CanSet() {
		v.Set(indirSrcV)
	}
}

// StructSetFields
// +param:
// fieldName -> fieldValue
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

// StructFieldsIndex
// +result:
// fieldName -> fieldIndex
func (tv TypVal) StructFieldsIndex() map[string][]int {
	v := tv.noPointer()

	if v.Kind() != Struct {
		return nil
	}

	typ := tv.Typ
	indexMap := make(map[string][]int, typ.NumField())
	for i := range typ.NumField() {
		fieldI := typ.Field(i)
		indexMap[fieldI.Name] = fieldI.Index
	}
	return indexMap
}

// StructFieldsType
// +result:
// fieldName -> fieldType
func (tv TypVal) StructFieldsType() map[string]reflect.Type {
	v := tv.noPointer()

	if v.Kind() != Struct {
		return nil
	}

	return structFieldsTypeMap(tv.Typ)
}

// no check Kind
func structFieldsTypeMap(typ reflect.Type) map[string]reflect.Type {
	fieldsTypeMap := make(map[string]reflect.Type, typ.NumField())
	for i := range typ.NumField() {
		fieldI := typ.Field(i)
		fieldsTypeMap[fieldI.Name] = fieldI.Type
	}
	return fieldsTypeMap
}

func (tv TypVal) StructFieldValue(fieldName string) any {
	v := tv.noPointer()

	if v.Kind() != Struct {
		return nil
	}

	fieldValue := v.FieldByName(fieldName)
	if fieldValue != _zeroValue && fieldValue.CanInterface() {
		return fieldValue.Interface()
	}
	return nil
}

func (tv TypVal) StructFieldsValues(fieldIndex map[string][]int) map[string]map[any]struct{} {
	v := tv.noPointer()

	if v.Kind() != Struct {
		return nil
	}

	// all field list
	fieldsIndex := tv.StructFieldsIndex()
	// common field list
	util.MapEvict(fieldsIndex, fieldIndex)

	fieldsValues := make(map[string]map[any]struct{}, len(fieldsIndex))
	for fieldName, fieldIndex := range fieldsIndex {
		if fi := v.FieldByIndex(fieldIndex); fi.CanInterface() {
			fieldsValues[fieldName] = map[any]struct{}{
				fi.Interface(): {},
			}
		}

	}
	return fieldsValues
}

/*func (tv TypVal) StructFieldsStrValues(fieldsIndex map[string][]int) map[string]string {
	v := tv.noPointer()

	if v.Kind() != Struct {
		return nil
	}

	fieldsValues := make(map[string]string, len(fieldsIndex))
	for fieldName, fieldIndex := range fieldsIndex {
		fieldI := v.FieldByIndex(fieldIndex)
		if fieldI.Kind() == String {
			fieldsValues[fieldName] = fieldI.Interface().(string)
		}
	}

	return fieldsValues
}*/

func (tv TypVal) Struct2Map(fields map[string]struct{}) map[string]any {
	v := tv.noPointer()

	if v.Kind() != Struct {
		return nil
	}

	// all field list
	fieldsIndex := tv.StructFieldsIndex()
	// common field list
	util.MapEvict(fieldsIndex, fields)

	structMap := make(map[string]any, tv.FieldLen())
	for fieldName, fieldIndex := range fieldsIndex {
		if fi := v.FieldByIndex(fieldIndex); fi.CanInterface() {
			structMap[fieldName] = fi.Interface()
		}
	}
	return structMap
}

// StructHasFields
// match all fields by FieldName and FieldType
// fields FieldValue must direct
func (tv TypVal) StructHasFields(fields map[string]any) map[string]struct{} {
	v := tv.noPointer()

	if v.Kind() != Struct {
		return nil
	}

	return structHasFields(tv.Typ, fields)
}

func structHasFields(typ reflect.Type, fields map[string]any) map[string]struct{} {
	fieldsTypeMap := structFieldsTypeMap(typ)
	if len(fieldsTypeMap) == 0 {
		return nil
	}
	okMap := make(map[string]struct{}, len(fieldsTypeMap))
	for fieldName, fieldValue := range fields {
		if fieldTyp, ok := fieldsTypeMap[fieldName]; ok { // has FieldName
			if directT(fieldValue) == fieldTyp { // direct match FiledType
				okMap[fieldName] = struct{}{}
			}
		}
	}
	return okMap
}

// Struct2TagKeyMap
// +example
// `app:"unique:1;union_unique:0;field:name;"`
// `app:"unique:0;union_unique:1;field:name_part1;follow:-"`
// `app:"unique:0;union_unique:1;field:name_part1;follow:+"`
func (tv TypVal) Struct2TagKeyMap(app, key string) map[string]any {
	v := tv.noPointer()
	if v.Kind() != Struct {
		return nil
	}
	tagKeyMap := make(map[string]any, tv.FieldLen())
	tagMap := tv.StructParseTag(app)
	for fieldName := range tagMap {
		if tagKey := tagMap.ParseGetVal(fieldName, key); len(tagKey) > 0 {
			tagKeyMap[tagKey] = tv.StructFieldValue(fieldName)
		}
	}
	return tagKeyMap
}

// StructParseTag
// +example
// `app1:"tag_val1" app2:"tag_val2" app3:"tag_val3"`
func (tv TypVal) StructParseTag(app string) (tagMap TagMap) {
	v := tv.noPointer()
	if v.Kind() != Struct {
		return
	}

	return structParseTag(tv.Typ, app)
}

func structParseTag(typ reflect.Type, app string) (tagMap TagMap) {
	tagMap = make(TagMap, typ.NumField())
	for i := range typ.NumField() {
		fieldI := typ.Field(i)
		tagMap[fieldI.Name] = fieldI.Tag.Get(app)
	}
	return
}

func (tv TypVal) StructParseTag2(app0, key0, app1, key1 string) (vv0, vv1 TagVV) {
	if tv.noPointer().Kind() != Struct {
		return
	}

	return structParseTag2(tv.Typ, app0, key0, app1, key1)
}

type TagVV map[string]string

// Name string `gorm:"column:name" app:"field:af"`
// Name -> name; Name -> af
func structParseTag2(typ reflect.Type, srcApp, srcKey, dstApp, dstKey string) (srcMap, dstMap TagVV) {
	srcMap, dstMap = make(TagVV, typ.NumField()), make(TagVV, typ.NumField())
	for i := range typ.NumField() {
		fieldI := typ.Field(i)
		srcv := parseTagKV(fieldI.Tag.Get(srcApp))[srcKey]
		dstv := parseTagKV(fieldI.Tag.Get(dstApp))[dstKey]

		//if len(srcv) == 0 {
		//	srcv = srcKey
		//}
		//if len(dstv) == 0 {
		//	dstv = dstKey
		//}

		srcMap[fieldI.Name] = srcv
		dstMap[fieldI.Name] = dstv
	}
	return
}

// StructParseTagKV
// +param: an app
// +result: FieldTagKV
func (tv TypVal) StructParseTagKV(app string) (fieldTagKv FieldTagKV) {
	if tv.noPointer().Kind() != Struct {
		return
	}

	return structParseTagKV(tv.Typ, app)
}

// FieldTagKV
// fieldName -> TagKV
type FieldTagKV map[string]TagKV

func structParseTagKV(typ reflect.Type, app string) (fieldTagKv FieldTagKV) {
	fieldTagKv = make(FieldTagKV, typ.NumField())
	for i := range typ.NumField() {
		field := typ.Field(i)

		fieldTagKv[field.Name] = parseTagKV(field.Tag.Get(app))
	}
	return
}

// TagKV
// as parseTagKV description
// +example:
// `app:"k1:v1;k2:v2;k3:v3;key;val"`
// +result:
// k1 -> v1
// k2 -> v2
// k3 -> v3
// key -> key
// val -> val
type TagKV map[string]string

// +example:
// `app:"k1:v1;k2:v2;k3:v3;key;val"`
// +result:
// k1 -> v1
// k2 -> v2
// k3 -> v3
// key -> key
// val -> val
func parseTagKV(tagValue string) (tagKv TagKV) {
	tagKv = make(TagKV, util.MagicMapCap)

	kvs := strings.Split(tagValue, ";")
	for _, kv := range kvs {
		keyVal := strings.Split(kv, ":")
		switch len(keyVal) {
		case 0:

		case 1:
			tagKv[keyVal[0]] = keyVal[0]

		case 2:
			tagKv[keyVal[0]] = keyVal[1]

		default:
			tagKv[keyVal[0]] = util.StringJoin(":", keyVal[1:]...)
		}
	}

	return
}

// TagMap
// in an app, fieldName -> fieldTag
// +example:
// Field0 string `app:"k1:v1;k2:v2;k3:v3;key;val"`
// +result:
// Field0 -> k1:v1;k2:v2;k3:v3;key;val
type TagMap map[string]string

// ParseGetVal
// get val in an app's, field's, key's val
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

// GetValList
// +example:
// Field0 string `app:"k1:v1;k2:v2;k3:v3;key;val"`
// Field1 string `app:"k1:v4;key-val"`
// +param:
// key = k1
// +result:
// []string{v1, v4}
func (m TagMap) GetValList(key string) []string {
	valList := make([]string, 0, len(m))
	for fieldName := range m {
		valList = append(valList, m.ParseGetVal(fieldName, key))
	}
	return valList
}

func StructMake(fields []FieldDesc) any {
	structTyp := structType(fields)
	if structTyp == nil {
		return nil
	}
	return NewOf(structTyp)
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
