package safe_reflect

import "reflect"

// Slice API elem type can indirect
// reflect.Slice is pointer

func (tv TypVal) SliceSet(index int, elem any) {
	v := tv.noPointer()

	if tv.sliceCanOpt(elem) && tv.FieldLen() > index {
		elemV := noPointer(v.Index(index))
		if elemV.CanSet() {
			elemV.Set(indirectV(elem))
		}
	}
}

func (tv TypVal) SliceSetStructFields(fields map[string]any) {
	v := tv.noPointer()
	if tv.Typ.Kind() != reflect.Slice &&
		!tv.Val.CanSet() {
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
