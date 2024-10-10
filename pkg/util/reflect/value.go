package reflect

import (
	"reflect"
)

// CopyFieldValue
// copy field value in src to dst which fields with same name and same type
// regardless src type same with dst type
// no deep copy, only one level iface
// +case 1
// src can be Struct, Struct Pointer, Struct Pointer to Pointer, Struct Pointer to..., ...
// dst must be Struct Pointer, Struct Pointer to Pointer, Struct Pointer to..., ...
func CopyFieldValue(src any, dst any) {
	srcVal := deref2NoPointer(reflect.ValueOf(src))
	dstVal := deref2OnePointer(reflect.ValueOf(dst))
	srcTyp := srcVal.Type()

	srcFieldMap := make(map[string]TypVal, fieldLen(srcVal))

	switch {
	case srcVal.Kind() == reflect.Struct &&
		underlyingIsStruct(dstVal) &&
		dstVal.Kind() == reflect.Pointer: // dstVal.CanAddr() && dstVal.CanSet()
		for srcI := range srcTyp.NumField() {
			srcFi := srcTyp.Field(srcI)
			srcFieldMap[srcFi.Name] = TypVal{
				Typ: srcFi.Type,
				Val: srcVal.Field(srcI), // srcVal.FieldByName(srcFi.Name)
			}
		}

		dstTyp := reflect.Indirect(dstVal).Type()
		dstValElem := dstVal.Elem()
		for dstI := range dstTyp.NumField() {
			dstFi := dstTyp.Field(dstI)
			if fieldV, ok := srcFieldMap[dstFi.Name]; ok && fieldV.Typ == dstFi.Type {
				dstValElemFi := dstValElem.Field(dstI)
				if dstValElemFi.CanSet() {
					dstValElemFi.Set(fieldV.Val)
				}
			}
		}

	}
}
