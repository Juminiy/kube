package safe_reflect

import "reflect"

// Array API elem type can indirect
// reflect.Array is value not pointer

func (tv TypVal) ArraySet(index int, elem any) {
	v := tv.noPointer()

	if tv.arrayCanOpt(elem) && tv.FieldLen() > index {
		elemI := deref2NoPointer(v.Index(index))
		if elemI.CanSet() {
			elemI.Set(indirectV(elem))
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
		of(v.Index(index)).StructSetFields(fields)
	}
}

func (tv TypVal) arrayCanOpt(elem any) bool {
	return tv.Typ.Kind() == reflect.Array &&
		tv.FieldLen() > 0 &&
		underlyingTypeEq(tv.Typ.Elem(), reflect.TypeOf(elem))
}
