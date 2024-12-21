package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"reflect"
)

type T struct {
	reflect.Type
}

func NewT(i any) T {
	return WrapT(reflect.TypeOf(i))
}

func WrapT(rt reflect.Type) T {
	return T{Type: rt}
}

func (t T) CanElem() bool {
	return util.ElemIn(t.Kind(),
		reflect.Array, reflect.Chan, reflect.Map, reflect.Pointer, reflect.Slice)
}

func (t T) Indirect() T {
	tCopy := t
	for util.ElemIn(tCopy.Kind(),
		reflect.Array, reflect.Slice, reflect.Chan, reflect.Pointer,
	) {
		tCopy = WrapT(tCopy.Elem())
	}
	return tCopy
}

func rts(i []any) []reflect.Type {
	return lo.Map(i, func(item any, index int) reflect.Type {
		return reflect.TypeOf(item)
	})
}
