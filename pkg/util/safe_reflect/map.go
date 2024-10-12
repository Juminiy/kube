package safe_reflect

import (
	"reflect"
)

// Map API
// +param key type and elem type must direct, because of key and elem alignment
// +desc reflect.Map is pointer

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
	return tv.mapCanOpt(key) && tv.mapKeyExist(directV(key))
}

// v is map
func (tv TypVal) mapKeyExistAssign(mapKey, mapElem any) {
	v := tv.noPointer()

	if tv.mapCanOpt2(mapKey, mapElem) {
		mapKeyV, mapElemV := directV(mapKey), directV(mapElem)
		if tv.mapKeyExist(mapKeyV) {
			v.SetMapIndex(mapKeyV, mapElemV)
		}
	}
}

// v is map
func (tv TypVal) mapDryAssign(mapKey, mapElem any) {
	v := tv.noPointer()

	if tv.mapCanOpt2(mapKey, mapElem) {
		v.SetMapIndex(directV(mapKey), directV(mapElem))
	}
}

// v is map
func (tv TypVal) mapDryDelete(mapKey any) {
	v := tv.noPointer()

	if tv.mapCanOpt(mapKey) {
		v.SetMapIndex(directV(mapKey), _nilValue)
	}
}

// v is map
func (tv TypVal) mapCanOpt(mapKey any) bool {
	return tv.Typ.Kind() == reflect.Map &&
		!tv.Val.IsNil() &&
		tv.Typ.Key() == directT(mapKey)
}

// v is map
func (tv TypVal) mapCanOpt2(mapKey, mapElem any) bool {
	return tv.Typ.Kind() == reflect.Map &&
		!tv.Val.IsNil() &&
		tv.Typ.Key() == directT(mapKey) &&
		tv.Typ.Elem() == directT(mapElem)
}

// v must be a map
// call this before must call mapCanOpt or mapCanOpt2
func (tv TypVal) mapKeyExist(keyV reflect.Value) bool {
	return tv.Val.MapIndex(keyV) != _nilValue
}
