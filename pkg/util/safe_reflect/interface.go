package safe_reflect

import (
	"reflect"
)

// interface is equal to any to say: recommend any

// WARNING: maybe cause dead-loop, be careful when use them

const (
	_unpackLoopMax = 8 // safe_guard
)

func InterfaceOf(v reflect.Value) any {
	if v.CanInterface() {
		return v.Interface()
	}
	return nil
}

func InterfacesOf(v []reflect.Value) []any {
	as := make([]any, len(v))
	for i := range v {
		as[i] = InterfaceOf(v[i])
	}
	return as
}

// NewOf
// new value for type:typ no pointer
func NewOf(typ reflect.Type) any {
	return reflect.New(typ).Elem().Interface()
}

// NewOf2
// new value for type:typ pointer
// sometimes useful
func NewOf2(typ reflect.Type) any {
	return reflect.New(typ).Interface()
}

// unpackOf
// do not export
func unpackOf(v any) TypVal {
	tv := Of(v)
	tv.unpack()
	return tv
}

func (tv *TypVal) unpack() reflect.Value {
	v, packed := unpackOk(tv.Val)
	if packed {
		tv.Val = v
		tv.Typ = rValueType(v)
	}
	return v
}

func unpack(v reflect.Value) reflect.Value {
	for i := 0; valueCanElem(v) && i < _unpackLoopMax; i++ {
		v = v.Elem()
	}
	return v
}

func unpackOk(v reflect.Value) (reflect.Value, bool) {
	packed := false
	for i := 0; valueCanElem(v) && i < _unpackLoopMax; i++ {
		v, packed = v.Elem(), true
	}
	return v, packed
}

func unpackV(v any) reflect.Value {
	return unpack(directV(v))
}

func unpackT(v any) reflect.Type {
	return rValueType(unpackV(v))
}

func unpackEqual(v0, v1 any) bool {
	return unpackV(v0) == unpackV(v1)
}

// Impl
// iFace must an iface or called not an empty interface
//
//	type _typeIFace interface {
//		Method0(in...)out...
//		Method1(in...)out...
//		...
//	}
//
// +param iFace
// var iFace = (*_typeIFace)(nil)
func Impl(v any, iFace any) bool {
	ifaceTyp, ok := ifaceType(iFace)
	if !ok {
		return false
	}
	return impl(directT(v), ifaceTyp)
}

func ifaceType(iFacePtr any) (reflect.Type, bool) {
	ifacePtrType := directT(iFacePtr)
	if !typeCanElem(ifacePtrType) || ifacePtrType.Elem().Kind() != IFace {
		return nil, false
	}
	return ifacePtrType.Elem(), true
}

func impl(v, iface reflect.Type) bool {
	return v.Implements(iface)
}

// IndirectImpl
// +param iFace same with Impl
// +param v
// specialized for
// T, *T, **T, *...*T -> T
// Arr, *Arr, *...*Arr -> Arr.Elem
// Slice, *Slice, *...*Slice -> Slice.Elem
func IndirectImpl(v any, iFace any) (inst any, inst2 any) {
	ifaceTyp, ok := ifaceType(iFace)
	if !ok {
		return
	}

	// pointer impl
	val := onePointer(directV(v))
	typ := rValueType(val)
	if impl(typ, ifaceTyp) &&
		val.CanInterface() {
		inst = val.Interface()
		inst2 = inst
		return
	}

	// value impl
	typ, val = indirectTV(v)
	if impl(typ, ifaceTyp) &&
		val.CanInterface() {
		inst = val.Interface()
		inst2 = inst
		return
	}

	switch typ.Kind() {
	case Arr, Slice:
		elemTyp := underlying(typ.Elem())
		if impl(elemTyp, ifaceTyp) {
			return NewOf(elemTyp), NewOf2(elemTyp)
		}

	default:
		return
	}

	return
}
