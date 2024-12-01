package safe_reflectv2

import (
	"reflect"
)

func direct(i any) (rv reflect.Value) {
	return reflect.ValueOf(i)
}
