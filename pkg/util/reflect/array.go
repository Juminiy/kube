package reflect

import "reflect"

func (tv TypVal) ArraySet(index int, elem any) {
	tv.noPointer()

	if tv.arrayCanOpt(elem) {
		elemOfI := tv.Val.Index(index)
		if elemOfI.CanAddr() {
			elemOfI.Set(reflect.ValueOf(elem))
		}
	}
}

func (tv TypVal) arrayCanOpt(elem any) bool {
	return tv.Typ.Kind() == reflect.Array &&
		tv.Val.Len() > 0 &&
		tv.Typ.Elem() == reflect.TypeOf(elem)
}
