package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util/zero_reflect"
	"reflect"
)

func directTV(v any) (typ reflect.Type, val reflect.Value) {
	return directT(v), directV(v)
}

func directT(v any) (typ reflect.Type) { return zero_reflect.TypeOf(v) }

func directV(v any) (val reflect.Value) { return reflect.ValueOf(v) }

func direct(v reflect.Value) TypVal {
	t := v.Type()
	return TypVal{
		Typ: t,
		Val: v,
		typ: t,
		//val: v,
	}
}

func directTs(v []any) []reflect.Type {
	ts := make([]reflect.Type, len(v))
	for i := range v {
		ts[i] = directT(v[i])
	}
	return ts
}

func directVs(v []any) []reflect.Value {
	vs := make([]reflect.Value, len(v))
	for i := range v {
		vs[i] = directV(v[i])
	}
	return vs
}
