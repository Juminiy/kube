package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	safe_reflectv2 "github.com/Juminiy/kube/pkg/util/safe_reflect/v2"
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

func NewVI(i any) V {
	return WrapVI(reflect.ValueOf(i))
}

func WrapVI(rv reflect.Value) V {
	return V{Value: rv}.Indirect()
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

	case reflect.Map:
		values = append(values, tv.MapValues())

	case reflect.Array:
		values = append(values, tv.ArrayStructValues()...)
		values = append(values, tv.ArrayMapValues()...)

	case reflect.Slice:
		values = append(values, tv.SliceStructValues()...)
		values = append(values, tv.SliceMapValues()...)

	default: // ignore
	}
	return values
}

func (tv Tv) CallMethod(name string, args []any) (rets []any, called bool) {
	callMethod := func(vv V) {
		rets, called = vv.CallMethod(name, args)
		if called {
			return
		}
		switch vv.Kind() {
		case reflect.Array:
			rets, called = vv.ArrayCallMethod(name, args)

		case reflect.Slice:
			rets, called = vv.SliceCallMethod(name, args)

		default: // ignore
		}
	}

	callMethod(tv.V)
	if !called {
		callMethod(tv.V.Indirect())
	}

	return
}

func (v V) SetField(nv map[string]any) {
	switch v.Kind() {
	case reflect.Struct:
		v.StructSet(nv)

	case reflect.Array:
		v.ArrayStructSetField2(nv)

	case reflect.Slice:
		v.SliceStructSetField2(nv)

	case reflect.Map:
		v.MapSetField(nv)

	default: // ignore
	}
}

func (v V) SetI(i any) {
	safe_reflectv2.Wrap(v.Value).SetILike(i)
}

func (v V) I() any {
	if v.IsValid() && v.CanInterface() {
		return v.Interface()
	}
	return nil
}
