package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
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
	case srcTyp.Kind() == Struct &&
		dstTyp.Kind() == Struct &&
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

func directSet(src, dst any) {
	sOf, dOf := Of(src), Of(dst)
	if sOf.Typ == dOf.Typ && dOf.Val.CanSet() {
		dOf.Val.Set(sOf.Val)
	}
}

func indirectSet(src, dst any) {
	sOf, dOf := IndirectOf(src), IndirectOf(dst)
	if sOf.Typ == dOf.Typ && dOf.Val.CanSet() {
		dOf.Val.Set(sOf.Val)
	}
}

var (
	_zeroValue = reflect.Value{} // comparable
)

func (tv TypVal) FieldLen() int {
	return fieldLen(tv.Val)
}

func fieldLen(v reflect.Value) int {
	switch v.Kind() {
	case Struct:
		return v.NumField()
	case Arr, Chan, Map, Slice, String:
		return v.Len()
	default:
		return 0
	}
}

func valueCanElem(v reflect.Value) bool {
	return util.ElemIn(v.Kind(),
		Ptr, Any,
	)
}

// a type's zero value
// return a reflect.Value has an abi.Type and flag, not the reflect.Value{}
func typeZeroValue(v any) reflect.Value {
	return reflect.Zero(directT(v))
}
