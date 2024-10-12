package safe_reflect

import (
	"reflect"
)

// Slice API
// +param elem type can indirect
// +desc reflect.Slice is pointer

// SliceSet
// set slice index to elem -> slice[index] = elem
func (tv TypVal) SliceSet(index int, elem any) {
	v := tv.noPointer()

	if tv.sliceCanOpt(elem) && tv.FieldLen() > index {
		elemV := noPointer(v.Index(index))
		if elemV.CanSet() {
			elemV.Set(indirectV(elem))
		}
	}
}

// SliceSetStructFields
// set slice struct fields fieldName to fieldVal
func (tv TypVal) SliceSetStructFields(fields map[string]any) {
	v := tv.noPointer()
	if v.Kind() != reflect.Slice {
		return
	}

	for index := range tv.FieldLen() {
		indirect(v.Index(index)).StructSetFields(fields)
	}
}

func (tv TypVal) sliceCanOpt(elem any) bool {
	return tv.Typ.Kind() == reflect.Slice &&
		!tv.Val.IsNil() &&
		underlyingEqual(tv.Typ.Elem(), reflect.TypeOf(elem))
}

// SliceSetOol
// set slice index to elem -> slice[index] = elem that allow length out of bound, but capacity inbound
func (tv TypVal) SliceSetOol(index int, elem any) {
	tv.SliceShiftLenInc(index + 1)
	tv.SliceSet(index, elem)
}

// SliceSetOoc
// set slice index to elem -> slice[index] = elem that allow length and capacity out of bound
func (tv TypVal) SliceSetOoc(index int, elem any) {
	tv.SliceGrow(index + 1)
	tv.SliceSetOol(index, elem)
}

func (tv TypVal) SliceSetLen(toLen int) {
	v := tv.noPointer()
	if v.Kind() == reflect.Slice && v.CanSet() && toLen <= v.Cap() {
		v.SetLen(toLen)
	}
}

func (tv TypVal) SliceSetCap(toCap int) {
	v := tv.noPointer()
	if v.Kind() == reflect.Slice && v.CanSet() && v.Len() <= toCap && toCap <= v.Cap() {
		v.SetCap(toCap)
	}
}

func (tv TypVal) SliceShiftLenInc(toLen int) {
	v := tv.noPointer()
	if v.Kind() == reflect.Slice && v.CanSet() && toLen <= v.Cap() && v.Len() < toLen {
		v.SetLen(toLen)
	}
}

func (tv TypVal) SliceShiftLenDec(toLen int) {
	v := tv.noPointer()
	if v.Kind() == reflect.Slice && v.CanSet() && toLen <= v.Cap() && v.Len() > toLen {
		v.SetLen(toLen)
	}
}

func (tv TypVal) SliceShiftLen2Cap() {
	v := tv.noPointer()
	if v.Kind() == reflect.Slice && v.CanSet() {
		v.SetLen(v.Cap())
	}
}

func (tv TypVal) SliceGrow(toCap int) {
	v := tv.noPointer()
	if v.Kind() == reflect.Slice && v.CanSet() {
		v.Grow(toCap)
	}
}
