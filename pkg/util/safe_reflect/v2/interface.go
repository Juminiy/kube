package safe_reflectv2

import "reflect"

func (v Value) isEFace() bool {
	return v.Kind() == reflect.Interface &&
		v.NumMethod() == 0
}

func (v Value) isIFace() bool {
	return v.Kind() == reflect.Interface &&
		v.NumMethod() > 0
}
