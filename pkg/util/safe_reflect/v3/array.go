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
