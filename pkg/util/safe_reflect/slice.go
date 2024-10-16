package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
)

// Slice API
// +param elem type can indirect
// +desc Slice is pointer

// SliceSet
// set slice index to elem -> slice[index] = elem
func (tv TypVal) SliceSet(index int, elem any) {
	v := tv.noPointer()

	if v.Kind() != Slice ||
		tv.FieldLen() <= index ||
		!tv.sliceCanOpt(elem) {
		return
	}

	if indirIndexV := noPointer(v.Index(index)); indirIndexV.CanSet() {
		indirIndexV.Set(indirectV(elem))
	}
}

// SliceSetStructFields
// set slice struct fields fieldName to fieldVal
func (tv TypVal) SliceSetStructFields(fields map[string]any) {
	v := tv.noPointer()

	if v.Kind() != Slice ||
		tv.FieldLen() == 0 {
		return
	}

	for index := range tv.FieldLen() {
		indirect(v.Index(index)).StructSetFields(fields)
	}
}

func (tv TypVal) sliceCanOpt(elem any) bool {
	return tv.Typ.Kind() == Slice &&
		!tv.Val.IsNil() &&
		tv.sliceElemTypeEq(elem)
}

// SliceSetOol
// set slice index to elem -> slice[index] = elem that allow length out of bound, but capacity inbound
func (tv TypVal) SliceSetOol(index int, elem any) {
	tv.sliceShiftLenInc(index + 1)
	tv.SliceSet(index, elem)
}

// SliceSetOoc
// set slice index to elem -> slice[index] = elem that allow length and capacity out of bound
func (tv TypVal) SliceSetOoc(index int, elem any) {
	tv.sliceGrowTo(index + 1)
	tv.SliceSetOol(index, elem)
}

// SliceSetMake
// set slice index to elem -> slice[index] = elem
// allow slice is nil, if slice is nil make a slice
// allow index out of bound capacity, auto resize length and resize capacity
func (tv TypVal) SliceSetMake(index int, elem any) {
	tv.sliceNilDrySet(index, elem)
}

func (tv TypVal) sliceNilDrySet(index int, elem any) {
	v := tv.noPointer()

	if v.Kind() != Slice || !tv.sliceElemTypeEq(elem) {
		return
	}

	if v.IsNil() && v.CanSet() {
		v.Set(reflect.MakeSlice(tv.Typ, index+1, (index+1)<<1))
	}

	if !v.IsNil() {
		tv.SliceSetOoc(index, elem)
	}
}

func (tv TypVal) sliceElemTypeEq(elem any) bool {
	return underlyingEqual(tv.Typ.Elem(), directT(elem))
}

func (tv TypVal) sliceSetLen(toLen int) {
	v := tv.noPointer()
	if v.Kind() == Slice && v.CanSet() &&
		toLen <= v.Cap() {
		v.SetLen(toLen)
	}
}

func (tv TypVal) sliceSetCap(toCap int) {
	v := tv.noPointer()
	if v.Kind() == Slice && v.CanSet() &&
		v.Len() <= toCap && toCap <= v.Cap() {
		v.SetCap(toCap)
	}
}

func (tv TypVal) sliceShiftLenInc(toLen int) {
	v := tv.noPointer()
	if v.Kind() == Slice && v.CanSet() &&
		toLen <= v.Cap() && v.Len() < toLen {
		v.SetLen(toLen)
	}
}

func (tv TypVal) sliceShiftLenDec(toLen int) {
	v := tv.noPointer()
	if v.Kind() == Slice && v.CanSet() &&
		toLen <= v.Cap() && v.Len() > toLen {
		v.SetLen(toLen)
	}
}

func (tv TypVal) sliceShiftLen2Cap() {
	v := tv.noPointer()
	if v.Kind() == Slice && v.CanSet() {
		v.SetLen(v.Cap())
	}
}

func (tv TypVal) sliceGrowTo(toCap int) {
	v := tv.noPointer()
	if v.Kind() == Slice && v.CanSet() && toCap > v.Cap() {
		v.Grow(toCap - v.Cap())
	}
}

func (tv TypVal) SliceAppend(elem any) {
	v := tv.noPointer()
	if v.Kind() != Slice || !v.CanSet() ||
		directT(elem) != tv.Typ.Elem() {
		return
	}
	v.Set(reflect.Append(v, directV(elem)))
}

func (tv TypVal) SliceAppends(elem ...any) {
	v := tv.noPointer()
	if v.Kind() != Slice || !v.CanSet() {
		return
	}
	for i := range elem {
		if directT(elem[i]) == tv.Typ.Elem() {
			v.Set(reflect.Append(v, directV(elem[i])))
		}
	}
}

func (tv TypVal) SliceAppendSlice(sl any) {
	v := tv.noPointer()
	slOf := Of(sl)
	if v.Kind() != Slice || !v.CanSet() ||
		slOf.Typ != tv.Typ || slOf.Typ.Elem() != tv.Typ.Elem() {
		return
	}
	v.Set(reflect.AppendSlice(v, slOf.Val))
}

func (tv TypVal) SliceStructFieldsValues(fields map[string]struct{}) map[string]map[any]struct{} {
	v := tv.noPointer()

	if v.Kind() != Slice ||
		tv.FieldLen() == 0 {
		return nil
	}

	// all field list
	fieldsIndex := direct(v.Index(0)).StructFieldsIndex()

	// common field list
	util.MapEvict(fieldsIndex, fields)

	fieldsValues := indirect(v.Index(0)).StructFieldsValues(fieldsIndex)
	for index := range tv.FieldLen() {
		util.MapMerge(fieldsValues, indirect(v.Index(index)).StructFieldsValues(fieldsIndex))
	}
	return fieldsValues
}

func SliceMake(elem any, length, capacity int) any {
	if elem == nil {
		return nil
	}
	if capacity <= 0 {
		capacity = util.MagicSliceCap
	}
	if length < 0 {
		length = 0
	}
	if capacity < length {
		capacity = length
	}

	//Pointer -> Slice(Addr)
	// old-version: worked
	//slOf := reflect.New(sliceType(elem)).Elem()
	//dirSlOf := direct(slOf)
	//dirSlOf.sliceGrowTo(capacity)
	//dirSlOf.sliceSetLen(length)

	// new-version: worked, optimized compare to old-version
	//Slice-> Slice
	return reflect.MakeSlice(sliceType(elem), length, capacity).Interface()
}

func sliceType(elem any) reflect.Type {
	return reflect.SliceOf(directT(elem))
}
