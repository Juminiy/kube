package safe_reflect

import (
	"reflect"
)

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
	return tv.funcInOutEq(funcInType(fnTyp), funcOutType(fnTyp))
}

func (tv TypVal) funcInOutEq(in, out []reflect.Type) bool {
	typ := tv.Typ
	if typ.NumIn() != len(in) ||
		typ.NumOut() != len(out) {
		return false
	}
	for i := range typ.NumIn() {
		if typ.In(i) != in[i] {
			return false
		}
	}
	for i := range typ.NumOut() {
		if typ.Out(i) != out[i] {
			return false
		}
	}
	return true
}

func funcInType(typ reflect.Type) []reflect.Type {
	inTyp := make([]reflect.Type, typ.NumIn())
	for i := range typ.NumIn() {
		inTyp[i] = typ.In(i)
	}
	return inTyp
}

func funcOutType(typ reflect.Type) []reflect.Type {
	outTyp := make([]reflect.Type, typ.NumOut())
	for i := range typ.NumOut() {
		outTyp[i] = typ.Out(i)
	}
	return outTyp
}

func (tv TypVal) FuncCall(in []any) ([]any, bool) {
	v := tv.noPointer()
	if v.Kind() != Func || v.IsNil() ||
		tv.Typ.NumIn() != len(in) {
		return nil, false
	}
	inTs := directTs(in)
	for i := range inTs {
		if inTs[i] != tv.Typ.In(i) {
			return nil, false
		}
	}
	return InterfacesOf(v.Call(directVs(in))), true
}

func (tv TypVal) HasMethod(methodName string, in, out []any) bool {
	tv.noPointer()

	typ := tv.Typ
	method, ok := typ.MethodByName(methodName)
	if !ok {
		return false
	}
	return direct(method.Func).funcInOutEq(directTs(in), directTs(out))
}

func FuncMake(in, out []any, variadic bool, metaFunc MetaFunc) any {
	if metaFunc == nil {
		return nil
	}
	return reflect.MakeFunc(funcType(in, out, variadic), metaFunc).Interface()
}

type MetaFunc func([]reflect.Value) []reflect.Value

func funcType(in, out []any, variadic bool) reflect.Type {
	inTyp, outTyp := directTs(in), directTs(out)
	if len(in) > 0 {
		variadic = variadic && inTyp[len(in)-1].Kind() == Slice
	}
	return reflect.FuncOf(inTyp, outTyp, variadic)
}
