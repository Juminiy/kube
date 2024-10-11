package safe_reflect

import "reflect"

// Array API elem type can indirect
// reflect.Array is value not pointer

func (tv TypVal) ArraySet(index int, elem any) {
	v := tv.noPointer()

	if tv.arrayCanOpt(elem) && tv.FieldLen() > index {
		elemV := noPointer(v.Index(index))
		if elemV.CanSet() {
			elemV.Set(indirectV(elem))
		}
	}
}

func (tv TypVal) ArraySetStructFields(fields map[string]any) {
	v := tv.noPointer()
	if tv.Typ.Kind() != reflect.Array &&
		!v.CanSet() {
		return
	}

	for index := range tv.FieldLen() {
		indirect(v.Index(index)).StructSetFields(fields)
	}
}

func (tv TypVal) arrayCanOpt(elem any) bool {
	return tv.Typ.Kind() == reflect.Array &&
		tv.FieldLen() > 0 &&
		underlyingEqual(tv.Typ.Elem(), reflect.TypeOf(elem))
}
