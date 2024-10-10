package safe_reflect

import "reflect"

// Func API fn type can indirect
// reflect.Func is value not pointer

func (tv TypVal) FuncSet(fn any) {
	v := tv.noPointer()

	if tv.funcCanOpt(fn) && v.CanSet() {
		v.Set(indirectV(fn))
	}
}

func (tv TypVal) funcCanOpt(fn any) bool {
	fnTyp := indirectT(fn)
	if tv.Typ.Kind() != reflect.Func || tv.Val.IsNil() || fnTyp.Kind() != reflect.Func {
		return false
	}

	if tv.Typ.NumIn() != fnTyp.NumIn() ||
		tv.Typ.NumOut() != fnTyp.NumOut() {
		return false
	}
	for inI := range tv.Typ.NumIn() {
		if tv.Typ.In(inI) != fnTyp.In(inI) {
			return false
		}
	}
	for outI := range tv.Typ.NumOut() {
		if tv.Typ.Out(outI) != fnTyp.Out(outI) {
			return false
		}
	}
	return true
}
