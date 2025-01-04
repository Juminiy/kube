package safe_reflectv3

import "reflect"

func (t T) ArrayStructTags(tagKey string) Tags {
	if t.Kind() != reflect.Array {
		return nil
	}
	return t.Indirect().StructTags(tagKey)
}

func (t T) ArrayStructTypes() Types {
	if t.Kind() != reflect.Array {
		return nil
	}
	return t.Indirect().StructTypes()
}

func (t T) ArrayStructFields() Fields {
	if t.Kind() != reflect.Array {
		return nil
	}
	return t.Indirect().StructFields()
}

func (t T) ArrayElemNew() Tv {
	if t.Kind() != reflect.Array {
		return Tv{}
	}
	return WrapT(t.Elem()).New()
}

func (v V) ArrayCallMethod(name string, args []any) (rets []any, called bool) {
	if v.Kind() != reflect.Array ||
		v.Len() == 0 {
		return nil, false
	}
	return WrapV(v.Index(0)).CallMethod(name, args)
}

func (v V) ArrayStructSetField(index int, nv map[string]any) {
	if v.Kind() != reflect.Array || v.Len() <= index {
		return
	}
	WrapVI(v.Index(index)).StructSet(nv)
}

func (v V) ArrayStructSetField2(nv map[string]any) {
	if v.Kind() != reflect.Array {
		return
	}
	for i := range v.Len() {
		WrapVI(v.Index(i)).StructSet(nv)
	}
}

func (v V) ArraySet(index int, i any) {
	if v.Kind() != reflect.Array || !v.CanSet() || v.Len() <= index {
		return
	}
	WrapV(v.Index(index)).SetI(i)
}

func (tv Tv) ArrayStructValues() []map[string]any {
	t, v := tv.T, tv.V
	if t.Kind() != reflect.Array || t.Indirect().Kind() != reflect.Struct {
		return nil
	}
	values := make([]map[string]any, v.Len())
	tv.IterIndex()(func(i int, rv reflect.Value) bool {
		values[i] = WrapI(rv).StructValues()
		return true
	})
	return values
}

func (tv Tv) ArrayMapValues() []map[string]any {
	t, v := tv.T, tv.V
	if t.Kind() != reflect.Array || t.Indirect().Kind() != reflect.Map {
		return nil
	}
	values := make([]map[string]any, v.Len())
	tv.IterIndex()(func(i int, rv reflect.Value) bool {
		values[i] = WrapI(rv).MapValues()
		return true
	})
	return values
}
