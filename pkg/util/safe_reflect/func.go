package safe_reflect

import "reflect"

// Func API
// +param fn type can indirect
// +desc reflect.Func is value not pointer

// FuncSet
// set func self to fn
func (tv TypVal) FuncSet(fn any) {
	v := tv.noPointer()

	if v.Kind() != reflect.Func || !v.CanSet() ||
		!tv.funcCanOpt(fn) {
		return
	}

	v.Set(indirectV(fn))
}

func (tv TypVal) funcCanOpt(fn any) bool {
	fnTyp := indirectT(fn)
	if tv.Typ.Kind() != reflect.Func || tv.Val.IsNil() ||
		fnTyp.Kind() != reflect.Func {
		return false
	}

	return tv.funcCanOptSlow(fnTyp)
}

func (tv TypVal) funcCanOptSlow(fnTyp reflect.Type) bool {
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
