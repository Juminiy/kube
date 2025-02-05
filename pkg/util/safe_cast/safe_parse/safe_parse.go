package safe_parse

import (
	"database/sql"
	"encoding/json" // json.RawMessage
	"github.com/Juminiy/kube/pkg/util"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Category int8

const (
	Null = Category(0)
	Bool = Category(1)
	Num  = Category(2)
	Time = Category(3)
	Text = Category(4)
)

type Type interface {
	Category() Category
	Bool() bool
	Number() Number
	Time() time.Time
	Text() string
	Get(kind reflect.Kind) (any, bool)
	GetByRT(rt reflect.Type) (any, bool)
}

type readable struct {
	boolV   *bool
	numberV *Number
	timeV   *time.Time
	stringV string
}

func (r readable) Category() Category {
	if r.timeV != nil {
		return Time
	} else if r.numberV != nil {
		return Num
	} else if r.boolV != nil {
		return Bool
	}
	return Text
}

func (r readable) Bool() bool {
	if r.boolV != nil {
		return *r.boolV
	}
	return false
}

func (r readable) Number() Number {
	if r.numberV != nil {
		return *r.numberV
	}
	return Number{}
}

func (r readable) Time() time.Time {
	if r.timeV != nil {
		return *r.timeV
	}
	return time.Time{}
}

func (r readable) Text() string {
	return r.stringV
}

func (r readable) Get(kind reflect.Kind) (v any, ok bool) {
	switch kind {
	case reflect.Bool:
		if r.boolV != nil {
			return *r.boolV, true
		}
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		if r.numberV != nil {
			return (*r.numberV).Get(kind)
		}
	case reflect.Interface, reflect.String:
		return r.stringV, true
	default:
		return r.getDefKind(kind)
	}
	return nil, false
}

func (r readable) GetByRT(rt reflect.Type) (v any, ok bool) {
	switch rt {
	case _TimeType:
		return r.Get(KStdTime)
	case _BytesType:
		return r.Get(KBytes)
	case _JSONRawType:
		return r.Get(KJSONRaw)
	case _UUIDType:
		return r.Get(KUUID)
	case _RawBytesType:
		return r.Get(KRawBytes)
	case _NullStringType:
		return r.Get(KNullString)
	case _NullInt64Type:
		return r.Get(KNullInt64)
	case _NullInt32Type:
		return r.Get(KNullInt32)
	case _NullInt16Type:
		return r.Get(KNullInt16)
	case _NullByteType:
		return r.Get(KNullByte)
	case _NullFloat64Type:
		return r.Get(KNullFloat64)
	case _NullBoolType:
		return r.Get(KNullBool)
	case _NullTimeType:
		return r.Get(KNullTime)
	}
	return nil, false
}

func (r readable) getDefKind(kind reflect.Kind) (v any, ok bool) {
	switch kind {
	case KStdTime:
		if r.timeV != nil {
			return *r.timeV, true
		}

	case KBytes:
		return []byte(r.stringV), true

	case KJSONRaw:
		return json.RawMessage(r.stringV), true

	case KUUID:
		if uuidV, err := uuid.Parse(r.stringV); err == nil {
			return uuidV, true
		}

	case KRawBytes:
		return sql.RawBytes(r.stringV), true

	case KNullString:
		return sql.NullString{
			String: r.stringV,
			Valid:  true,
		}, true

	case KNullInt64:
		if numV, ok := r.Number().Get(reflect.Int64); ok {
			return sql.NullInt64{
				Int64: numV.(int64),
				Valid: true,
			}, true
		}

	case KNullInt32:
		if numV, ok := r.Number().Get(reflect.Int32); ok {
			return sql.NullInt32{
				Int32: numV.(int32),
				Valid: true,
			}, true
		}

	case KNullInt16:
		if numV, ok := r.Number().Get(reflect.Int16); ok {
			return sql.NullInt16{
				Int16: numV.(int16),
				Valid: true,
			}, true
		}

	case KNullByte:
		if numV, ok := r.Number().Get(reflect.Uint8); ok {
			return sql.NullByte{
				Byte:  byte(numV.(uint8)),
				Valid: true,
			}, true
		}

	case KNullFloat64:
		if numV, ok := r.Number().Get(reflect.Float64); ok {
			return sql.NullFloat64{
				Float64: numV.(float64),
				Valid:   true,
			}, true
		}

	case KNullBool:
		if r.boolV != nil {
			return sql.NullBool{
				Bool:  *r.boolV,
				Valid: true,
			}, true
		}

	case KNullTime:
		if r.timeV != nil {
			return sql.NullTime{
				Time:  *r.timeV,
				Valid: true,
			}, true
		}

	default: // ignore case
		return nil, false
	}

	return nil, false
}

func Parse(s string) Type {
	readV := readable{}

	// string, reflect.String, Type.Text()
	readV.stringV = s

	// bool, reflect.Bool, Type.Bool()
	boolV, ok := ParseBool(s)
	if ok {
		readV.boolV = util.New(boolV)
	}

	// Number, reflect.Int~reflect.F64, Type.Number()
	numberV := ParseNumber(s)
	for _kind := reflect.Int; _kind <= reflect.Float64; _kind++ {
		if _, ok := numberV.Get(_kind); ok {
			readV.numberV = util.New(numberV)
			break
		}
	}

	// time.Time, Type.Time()
	timeV, ok := ParseTime(s)
	if ok {
		readV.timeV = util.New(timeV)
	}

	return readV
}
