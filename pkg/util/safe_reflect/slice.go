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

// SliceAppend
// stand for: sl = append(sl, elem)
func (tv TypVal) SliceAppend(elem any) {
	v := tv.noPointer()
	if v.Kind() != Slice || !v.CanSet() ||
		directT(elem) != tv.Typ.Elem() {
		return
	}
	v.Set(reflect.Append(v, directV(elem)))
}

// SliceAppends
// stand for: sl = append(sl, elem...)
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

// SliceAppendSlice
// stand for: sl = append(sl, slAdd...)
func (tv TypVal) SliceAppendSlice(slAdd any) {
	v := tv.noPointer()
	slOf := Of(slAdd)
	if v.Kind() != Slice || !v.CanSet() ||
		slOf.Typ != tv.Typ || slOf.Typ.Elem() != tv.Typ.Elem() {
		return
	}
	v.Set(reflect.AppendSlice(v, slOf.Val))
}

// SliceStructFieldValues
// []Struct get field(fieldName) values
func (tv TypVal) SliceStructFieldValues(fieldName string) map[any]struct{} {
	v := tv.noPointer()

	if v.Kind() != Slice ||
		tv.FieldLen() == 0 {
		return nil
	}
	directElem0 := direct(v.Index(0))

	fieldIndex := directElem0.structFieldIndexByName(fieldName)
	if len(fieldIndex) == 0 {
		return nil
	}
	fieldValues := make(map[any]struct{}, tv.FieldLen())
	for index := range tv.FieldLen() {
		if elemI := v.Index(index); elemI.Kind() == Struct {
			if elemIFi := elemI.FieldByIndex(fieldIndex); elemIFi.CanInterface() {
				fieldValues[elemIFi.Interface()] = struct{}{}
			}
		}
	}
	return fieldValues
}

// SliceStructFieldsValues
// []Struct get fields(fieldNames) values
func (tv TypVal) SliceStructFieldsValues(fields map[string]struct{}) map[string]map[any]struct{} {
	v := tv.noPointer()

	if v.Kind() != Slice ||
		tv.FieldLen() == 0 {
		return nil
	}
	directElem0 := direct(v.Index(0))

	// all field list
	fieldsIndex := directElem0.StructFieldsIndex()

	// common field list
	util.MapEvict(fieldsIndex, fields)

	fieldsValues := directElem0.StructFieldsValues(fieldsIndex)
	for index := range tv.FieldLen() {
		util.MapMerge(fieldsValues, direct(v.Index(index)).StructFieldsValues(fieldsIndex))
	}
	return fieldsValues
}

// SliceStruct2SliceMap
// Struct to map[fieldName]fieldValue
func (tv TypVal) SliceStruct2SliceMap(fields map[string]struct{}) []map[string]any {
	v := tv.noPointer()

	if v.Kind() != Slice ||
		tv.FieldLen() == 0 {
		return nil
	}

	recordValues := make([]map[string]any, tv.FieldLen())
	for index := range tv.FieldLen() {
		recordValues[index] = direct(v.Index(index)).Struct2Map(fields)
	}
	return recordValues
}

// StructMakeSlice
// Type Struct -> return []Struct
func (tv TypVal) StructMakeSlice(length, capacity int) any {
	v := tv.noPointer()
	if v.Kind() != Struct {
		return nil
	}
	avaLenCap(&length, &capacity)
	return reflect.MakeSlice(reflect.SliceOf(tv.Typ), length, capacity).Interface()
}

// SliceOrArrayStructHasFields
// like StructHasFields only check type not value
func (tv TypVal) SliceOrArrayStructHasFields(fields map[string]any) map[string]struct{} {
	v := tv.noPointer()
	if v.Kind() != Slice && v.Kind() != Arr {
		return nil
	}

	underElemTyp := underlying(tv.Typ.Elem())
	if underElemTyp.Kind() != Struct {
		return nil
	}
	return structHasFields(underElemTyp, fields)
}

// SliceOrArrayStructGetTagVal
// get slice-struct or array-struct fields app key tag value
func (tv TypVal) SliceOrArrayStructGetTagVal(app, key string) []string {
	v := tv.noPointer()
	if v.Kind() != Slice && v.Kind() != Arr {
		return nil
	}

	underElemTyp := underlying(tv.Typ.Elem())
	if underElemTyp.Kind() != Struct {
		return nil
	}
	return structParseTag(underElemTyp, app).GetValList(key)
}

func SliceMake(elem any, length, capacity int) any {
	if elem == nil {
		return nil
	}

	avaLenCap(&length, &capacity)

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

func avaLenCap(length, capacity *int) {
	if *capacity <= 0 {
		*capacity = util.MagicSliceCap
	}
	if *length < 0 {
		*length = 0
	}
	if *capacity < *length {
		*capacity = *length
	}
}
