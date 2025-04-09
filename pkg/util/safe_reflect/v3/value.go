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
	return lo.Map(i, func(item any, _ int) reflect.Value {
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
	return lo.Map(rv, func(item reflect.Value, _ int) any {
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

	case reflect.Map:
		v.MapSetField(nv)

	case reflect.Array:
		if elemKind := WrapT(v.Type()).Indirect().Kind(); elemKind == reflect.Struct {
			v.ArrayStructSetField2(nv)
		} else if elemKind == reflect.Map {
			v.ArrayMapSetField2(nv)
		}

	case reflect.Slice:
		if elemKind := WrapT(v.Type()).Indirect().Kind(); elemKind == reflect.Struct {
			v.SliceStructSetField2(nv)
		} else if elemKind == reflect.Map {
			v.SliceMapSetField2(nv)
		}

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

// ToAnySlice
// []Tpl -> []any
// [N]Tpl -> []any
// Tpl -> []any
func ToAnySlice(i any) []any {
	tv := Direct(i)
	if tv.Value == _ZeroValue {
		return []any{nil}
	}
	aS := make([]any, 0, util.MagicSliceCap)
	switch tv.Type.Kind() {
	case reflect.Array, reflect.Slice:
		tv.IterIndex()(func(_ int, value reflect.Value) bool {
			if value.IsValid() && value.CanInterface() {
				aS = append(aS, value.Interface())
			}
			return true
		})
	default:
		return append(aS, i)
	}
	return aS
}

// CopyFieldValue
// reinforce of safe_reflect.CopyFieldValue
func CopyFieldValue(src any, dst any) {
	srcOf, dstOf := Indirect(src), Indirect(dst)

	switch {
	case srcOf.V.Value != _ZeroValue &&
		dstOf.V.Value != _ZeroValue &&
		srcOf.T.Kind() == reflect.Struct &&
		dstOf.T.Kind() == reflect.Struct &&
		dstOf.V.CanSet():
		if srcOf.Type == dstOf.Type {
			dstOf.Set(srcOf.Value)
			return
		}
		srcFieldValues := srcOf.StructToMap()
		for idx := range dstOf.T.NumField() {
			if srcFieldValue, ok := srcFieldValues[dstOf.T.Field(idx).Name]; ok {
				WrapV(dstOf.V.Field(idx)).SetI(srcFieldValue)
			}
		}
	}

}
