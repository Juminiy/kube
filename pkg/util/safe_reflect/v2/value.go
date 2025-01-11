package safe_reflectv2

import (
	"github.com/Juminiy/kube/pkg/util/safe_cast/safe_parse"
	"github.com/spf13/cast"
	"reflect"
)

type Value struct {
	reflect.Value
}

func Wrap(rv reflect.Value) Value {
	/*v := Value{Value: rv}
	if rv.CanInterface() {
		v.i = rv.Interface()
	}*/
	return Value{Value: rv}
}

func Wrap2(rv reflect.Value, i any) Value {
	return Value{Value: rv}
}

func (v Value) SetI(i any) {
	indirv := v.indirect()
	if !indirv.CanSet() {
		return
	}
	iv := Direct(i)
	if iv.Value == _ZeroValue ||
		iv.IsZero() {
		indirv.SetZero()
		return
	}
	if indirv.isEFace() || // var i0 any
		indirv.Type() == iv.Type() { // var i0, i T
		indirv.Set(iv.Value)
		return
	}
}

func (v Value) SetILike(i any) {
	indirv := v.indirect()
	if !indirv.CanSet() {
		return
	}
	directi := Direct(i)
	if directi.Value == _ZeroValue ||
		directi.IsZero() {
		indirv.SetZero()
		return
	}
	if indirv.isEFace() || // var i0 any
		indirv.Type() == directi.Type() { // var i0, i T
		indirv.Set(directi.Value)
		return
	} else if indiri := directi.indirect(); indirv.Type() == indiri.Type() {
		indirv.Set(indiri.Value)
		return
	}
	parsed := safe_parse.Parse(cast.ToString(i))
	if pv, ok := parsed.Get(indirv.Kind()); ok {
		indirv.Set(direct(pv))
		return
	}
}

func (v Value) isNil() bool {
	return v.IsNil()
}

func wrapPtr(v any, c int) any {
	for ; c > 0; c-- {
		v = &v
	}
	return v
}

var _ZeroValue = reflect.Value{}
