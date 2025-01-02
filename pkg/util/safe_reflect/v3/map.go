package safe_reflectv3

import (
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"reflect"
	"slices"
)

const MapElemType = "~map_elem_r_type~"

func (t T) MapElemType() reflect.Type {
	if t.Kind() != reflect.Map {
		return nil
	}
	return t.Type.Elem()
}

func (t T) MapElemNew() Tv {
	if t.Kind() != reflect.Map {
		return Tv{}
	}
	return WrapT(t.MapElemType()).New()
}

func (v V) MapValues() map[string]any {
	return lo.SliceToMap(v.MapRange(), func(item MapKeyValue) (string, any) {
		rk, rv := Any(item.Key), Any(item.Value)
		if rk != nil && rv != nil {
			return cast.ToString(rk), rv
		}
		return "", nil
	})
}

func (v V) MapDeleteZero() {
	for _, kv := range lo.Filter(v.MapRange(), func(item MapKeyValue, index int) bool {
		return item.Value.IsZero()
	}) {
		v.SetMapIndex(kv.Key, _ZeroValue)
	}
}

func (v V) MapSetField(nv map[string]any) {
	if v.Kind() != reflect.Map || v.Type().Kind() != reflect.String {
		return
	}
	slices.All(Indirect(nv).MapRange())(func(_ int, kv MapKeyValue) bool {
		if kv.Value != _ZeroValue && kv.Value.Type() == v.Type().Elem() {
			v.SetMapIndex(kv.Key, kv.Value)
		}
		return true
	})
}

type MapKeyValue struct {
	Key   reflect.Value
	Value reflect.Value
}

func (v V) MapRange() []MapKeyValue {
	if v.Kind() != reflect.Map {
		return nil
	}
	keyValues := make([]MapKeyValue, v.Len())
	for i, miter := 0, v.Value.MapRange(); miter.Next(); i++ {
		keyValues[i] = MapKeyValue{Key: miter.Key(), Value: miter.Value()}
	}
	return keyValues
}
