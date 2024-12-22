package safe_reflectv3

import (
	"reflect"
)

func (t T) SliceStructTags(tagKey string) Tags {
	if t.Kind() != reflect.Slice {
		return nil
	}
	return t.Indirect().StructTags(tagKey)
}

func (t T) SliceStructTypes() Types {
	if t.Kind() != reflect.Slice {
		return nil
	}
	return t.Indirect().StructTypes()
}

func (t T) SliceStructFields() Fields {
	if t.Kind() != reflect.Slice {
		return nil
	}
	return t.Indirect().StructFields()
}

func (v V) SliceCallMethod(name string, args []any) (rets []any, called bool) {
	if v.Kind() != reflect.Slice ||
		v.Len() == 0 {
		return nil, false
	}
	return WrapV(v.Index(0)).Indirect().CallMethod(name, args)
}

func (v V) SliceStructSetField(index int, nv map[string]any) {
	if v.Kind() != reflect.Slice || !v.CanSet() || v.Len() <= index {
		return
	}
	WrapV(v.Index(index)).StructSet(nv)
}

func (tv Tv) SliceStructValues() []map[string]any {
	t, v := tv.T, tv.V
	if t.Kind() != reflect.Slice {
		return nil
	}
	values := make([]map[string]any, v.Len())
	for i := range v.Len() {
		values[i] = Wrap(v.Index(i)).StructValues()
	}
	return values
}
