package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
)

// Map API
// +param key type must direct, because of comparable
// +param elem type must direct, because of alignment and ptr
// +desc Map is pointer

// MapAssign key exist assign
func (tv TypVal) MapAssign(key, elem any) {
	tv.mapKeyExistAssign(key, elem)
}

func (tv TypVal) mapKeyExistAssign(key, elem any) {
	v := tv.noPointer()

	if tv.mapCanOpt2(key, elem) {
		mapKeyV, mapElemV := directV(key), directV(elem)
		if tv.mapKeyExist(mapKeyV) {
			v.SetMapIndex(mapKeyV, mapElemV)
		}
	}
}

// MapAssign2 dry assign
func (tv TypVal) MapAssign2(key, elem any) {
	tv.mapDryAssign(key, elem)
}

func (tv TypVal) mapDryAssign(key, elem any) {
	v := tv.noPointer()

	if tv.mapCanOpt2(key, elem) {
		v.SetMapIndex(directV(key), directV(elem))
	}
}

// MapDelete dry delete
func (tv TypVal) MapDelete(key any) {
	tv.mapDryDelete(key)
}

func (tv TypVal) mapDryDelete(key any) {
	v := tv.noPointer()

	if tv.mapCanOpt(key) {
		v.SetMapIndex(directV(key), _zeroValue)
	}
}

// MapKeyOk key exist
func (tv TypVal) MapKeyOk(key any) bool {
	return tv.mapCanOpt(key) && tv.mapKeyExist(directV(key))
}

// call before must call mapCanOpt or mapCanOpt2
func (tv TypVal) mapKeyExist(vOfKey reflect.Value) bool {
	return tv.Val.MapIndex(vOfKey) != _zeroValue
}

// MapAssignMake if nil make map
func (tv TypVal) MapAssignMake(key, elem any) {
	tv.mapNilDryAssign(key, elem)
}

func (tv TypVal) mapNilDryAssign(key, elem any) {
	v := tv.noPointer()

	if v.Kind() != Map || !tv.mapKeyElemTypeEq(key, elem) {
		return
	}

	if v.IsNil() && v.CanSet() {
		v.Set(reflect.MakeMapWithSize(tv.Typ, util.MagicMapCap))
	}

	if !v.IsNil() {
		v.SetMapIndex(directV(key), directV(elem))
	}
}

func (tv TypVal) mapCanOpt(key any) bool {
	return tv.Typ.Kind() == Map &&
		!tv.Val.IsNil() &&
		tv.Typ.Key() == directT(key)
}

func (tv TypVal) mapCanOpt2(key, elem any) bool {
	return tv.Typ.Kind() == Map &&
		!tv.Val.IsNil() &&
		tv.mapKeyElemTypeEq(key, elem)
}

// call before must call noPointer and checkKind
func (tv TypVal) mapKeyElemTypeEq(key, elem any) bool {
	mapElemTyp := tv.Typ.Elem()
	return tv.Typ.Key() == directT(key) &&
		(mapElemTyp.Kind() == Any || // map[key_type]any
			mapElemTyp == directT(elem)) // map[key_type]elem_type
}
