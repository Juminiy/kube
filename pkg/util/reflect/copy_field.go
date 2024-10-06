package reflect

import (
	"reflect"
)

// CopyFieldValue
// copy field value in src to dst which fields with same name and same type
// regardless src type same with dst type
// no deep copy, only one level iface
// +case 1
// src can be Struct, Struct Pointer, Struct Pointer Pointer, Struct Pointer ..., ...
// dst must be Struct Pointer, Struct Pointer Pointer, Struct Pointer ..., ...
func CopyFieldValue(src any, dst any) {
	srcVal := deref2NoPointer(reflect.ValueOf(src))
	dstVal := deref2OnePointer(reflect.ValueOf(dst))
	srcTyp := srcVal.Type()

	srcFieldMap := make(map[string]typVal, fieldSize(srcVal))

	switch {
	case srcVal.Kind() == reflect.Struct &&
		underlyingStruct(dstVal) &&
		dstVal.Kind() == reflect.Pointer: // dstVal.CanAddr() && dstVal.CanSet()
		for srcI := range srcTyp.NumField() {
			srcFi := srcTyp.Field(srcI)
			srcFieldMap[srcFi.Name] = typVal{
				typ: srcFi.Type,
				val: srcVal.Field(srcI), // srcVal.FieldByName(srcFi.Name)
			}
		}

		dstTyp := reflect.Indirect(dstVal).Type()
		dstValElem := dstVal.Elem()
		for dstI := range dstTyp.NumField() {
			dstFi := dstTyp.Field(dstI)
			if fieldV, ok := srcFieldMap[dstFi.Name]; ok && fieldV.typ == dstFi.Type {
				dstValElemFi := dstValElem.Field(dstI)
				if dstValElemFi.CanSet() {
					dstValElemFi.Set(fieldV.val)
				}
			}
		}

	}
}

type typVal struct {
	typ reflect.Type
	val reflect.Value
}
