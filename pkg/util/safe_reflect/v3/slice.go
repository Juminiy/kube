package safe_reflectv3

import "reflect"

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
