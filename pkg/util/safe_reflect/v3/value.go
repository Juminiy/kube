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

func (tv Tv) Values() []map[string]any {
	t, _ := tv.T, tv.V
	var values []map[string]any
	switch t.Kind() {
	case reflect.Struct:
		values = append(values, tv.StructValues())

	case reflect.Array:
		values = append(values, tv.ArrayStructValues()...)

	case reflect.Slice:
		values = append(values, tv.SliceStructValues()...)

	case reflect.Map:
		values = append(values, tv.MapValues())

	default: // ignore
	}
	return values
}

func (tv Tv) CallMethod(name string, args []any) (rets []any, called bool) {
	t, v := tv.T, tv.V
	switch t.Kind() {
	case reflect.Array:
		rets, called = v.ArrayCallMethod(name, args)

	case reflect.Slice:
		rets, called = v.SliceCallMethod(name, args)

	default:
		rets, called = v.CallMethod(name, args)

	}
	return
}

func (v V) SetField(nv map[string]any, index ...int) {
	switch v.Kind() {
	case reflect.Struct:
		v.StructSet(nv)

	case reflect.Array:
		if len(index) > 0 {
			v.ArrayStructSetField(index[0], nv)
		}

	case reflect.Slice:
		if len(index) > 0 {
			v.SliceStructSetField(index[0], nv)
		}

	case reflect.Map:
		v.MapSetField(nv)

	default: // ignore
	}
}

func (v V) SetI(i any) {
	V2Wrap(v.Value).SetILike(i)
}
