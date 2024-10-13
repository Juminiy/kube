package safe_reflect

// Array API
// +param elem type can indirect
// +desc Arr is pointer but no address

// ArraySet
// set array index to elem -> arr[index] = elem
func (tv TypVal) ArraySet(index int, elem any) {
	v := tv.noPointer()

	if v.Kind() != Arr || !v.CanSet() ||
		tv.FieldLen() <= index ||
		!tv.arrayCanOpt(elem) {
		return
	}

	if indirIndexV := noPointer(v.Index(index)); indirIndexV.CanSet() {
		indirIndexV.Set(indirectV(elem))
	}
}

// ArraySetStructFields
// set array struct fields fieldName to fieldVal
func (tv TypVal) ArraySetStructFields(fields map[string]any) {
	v := tv.noPointer()

	if v.Kind() != Arr || !v.CanSet() ||
		tv.FieldLen() == 0 {
		return
	}

	for index := range tv.FieldLen() {
		indirect(v.Index(index)).StructSetFields(fields)
	}
}

func (tv TypVal) arrayCanOpt(elem any) bool {
	return tv.Typ.Kind() == Arr &&
		tv.FieldLen() > 0 &&
		underlyingEqual(tv.Typ.Elem(), directT(elem))
}
