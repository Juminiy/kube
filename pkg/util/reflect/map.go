package reflect

import (
	"reflect"
)

// v maybe indirect map
func mapKeyExistAssign(v reflect.Value, mapKey, mapElem any) {
	v = deref2NoPointer(v)

	if mapCanOpt2(v, mapKey, mapElem) {
		mapKeyV, mapElemV := reflect.ValueOf(mapKey), reflect.ValueOf(mapElem)
		if mapKeyExist(v, mapKeyV) {
			v.SetMapIndex(mapKeyV, mapElemV)
		}
	}
}

// v maybe indirect map
func mapDryAssign(v reflect.Value, mapKey, mapElem any) {
	v = deref2NoPointer(v)

	if mapCanOpt2(v, mapKey, mapElem) {
		v.SetMapIndex(reflect.ValueOf(mapKey), reflect.ValueOf(mapElem))
	}
}

// v maybe indirect map
func mapDryDelete(v reflect.Value, mapKey any) {
	v = deref2NoPointer(v)

	if mapCanOpt(v, mapKey) {
		v.SetMapIndex(reflect.ValueOf(mapKey), _nilValue)
	}
}

// v must be direct map
func mapCanOpt(v reflect.Value, mapKey any) bool {
	return v.Kind() == reflect.Map &&
		!v.IsNil() &&
		v.Type().Key() == reflect.TypeOf(mapKey)
}

// v must be direct map
func mapCanOpt2(v reflect.Value, mapKey, mapElem any) bool {
	return v.Kind() == reflect.Map &&
		!v.IsNil() &&
		v.Type().Key() == reflect.TypeOf(mapKey) &&
		v.Type().Elem() == reflect.TypeOf(mapElem)
}

// v must be direct map && must not nil
func mapKeyExist(v, keyV reflect.Value) bool {
	return v.MapIndex(keyV) != _nilValue
}
