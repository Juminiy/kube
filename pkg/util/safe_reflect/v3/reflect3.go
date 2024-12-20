package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
)

type Tv struct {
	V reflect.Value
	T reflect.Type
}

func Indirect(i any) Tv {
	v := reflect.ValueOf(i)
	for canElem(v) {
		vElem := v.Elem()
		if !canElem(vElem) || (canElem(vElem) && vElem.Elem() == v) {
			v = vElem
			break
		}
		v = vElem
	}
	if !v.IsValid() {
		return Tv{}
	}
	return Tv{
		V: v,
		T: v.Type(),
	}
}

func canElem(v reflect.Value) bool {
	return util.ElemIn(v.Kind(), reflect.Interface, reflect.Pointer)
}
