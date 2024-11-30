package safe_reflect

import "reflect"

func IndirectRVal(v reflect.Value) reflect.Value {
	return indirect(v).Val
}

func indirectTV(v any) (typ reflect.Type, val reflect.Value) {
	return indirectT(v), indirectV(v)
}

func indirectT(v any) (typ reflect.Type) {
	return underlying(directT(v))
}

func indirectV(v any) (val reflect.Value) {
	return noPointer(directV(v))
}

func indirect(v reflect.Value) TypVal {
	indirV := noPointer(v)
	return TypVal{
		Typ: rValueType(indirV),
		Val: indirV,
		typ: rValueType(v),
		//val: v,
	}
}

func indirectTs(v []any) []reflect.Type {
	ts := make([]reflect.Type, len(v))
	for i := range v {
		ts[i] = indirectT(v[i])
	}
	return ts
}

func indirectVs(v []any) []reflect.Value {
	vs := make([]reflect.Value, len(v))
	for i := range v {
		vs[i] = indirectV(v[i])
	}
	return vs
}
