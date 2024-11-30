package safe_validator

import (
	"github.com/Juminiy/kube/pkg/util/safe_cast/safe_parse"
	"reflect"
)

func (f fieldOf) setDefault(tagv string) {
	if !f.rval.CanSet() || !f.rval.IsZero() {
		return
	}

	switch f.rkind {
	case kBool:
		v, ok := safe_parse.ParseBool(tagv)
		if ok {
			f.rval.SetBool(v)
		}
	case kInt:
		v, ok := safe_parse.ParseInt(tagv)
		if ok {
			f.rval.SetInt(int64(v))
		}
	case kI8:
		v, ok := safe_parse.ParseInt8(tagv)
		if ok {
			f.rval.SetInt(int64(v))
		}
	case kI16:
		v, ok := safe_parse.ParseInt16(tagv)
		if ok {
			f.rval.SetInt(int64(v))
		}
	case kI32:
		v, ok := safe_parse.ParseInt32(tagv)
		if ok {
			f.rval.SetInt(int64(v))
		}
	case kI64:
		v, ok := safe_parse.ParseInt64(tagv)
		if ok {
			f.rval.SetInt(v)
		}
	case kUint:
		v, ok := safe_parse.ParseUint(tagv)
		if ok {
			f.rval.SetUint(uint64(v))
		}
	case kU8:
		v, ok := safe_parse.ParseUint8(tagv)
		if ok {
			f.rval.SetUint(uint64(v))
		}
	case kU16:
		v, ok := safe_parse.ParseUint16(tagv)
		if ok {
			f.rval.SetUint(uint64(v))
		}
	case kU32:
		v, ok := safe_parse.ParseUint32(tagv)
		if ok {
			f.rval.SetUint(uint64(v))
		}
	case kU64:
		v, ok := safe_parse.ParseUint64(tagv)
		if ok {
			f.rval.SetUint(v)
		}
	case kUPtr:
		v, ok := safe_parse.ParseUintptr(tagv)
		if ok {
			f.rval.SetUint(uint64(v))
		}
	case kF32:
		v, ok := safe_parse.ParseFloat32(tagv)
		if ok {
			f.rval.SetFloat(float64(v))
		}
	case kF64:
		v, ok := safe_parse.ParseFloat64(tagv)
		if ok {
			f.rval.SetFloat(v)
		}
	case kString:
		f.rval.SetString(tagv)
	case kPtr:
		f.setDefaultToPtr(tagv)
	default:
		panic(errTagKindCheckErr)
	}
}

func (f fieldOf) setDefaultToPtr(tagv string) {
	parsedTypedV := safe_parse.Parse(tagv)
	rvalElem := f.rval.Elem()
	if rvalElem.CanSet() {
		val, ok := parsedTypedV.Get(rvalElem.Kind())
		if ok {
			rvalElem.Set(reflect.ValueOf(val))
		}
	}

}
