package safe_reflect

import "reflect"

// Slice API elem type can indirect
// reflect.Slice is pointer

// SliceSet
// set slice index to elem -> slice[index] = elem
func (tv TypVal) SliceSet(index int, elem any) {
	v := tv.noPointer()

	if tv.sliceCanOpt(elem) && tv.FieldLen() > index {
		elemV := noPointer(v.Index(index))
		if elemV.CanSet() {
			elemV.Set(indirectV(elem))
		}
	}
}

// SliceSetStructFields
// set slice struct fields fieldName to fieldVal
func (tv TypVal) SliceSetStructFields(fields map[string]any) {
	v := tv.noPointer()
	if v.Kind() != reflect.Slice ||
		!v.CanSet() {
		return
	}

	for index := range tv.FieldLen() {
		indirect(v.Index(index)).StructSetFields(fields)
	}
}

func (tv TypVal) sliceCanOpt(elem any) bool {
	return tv.Typ.Kind() == reflect.Slice &&
		!tv.Val.IsNil() &&
		underlyingEqual(tv.Typ.Elem(), reflect.TypeOf(elem))
}

// SliceSetOob
// slice set out of bound index to elem -> slice[index] = elem
func (tv TypVal) SliceSetOob(index int, elem any) {
	v := tv.noPointer()
	if v.Kind() != reflect.Slice {
		return
	}

	if tv.FieldLen() <= index {
		v.SetLen(index + 1)
	}

	tv.SliceSet(index, elem)
}
