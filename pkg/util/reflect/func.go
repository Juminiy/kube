package reflect

import "reflect"

func (tv TypVal) FuncSet(v any) {
	tv.noPointer()

	if tv.funcCanOpt(v) {
		tv.Val.Set(reflect.ValueOf(v))
	}
}

func (tv TypVal) funcCanOpt(v any) bool {
	vTyp := reflect.TypeOf(v)
	if tv.Typ.Kind() != reflect.Func || tv.Val.IsNil() || vTyp.Kind() != reflect.Func {
		return false
	}

	if tv.Typ.NumIn() != vTyp.NumIn() ||
		tv.Typ.NumOut() != vTyp.NumOut() {
		return false
	}
	for inI := range tv.Typ.NumIn() {
		if tv.Typ.In(inI) != vTyp.In(inI) {
			return false
		}
	}
	for outI := range tv.Typ.NumOut() {
		if tv.Typ.Out(outI) != vTyp.Out(outI) {
			return false
		}
	}
	return true
}
