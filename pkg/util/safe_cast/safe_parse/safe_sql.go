package safe_parse

import (
	"database/sql"
	"encoding/json" // json.RawMessage
	"github.com/google/uuid"
	"reflect"
	"time"
)

const (
	_DefKind               reflect.Kind = 100000 + reflect.UnsafePointer
	KStdTime               reflect.Kind = _DefKind + 1
	KStdTimePtr            reflect.Kind = _DefKind + 2
	KBytes                 reflect.Kind = _DefKind + 3
	KBytesPtr              reflect.Kind = _DefKind + 4
	KJSONRaw               reflect.Kind = _DefKind + 5
	KJSONRawPtr            reflect.Kind = _DefKind + 6
	KUUID                  reflect.Kind = _DefKind + 7
	KUUIDPtr               reflect.Kind = _DefKind + 8
	KRawBytes              reflect.Kind = _DefKind + 9
	KRawBytesPtr           reflect.Kind = _DefKind + 10
	KNullString            reflect.Kind = _DefKind + 11
	KNullStringPtr         reflect.Kind = _DefKind + 12
	KUnderlyingNullString  reflect.Kind = _DefKind + 13
	KNullInt64             reflect.Kind = _DefKind + 14
	KNullInt64Ptr          reflect.Kind = _DefKind + 15
	KUnderlyingNullInt64   reflect.Kind = _DefKind + 16
	KNullInt32             reflect.Kind = _DefKind + 17
	KNullInt32Ptr          reflect.Kind = _DefKind + 18
	KUnderlyingNullInt32   reflect.Kind = _DefKind + 19
	KNullInt16             reflect.Kind = _DefKind + 20
	KNullInt16Ptr          reflect.Kind = _DefKind + 21
	KUnderlyingNullInt16   reflect.Kind = _DefKind + 22
	KNullByte              reflect.Kind = _DefKind + 23
	KNullBytePtr           reflect.Kind = _DefKind + 24
	KUnderlyingNullByte    reflect.Kind = _DefKind + 25
	KNullFloat64           reflect.Kind = _DefKind + 26
	KNullFloat64Ptr        reflect.Kind = _DefKind + 27
	KUnderlyingNullFloat64 reflect.Kind = _DefKind + 28
	KNullBool              reflect.Kind = _DefKind + 29
	KNullBoolPtr           reflect.Kind = _DefKind + 30
	KUnderlyingNullBool    reflect.Kind = _DefKind + 31
	KNullTime              reflect.Kind = _DefKind + 32
	KNullTimePtr           reflect.Kind = _DefKind + 33
	KUnderlyingNullTime    reflect.Kind = _DefKind + 34
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
