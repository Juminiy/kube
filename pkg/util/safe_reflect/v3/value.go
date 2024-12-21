package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"reflect"
)

type V struct {
	reflect.Value
}

func NewV(i any) V {
	return WrapV(reflect.ValueOf(i))
}

func WrapV(rv reflect.Value) V {
	return V{Value: rv}
}

func (v V) CanElem() bool {
	return util.ElemIn(v.Kind(),
		reflect.Interface, reflect.Pointer)
}

func (v V) Indirect() V {
	vCopy := v
	for vCopy.CanElem() {
		vElem := WrapV(vCopy.Elem())
		if !vElem.CanElem() ||
			(vElem.Elem() == vCopy.Value) {
			vCopy = vElem
			break
		}
		vCopy = vElem
	}
	return vCopy
}

var _ZeroValue = reflect.Value{}

func ZeroValue() reflect.Value {
	return _ZeroValue
}

func rvs(i []any) []reflect.Value {
	return lo.Map(i, func(item any, index int) reflect.Value {
		return reflect.ValueOf(item)
	})
}

func Any(rv reflect.Value) any {
	if rv.IsValid() && rv.CanInterface() {
		return rv.Interface()
	}
	return nil
}

func Anys(rv []reflect.Value) []any {
	return lo.Map(rv, func(item reflect.Value, index int) any {
		return Any(item)
	})
}
