package zero_reflect

import "reflect"

func TypeOf(v any) reflect.Type {
	return reflect.TypeOf(v)
}
