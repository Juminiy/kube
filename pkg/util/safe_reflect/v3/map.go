package safe_reflectv3

import (
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"reflect"
	"slices"
)

const MapElemType = "~map_elem_r_type~"

func (t T) MapKeyType() reflect.Type {
	if t.Kind() != reflect.Map {
		return nil
	}
	return t.Key()
}

func (t T) MapElemType() reflect.Type {
	if t.Kind() != reflect.Map {
		return nil
	}
	return t.Elem()
}

func (t T) MapElemNew() Tv {
	if t.Kind() != reflect.Map {
		return Tv{}
	}
	return WrapT(t.MapElemType()).New()
}

func (v V) MapValues() map[string]any {
	return lo.SliceToMap(v.MapRange(), func(item MapKeyElem) (string, any) {
		rk, rv := Any(item.Key), Any(item.Elem)
		if rk != nil && rv != nil {
			return cast.ToString(rk), rv
		}
		return "", nil
	})
}

func (v V) MapDeleteZero() {
	for _, ke := range lo.Filter(v.MapRange(), func(item MapKeyElem, index int) bool {
		return item.Elem.IsZero()
	}) {
		v.SetMapIndex(ke.Key, _ZeroValue)
	}
}

func (v V) MapSetField(nv map[string]any) {
	keyType, elemType := v.Type().Key(), v.Type().Elem()
	if v.Kind() != reflect.Map || keyType.Kind() != reflect.String {
		return
	}
	slices.All(Direct(nv).MapRange())(func(_ int, item MapKeyElem) bool {
		if item.Elem == _ZeroValue || item.Elem.IsNil() {
			v.SetMapIndex(item.Key, _ZeroValue)
		} else if elemType.Kind() == reflect.Interface || elemType == item.Elem.Type() {
			v.SetMapIndex(item.Key, item.Elem)
		} else if elemIndir := WrapV(item.Elem).Indirect(); elemType == elemIndir.Type() {
			v.SetMapIndex(item.Key, elemIndir.Value)
		}
		return true
	})
}

type MapKeyElem struct {
	Key  reflect.Value
	Elem reflect.Value
}

func (v V) MapRange() []MapKeyElem {
	if v.Kind() != reflect.Map {
		return nil
	}
	keyValues := make([]MapKeyElem, v.Len())
	for i, miter := 0, v.Value.MapRange(); miter.Next(); i++ {
		keyValues[i] = MapKeyElem{Key: miter.Key(), Elem: miter.Value()}
	}
	return keyValues
}
