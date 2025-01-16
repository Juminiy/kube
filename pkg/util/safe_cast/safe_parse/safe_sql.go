package safe_parse

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
	"time"
)

const (
	_DefKind     reflect.Kind = 100000 + reflect.UnsafePointer
	KStdTime     reflect.Kind = _DefKind + 1
	KBytes       reflect.Kind = _DefKind + 2
	KJSONRaw     reflect.Kind = _DefKind + 3
	KUUID        reflect.Kind = _DefKind + 4
	KRawBytes    reflect.Kind = _DefKind + 5
	KNullString  reflect.Kind = _DefKind + 6
	KNullInt64   reflect.Kind = _DefKind + 7
	KNullInt32   reflect.Kind = _DefKind + 8
	KNullInt16   reflect.Kind = _DefKind + 9
	KNullByte    reflect.Kind = _DefKind + 10
	KNullFloat64 reflect.Kind = _DefKind + 11
	KNullBool    reflect.Kind = _DefKind + 12
	KNullTime    reflect.Kind = _DefKind + 13
)

var _TimeType = reflect.TypeOf(time.Time{})
var _TimePType = reflect.TypeOf(&time.Time{})

var _BytesType = reflect.TypeOf([]byte{})
var _BytesPType = reflect.TypeOf(&[]byte{})

var _JSONRawType = reflect.TypeOf(json.RawMessage{})
var _JSONRawPType = reflect.TypeOf(&json.RawMessage{})

var _UUIDType = reflect.TypeOf(uuid.UUID{})
var _UUIDPType = reflect.TypeOf(&uuid.UUID{})

var _RawBytesType = reflect.TypeOf(sql.RawBytes{})
var _RawBytesPType = reflect.TypeOf(&sql.RawBytes{})

var _NullStringType = reflect.TypeOf(sql.NullString{})
var _NullStringPType = reflect.TypeOf(&sql.NullString{})

var _NullInt64Type = reflect.TypeOf(sql.NullInt64{})
var _NullInt64PType = reflect.TypeOf(&sql.NullInt64{})

var _NullInt32Type = reflect.TypeOf(sql.NullInt32{})
var _NullInt32PType = reflect.TypeOf(&sql.NullInt32{})

var _NullInt16Type = reflect.TypeOf(sql.NullInt16{})
var _NullInt16PType = reflect.TypeOf(&sql.NullInt16{})

var _NullByteType = reflect.TypeOf(sql.NullByte{})
var _NullBytePType = reflect.TypeOf(&sql.NullByte{})

var _NullFloat64Type = reflect.TypeOf(sql.NullFloat64{})
var _NullFloat64PType = reflect.TypeOf(&sql.NullFloat64{})

var _NullBoolType = reflect.TypeOf(sql.NullBool{})
var _NullBoolPType = reflect.TypeOf(&sql.NullBool{})

var _NullTimeType = reflect.TypeOf(sql.NullTime{})
var _NullTimePType = reflect.TypeOf(&sql.NullTime{})
