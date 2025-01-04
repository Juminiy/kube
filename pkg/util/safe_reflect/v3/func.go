package safe_reflectv3

import (
	"reflect"
)

func (v V) CallMethod(name string, args []any) (rets []any, called bool) {
	method := v.MethodByName(name)
	if method == _ZeroValue {
		return
	}
	methodDirect := Wrap(method)
	rets, called = methodDirect.FuncCall(args)
	if !called {
		rets, called = methodDirect.Indirect().FuncCall(args)
	}
	return
}

func (tv Tv) FuncCall(in []any) (out []any, called bool) {
	if tv.Type.Kind() != reflect.Func {
		return
	}
	numIn := tv.NumIn()
	if numIn != len(in) {
		return
	}
	rtin := rts(in)
	for i := range numIn {
		if rtin[i] != tv.In(i) {
			return
		}
	}
	return Anys(tv.Call(rvs(in))), true
}
