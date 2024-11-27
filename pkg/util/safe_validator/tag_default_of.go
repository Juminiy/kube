package safe_validator

import "github.com/Juminiy/kube/pkg/util/safe_cast/safe_parse"

func (f fieldOf) setDefault(tagv string) {
	if !f.rval.CanSet() || !f.rval.IsZero() {
		return
	}
	switch f.rkind {
	case kBool:
		v, ok := safe_parse.ParseBoolOk(tagv)
		if ok {
			f.rval.SetBool(v)
		}
	case kInt:
		v, ok := safe_parse.ParseIntOk(tagv)
		if ok {
			f.rval.SetInt(int64(v))
		}
	case kI8:
		v, ok := safe_parse.ParseInt8Ok(tagv)
		if ok {
			f.rval.SetInt(int64(v))
		}
	case kI16:
		v, ok := safe_parse.ParseInt16Ok(tagv)
		if ok {
			f.rval.SetInt(int64(v))
		}
	case kI32:
		v, ok := safe_parse.ParseInt32Ok(tagv)
		if ok {
			f.rval.SetInt(int64(v))
		}
	case kI64:
		v, ok := safe_parse.ParseInt64Ok(tagv)
		if ok {
			f.rval.SetInt(v)
		}
	case kUint:
		v, ok := safe_parse.ParseUintOk(tagv)
		if ok {
			f.rval.SetUint(uint64(v))
		}
	case kU8:
		v, ok := safe_parse.ParseUint8Ok(tagv)
		if ok {
			f.rval.SetUint(uint64(v))
		}
	case kU16:
		v, ok := safe_parse.ParseUint16Ok(tagv)
		if ok {
			f.rval.SetUint(uint64(v))
		}
	case kU32:
		v, ok := safe_parse.ParseUint32Ok(tagv)
		if ok {
			f.rval.SetUint(uint64(v))
		}
	case kU64:
		v, ok := safe_parse.ParseUint64Ok(tagv)
		if ok {
			f.rval.SetUint(v)
		}
	case kUPtr:
		v, ok := safe_parse.ParseUintptrOk(tagv)
		if ok {
			f.rval.SetUint(uint64(v))
		}
	case kF32:
		v, ok := safe_parse.ParseFloat32Ok(tagv)
		if ok {
			f.rval.SetFloat(float64(v))
		}
	case kF64:
		v, ok := safe_parse.ParseFloat64Ok(tagv)
		if ok {
			f.rval.SetFloat(v)
		}
	case kString:
		f.rval.SetString(tagv)
	default:
		panic(errTagKindCheckErr)
	}
}
