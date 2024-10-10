package safe_reflect

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
	srcTyp, srcVal := indirectTV(src)
	dstTyp, dstVal := indirectTV(dst)

	srcFieldMap := make(map[string]TypVal, fieldLen(srcVal))

	switch {
	case srcTyp.Kind() == reflect.Struct &&
		dstTyp.Kind() == reflect.Struct &&
		dstVal.CanSet():
		for index := range srcTyp.NumField() {
			srcFi := srcTyp.Field(index)
			srcFieldMap[srcFi.Name] = TypVal{
				Typ: srcFi.Type,
				Val: srcVal.Field(index),
			}
		}

		for index := range dstTyp.NumField() {
			dstFi := dstTyp.Field(index)
			if fieldV, ok := srcFieldMap[dstFi.Name]; ok && fieldV.Typ == dstFi.Type {
				dstValFieldI := dstVal.Field(index)
				if dstValFieldI.CanSet() {
					dstValFieldI.Set(fieldV.Val)
				}
			}
		}

	}
}
