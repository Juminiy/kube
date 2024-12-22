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

func (v V) ArrayCallMethod(name string, args []any) (rets []any, called bool) {
	if v.Kind() != reflect.Array ||
		v.Len() == 0 {
		return nil, false
	}
	return WrapV(v.Index(0)).Indirect().CallMethod(name, args)
}

func (v V) ArrayStructSetField(index int, nv map[string]any) {
	if v.Kind() != reflect.Array || !v.CanSet() || v.Len() <= index {
		return
	}
	WrapV(v.Index(index)).StructSet(nv)
}

func (tv Tv) ArrayStructValues() []map[string]any {
	t, v := tv.T, tv.V
	if t.Kind() != reflect.Array {
		return nil
	}
	values := make([]map[string]any, v.Len())
	for i := range v.Len() {
		values[i] = Wrap(v.Index(i)).StructValues()
	}
	return values
}
