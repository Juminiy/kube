package reflect

import "reflect"

func (tv TypVal) SliceSet(index int, elem any) {
	tv.Val = tv.noPointer()

	if tv.sliceCanOpt(elem) {
		tv.Val.Index(index).Set(reflect.ValueOf(elem))
	}
}

func (tv TypVal) SliceSetStructField(fields map[string]any) {
	tv.Val = tv.noPointer()

}

func (tv TypVal) sliceCanOpt(elem any) bool {
	return tv.Typ.Kind() == reflect.Slice &&
		!tv.Val.IsNil() &&
		tv.Typ.Elem() == reflect.TypeOf(elem)
}
