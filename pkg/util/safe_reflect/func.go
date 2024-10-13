package safe_reflect

import "reflect"

// Func API
// +param fn type can indirect
// +desc Func is value not pointer

// FuncSet
// set func self to fn
func (tv TypVal) FuncSet(fn any) {
	v := tv.noPointer()

	if v.Kind() != Func || !v.CanSet() ||
		!tv.funcCanOpt(fn) {
		return
	}

	v.Set(indirectV(fn))
}

func (tv TypVal) funcCanOpt(fn any) bool {
	fnTyp := indirectT(fn)
	if tv.Typ.Kind() != Func || tv.Val.IsNil() ||
		fnTyp.Kind() != Func {
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

func (tv TypVal) funcType(in, out []any, variadic bool) reflect.Type {
	inTyp, outTyp := directTs(in), directTs(out)
	if len(in) > 0 {
		variadic = variadic && inTyp[len(in)-1].Kind() == Slice
	}
	return reflect.FuncOf(inTyp, outTyp, variadic)
}

func (tv TypVal) funcMake(in, out []any, variadic bool) any {
	v := tv.noPointer()
	if v.Kind() != Func {
		return nil
	}

	return nil
}

func (tv TypVal) funcCall(in []any) []any {
	v := tv.noPointer()

	if v.Kind() == Func && !v.IsNil() {
		return InterfacesOf(v.Call(directVs(in)))
	}
	return nil
}
