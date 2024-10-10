package reflect

import (
	"reflect"
)

// MapAssign key exist assign
func (tv TypVal) MapAssign(key, elem any) {
	tv.mapKeyExistAssign(key, elem)
}

// MapAssign2 dry assign
func (tv TypVal) MapAssign2(key, elem any) {
	tv.mapDryAssign(key, elem)
}

// MapDelete dry delete
func (tv TypVal) MapDelete(key any) {
	tv.mapDryDelete(key)
}

// MapKeyOk key exist
func (tv TypVal) MapKeyOk(key any) bool {
	return tv.mapCanOpt(key) && tv.mapKeyExist(reflect.ValueOf(key))
}

// v is map
func (tv TypVal) mapKeyExistAssign(mapKey, mapElem any) {
	v := tv.noPointer()

	if tv.mapCanOpt2(mapKey, mapElem) {
		mapKeyV, mapElemV := reflect.ValueOf(mapKey), reflect.ValueOf(mapElem)
		if tv.mapKeyExist(mapKeyV) {
			v.SetMapIndex(mapKeyV, mapElemV)
		}
	}
}

// v is map
func (tv TypVal) mapDryAssign(mapKey, mapElem any) {
	v := tv.noPointer()

	if tv.mapCanOpt2(mapKey, mapElem) {
		v.SetMapIndex(reflect.ValueOf(mapKey), reflect.ValueOf(mapElem))
	}
}

// v is map
func (tv TypVal) mapDryDelete(mapKey any) {
	v := tv.noPointer()

	if tv.mapCanOpt(mapKey) {
		v.SetMapIndex(reflect.ValueOf(mapKey), _nilValue)
	}
}

// v is map
func (tv TypVal) mapCanOpt(mapKey any) bool {
	return tv.Typ.Kind() == reflect.Map &&
		!tv.Val.IsNil() &&
		tv.Typ.Key() == reflect.TypeOf(mapKey)
}

// v is map
func (tv TypVal) mapCanOpt2(mapKey, mapElem any) bool {
	return tv.Typ.Kind() == reflect.Map &&
		!tv.Val.IsNil() &&
		tv.Typ.Key() == reflect.TypeOf(mapKey) &&
		tv.Typ.Elem() == reflect.TypeOf(mapElem)
}

// v must be a map
// call this before must call mapCanOpt or mapCanOpt2
func (tv TypVal) mapKeyExist(keyV reflect.Value) bool {
	return tv.Val.MapIndex(keyV) != _nilValue
}
